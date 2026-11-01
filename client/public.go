/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:15:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:50:15
 * @FilePath: \go-namesilo\client\public.go
 * @Description: 公开 API 请求方法
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"encoding/json"
	"fmt"

	namesilo "github.com/HitGameAI/apex-namesilo"
	"github.com/HitGameAI/apex-toolbox/pkg/osx"
)

// DoPublicRequest 执行公开 API 请求（不需要 API Key）
// 参数:
//   - ctx: 上下文
//   - method: HTTP 方法（GET, POST 等）
//   - path: API 路径（相对于 PublicURL）
//   - body: 请求体（POST 时使用）
//   - result: 响应结果对象
func (c *Client) DoPublicRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	// 构建完整 URL
	fullURL, err := osx.JoinURL(c.config.PublicURL, path)
	if err != nil {
		return namesilo.WrapError(ErrCodeAPIRequest, "failed to build public request URL", err)
	}

	// 创建请求
	req := c.httpClient.NewRequest(method, fullURL).
		WithContext(ctx).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36").
		SetHeader("Accept", "application/json")

	if method == "POST" {
		req.SetHeader("Content-Type", "application/json")
		if body != nil {
			req.SetBody(body)
		}
	}

	resp, err := req.Send()
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		// 读取错误响应体以便调试
		bodyBytes, _ := resp.Bytes()
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	if result != nil {
		bodyBytes, err := resp.Bytes()
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}
		if err := json.Unmarshal(bodyBytes, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
