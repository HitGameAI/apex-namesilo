/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:15:17
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 21:00:27
 * @FilePath: \go-namesilo\client\client.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
	"github.com/HitGameAI/apex-toolbox/pkg/mathx"
)

// Client NameSilo API 客户端
type Client struct {
	config     *Config
	httpClient *httpx.Client
	logger     Logger // 日志记录器
}

// New 创建一个新的 NameSilo 客户端
//
// 参数:
//   - apiKey: NameSilo API Key，必填。作为默认值，可被 opts 中的 WithAPIKey 覆盖
//   - opts: 可选配置项，用于自定义客户端行为
//
// 配置优先级:
//   - apiKey 参数会作为 APIKey 的默认值
//   - 如果 opts 中使用了 WithAPIKey，则会覆盖 apiKey 参数
//   - 其他配置项（BaseURL、Timeout 等）如果在 opts 中指定，会覆盖默认配置
//
// 示例:
//
//	示例 1: 使用参数传入的 API Key
//	client, err := New("your-api-key")
//	结果: client.config.APIKey = "your-api-key"
//
//	示例 2: opts 中的 WithAPIKey 会覆盖参数
//	client, err := New("default-key", WithAPIKey("override-key"))
//	结果: client.config.APIKey = "override-key" (opts 优先级更高)
//
//	示例 3: 自定义其他配置
//	client, err := New("your-api-key",
//	    WithBaseURL("https://custom.namesilo.com"),
//	    WithTimeout(30*time.Second),
//	    WithDebug(true),
//	)
//	结果:
//	client.config.APIKey = "your-api-key"
//	client.config.BaseURL = "https://custom.namesilo.com"
//	client.config.Timeout = 30s
//	client.config.Debug = true
func New(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, ErrAPIKeyRequired
	}

	// 初始化配置，使用参数传入的 APIKey 作为默认值
	config := &Config{
		APIKey:       apiKey,
		BaseURL:      DefaultBaseURL,
		PublicURL:    DefaultPublicURL,
		Timeout:      DefaultTimeout,
		Debug:        true,
		ResponseType: DefaultType, // 默认使用 XML
		APIVersion:   DefaultAPIVersion,
	}

	// 应用选项，opts 中的设置会覆盖默认配置
	for _, opt := range opts {
		opt(config)
	}

	// 规范化 URL，确保包含协议前缀
	config.BaseURL = httpx.NormalizeBaseURL(config.BaseURL)
	config.PublicURL = httpx.NormalizeBaseURL(config.PublicURL)

	// 创建 HTTP 客户端
	httpClient := httpx.NewClient(
		httpx.WithTimeout(config.Timeout),
	)

	// 设置日志记录器
	logger := mathx.IF(config.Logger == nil || !config.Debug, NoLogger, config.Logger)

	client := &Client{
		config:     config,
		httpClient: httpClient,
		logger:     logger,
	}

	// 打印最终配置信息（仅在 Debug 模式下）
	logger.InfoContext(context.TODO(), "NameSilo Client 初始化完成",
		"APIKey", maskAPIKey(config.APIKey),
		"BaseURL", config.BaseURL,
		"Timeout", config.Timeout,
		"Debug", config.Debug,
		"ResponseType", config.ResponseType,
		"APIVersion", config.APIVersion,
	)

	return client, nil
}

// maskAPIKey 遮蔽 API Key 的敏感部分，只显示前后几位
func maskAPIKey(apiKey string) string {
	if len(apiKey) <= 8 {
		return "***"
	}
	return apiKey[:4] + "****" + apiKey[len(apiKey)-4:]
}

// GetConfig 获取客户端配置
func (c *Client) GetConfig() *Config {
	return c.config
}

// GetHTTPClient 获取 HTTP 客户端
func (c *Client) GetHTTPClient() *httpx.Client {
	return c.httpClient
}
