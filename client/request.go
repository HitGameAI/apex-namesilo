/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 21:37:18
 * @FilePath: \go-namesilo\client\request.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"

	namesilo "github.com/kamalyes/go-namesilo"
	"github.com/kamalyes/go-toolbox/pkg/errorx"
	"github.com/kamalyes/go-toolbox/pkg/osx"
)

// DoRequest 执行 API 请求（公开给子模块使用）
func (c *Client) DoRequest(ctx context.Context, operation string, params map[string]string) ([]byte, error) {
	// 构建请求参数
	queryParams := url.Values{}

	// 添加业务参数（先添加，以便系统参数可以覆盖）
	for key, value := range params {
		if value != "" {
			queryParams.Set(key, value)
		}
	}

	// 强制设置系统参数（这些参数不应被 params 覆盖）
	queryParams.Set("version", c.config.APIVersion)
	queryParams.Set("type", c.config.ResponseType.String())
	queryParams.Set("key", c.config.APIKey)

	// 构建完整 URL（NameSilo API 使用 GET 方法）
	baseURL, err := osx.JoinURL(c.config.BaseURL, operation)
	if err != nil {
		return nil, namesilo.WrapError(ErrCodeAPIRequest, "failed to build request URL", err)
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// 用于记录响应信息
	var respStatus string
	var respHeaders interface{}
	var respBody []byte
	var apiErr error
	var recoverErr error

	// defer 记录请求和响应日志
	defer func() {
		c.logRequestResponse(ctx, operation, fullURL, respStatus, respHeaders, respBody, apiErr, recoverErr)
	}()

	// 使用 errorx.Recover 捕获 panic
	recoverErr = errorx.Recover(func() {
		// 使用 GET 请求（NameSilo API 标准方法）
		// httpx 已修复参数覆盖问题，现在可以直接传完整 URL
		resp, err := c.httpClient.Get(fullURL).WithContext(ctx).Send()
		if err != nil {
			apiErr = err
			return
		}
		defer resp.Close()

		// 记录响应状态和头
		respStatus = resp.Status
		respHeaders = resp.Header

		// 读取响应体
		body, err := resp.Bytes()
		if err != nil {
			apiErr = err
			return
		}

		respBody = body
	})

	// 如果发生 panic，返回错误
	if recoverErr != nil {
		return nil, namesilo.WrapError(ErrCodeAPIRequest, fmt.Sprintf("panic occurred during request to operation: %s", operation), recoverErr)
	}

	// 如果有 API 错误，返回错误
	if apiErr != nil {
		return nil, namesilo.WrapError(ErrCodeAPIRequest, fmt.Sprintf("failed to send request to operation: %s", operation), apiErr)
	}

	// 检查响应是否为空
	if len(respBody) == 0 {
		apiErr = fmt.Errorf("empty response body")
		return nil, namesilo.NewError(ErrCodeAPIResponse, fmt.Sprintf("received empty response from NameSilo API for operation: %s", operation))
	}

	return respBody, nil
}

// ParseResponse 解析响应（公开给子模块使用）
func (c *Client) ParseResponse(data []byte, v interface{}) error {
	// 检查响应是否为空
	if len(data) == 0 {
		return namesilo.WrapError(ErrCodeAPIResponse, "received empty response from NameSilo API", fmt.Errorf("empty response body"))
	}

	// 根据配置的响应类型进行解析
	var err error
	switch c.config.ResponseType {
	case ResponseTypeJSON:
		err = json.Unmarshal(data, v)
	default:
		// 默认使用 XML 解析
		err = xml.Unmarshal(data, v)
	}

	if err != nil {
		return namesilo.WrapError(ErrCodeAPIResponse,
			fmt.Sprintf("failed to unmarshal response as %s, response preview: %s", c.config.ResponseType.String(), string(data)), err)
	}

	return nil
}

// logRequestResponse 记录请求和响应日志
func (c *Client) logRequestResponse(ctx context.Context, operation, fullURL string,
	respStatus string, respHeaders interface{}, respBody []byte, apiErr, recoverErr error) {

	if recoverErr != nil {
		c.logger.ErrorContextKV(ctx, "🔗 NameSilo API panic",
			"operation", operation,
			"url", fullURL,
			"error", recoverErr.Error(),
		)
		return
	}

	if apiErr != nil {
		if errors.Is(apiErr, context.DeadlineExceeded) || errors.Is(apiErr, context.Canceled) {
			c.logger.WarnContextKV(ctx, "⏱️ NameSilo API timeout",
				"operation", operation,
				"url", fullURL,
				"error", apiErr.Error(),
			)
			return
		}
		c.logger.ErrorContextKV(ctx, "❌ NameSilo API error",
			"operation", operation,
			"url", fullURL,
			"status", respStatus,
			"error", apiErr.Error(),
		)
		return
	}

	c.logger.InfoContextKV(ctx, "🎯 NameSilo API request",
		"operation", operation,
		"url", fullURL,
		"status", respStatus,
		"headers", respHeaders,
		"bodyLen", len(respBody),
	)
	if len(respBody) > 0 {
		c.logger.DebugContextKV(ctx, "📏 NameSilo API response body",
			"operation", operation,
			"body", string(respBody),
		)
	}
}
