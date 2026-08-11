/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 21:15:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 21:15:00
 * @FilePath: \go-namesilo\client\client_test.go
 * @Description: 测试客户端创建和配置优先级
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"testing"
	"time"

	logger "github.com/kamalyes/go-logger"
	"github.com/stretchr/testify/assert"
)

// TestNew_WithAPIKeyParameter 测试使用参数传入 API Key
func TestNew_WithAPIKeyParameter(t *testing.T) {
	apiKey := "test-api-key-from-param"
	client, err := New(apiKey)

	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, client, "客户端不应该为 nil")
	assert.Equal(t, apiKey, client.config.APIKey, "API Key 应该等于参数传入的值")
	assert.Equal(t, DefaultBaseURL, client.config.BaseURL, "BaseURL 应该使用默认值")
	assert.Equal(t, DefaultTimeout, client.config.Timeout, "Timeout 应该使用默认值")
}

// TestNew_WithEmptyAPIKey 测试传入空 API Key
func TestNew_WithEmptyAPIKey(t *testing.T) {
	client, err := New("")

	assert.Error(t, err, "应该返回错误")
	assert.Equal(t, ErrAPIKeyRequired, err, "错误应该是 ErrAPIKeyRequired")
	assert.Nil(t, client, "客户端应该为 nil")
}

// TestNew_OptsOverrideParameter 测试 opts 中的 WithAPIKey 覆盖参数
func TestNew_OptsOverrideParameter(t *testing.T) {
	paramAPIKey := "param-api-key"
	optsAPIKey := "opts-api-key"

	client, err := New(paramAPIKey, WithAPIKey(optsAPIKey))

	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, client, "客户端不应该为 nil")
	assert.Equal(t, optsAPIKey, client.config.APIKey, "API Key 应该等于 opts 中设置的值")
	assert.NotEqual(t, paramAPIKey, client.config.APIKey, "API Key 不应该等于参数传入的值")
}

// TestNew_WithMultipleOptions 测试使用多个配置选项
func TestNew_WithMultipleOptions(t *testing.T) {
	apiKey := "test-api-key"
	customBaseURL := "https://custom.namesilo.com"
	customTimeout := 30 * time.Second
	customVersion := "2.0"

	client, err := New(apiKey,
		WithBaseURL(customBaseURL),
		WithTimeout(customTimeout),
		WithAPIVersion(customVersion),
		WithDebug(false),
		WithResponseType(ResponseTypeJSON),
	)

	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, client, "客户端不应该为 nil")
	assert.Equal(t, apiKey, client.config.APIKey, "API Key 应该正确")
	assert.Equal(t, "https://custom.namesilo.com", client.config.BaseURL, "BaseURL 应该被覆盖")
	assert.Equal(t, customTimeout, client.config.Timeout, "Timeout 应该被覆盖")
	assert.Equal(t, customVersion, client.config.APIVersion, "API Version 应该被覆盖")
	assert.False(t, client.config.Debug, "Debug 应该被设置为 false")
	assert.Equal(t, ResponseTypeJSON, client.config.ResponseType, "ResponseType 应该被设置为 JSON")
}

// TestNew_OptsOverrideMultipleConfigs 测试 opts 覆盖多个配置包括 APIKey
func TestNew_OptsOverrideMultipleConfigs(t *testing.T) {
	paramAPIKey := "param-key"
	optsAPIKey := "opts-key"
	customBaseURL := "https://override.namesilo.com"

	client, err := New(paramAPIKey,
		WithBaseURL(customBaseURL),
		WithAPIKey(optsAPIKey), // 这应该覆盖参数传入的 APIKey
		WithDebug(false),
	)

	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, client, "客户端不应该为 nil")
	assert.Equal(t, optsAPIKey, client.config.APIKey, "API Key 应该等于 opts 中设置的值")
	assert.Equal(t, "https://override.namesilo.com", client.config.BaseURL, "BaseURL 应该正确")
	assert.False(t, client.config.Debug, "Debug 应该被覆盖")
}

// TestNew_WithLogger 测试自定义 Logger
func TestNew_WithLogger(t *testing.T) {
	apiKey := "test-api-key"
	testLogger := &TestLogger{
		EmptyLogger: logger.NewEmptyLogger(),
		DebugLogs:   make([]string, 0),
		InfoLogs:    make([]string, 0),
		WarnLogs:    make([]string, 0),
		ErrorLogs:   make([]string, 0),
	}

	client, err := New(apiKey,
		WithDebug(true),
		WithLogger(testLogger),
	)

	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, client, "客户端不应该为 nil")
	assert.Equal(t, testLogger, client.logger, "Logger 应该被正确设置")
}

// TestNew_DebugOffIgnoresLogger 测试 Debug 关闭时忽略 Logger
func TestNew_DebugOffIgnoresLogger(t *testing.T) {
	apiKey := "test-api-key"
	testLogger := &TestLogger{
		EmptyLogger: logger.NewEmptyLogger(),
		DebugLogs:   make([]string, 0),
		InfoLogs:    make([]string, 0),
		WarnLogs:    make([]string, 0),
		ErrorLogs:   make([]string, 0),
	}

	client, err := New(apiKey,
		WithDebug(false),
		WithLogger(testLogger),
	)

	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, client, "客户端不应该为 nil")
	assert.Equal(t, NoLogger, client.logger, "Debug 关闭时应该使用 NoLogger")
}

// TestNew_ConfigPriority 测试配置优先级：参数 < opts
func TestNew_ConfigPriority(t *testing.T) {
	t.Run("参数设置APIKey,opts未设置", func(t *testing.T) {
		client, err := New("param-key")
		assert.NoError(t, err)
		assert.Equal(t, "param-key", client.config.APIKey, "应该使用参数的值")
	})

	t.Run("参数设置APIKey,opts也设置", func(t *testing.T) {
		client, err := New("param-key", WithAPIKey("opts-key"))
		assert.NoError(t, err)
		assert.Equal(t, "opts-key", client.config.APIKey, "应该使用 opts 的值（优先级更高）")
	})

	t.Run("opts最后设置的值生效", func(t *testing.T) {
		client, err := New("param-key",
			WithAPIKey("opts-key-1"),
			WithAPIKey("opts-key-2"), // 最后一个应该生效
		)
		assert.NoError(t, err)
		assert.Equal(t, "opts-key-2", client.config.APIKey, "应该使用最后设置的 opts 值")
	})
}

// TestGetConfig 测试获取配置
func TestGetConfig(t *testing.T) {
	apiKey := "test-api-key"
	client, err := New(apiKey)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	config := client.GetConfig()
	assert.NotNil(t, config, "配置不应该为 nil")
	assert.Equal(t, apiKey, config.APIKey, "配置中的 API Key 应该正确")
}

// TestGetHTTPClient 测试获取 HTTP 客户端
func TestGetHTTPClient(t *testing.T) {
	apiKey := "test-api-key"
	client, err := New(apiKey)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	httpClient := client.GetHTTPClient()
	assert.NotNil(t, httpClient, "HTTP 客户端不应该为 nil")
}
