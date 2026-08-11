/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 21:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 21:00:00
 * @FilePath: \go-namesilo\client\request_test.go
 * @Description: 测试请求和日志打印
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package client

import (
	"context"
	"fmt"
	"strings"
	"testing"

	logger "github.com/kamalyes/go-logger"
	"github.com/stretchr/testify/assert"
)

// TestLogger 测试用的日志记录器，嵌入 *logger.EmptyLogger 自动实现 ILogger 全部方法
type TestLogger struct {
	*logger.EmptyLogger
	DebugLogs []string
	InfoLogs  []string
	WarnLogs  []string
	ErrorLogs []string
}

func (l *TestLogger) DebugContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.DebugLogs = append(l.DebugLogs, logStr)
	fmt.Println("[DEBUG]", logStr)
}

func (l *TestLogger) InfoContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.InfoLogs = append(l.InfoLogs, logStr)
	fmt.Println("[INFO]", logStr)
}

func (l *TestLogger) WarnContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.WarnLogs = append(l.WarnLogs, logStr)
	fmt.Println("[WARN]", logStr)
}

func (l *TestLogger) ErrorContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	logStr := fmt.Sprintf("%s", msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			logStr += fmt.Sprintf(" | %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	l.ErrorLogs = append(l.ErrorLogs, logStr)
	fmt.Println("[ERROR]", logStr)
}

// TestDoRequestWithLogger 测试带日志的请求
func TestDoRequestWithLogger(t *testing.T) {
	// 创建测试日志记录器
	testLogger := &TestLogger{
		EmptyLogger: logger.NewEmptyLogger(),
		DebugLogs:   make([]string, 0),
		InfoLogs:    make([]string, 0),
		WarnLogs:    make([]string, 0),
		ErrorLogs:   make([]string, 0),
	}

	// 创建客户端
	c, err := New("test-api-key",
		WithDebug(true),
		WithLogger(testLogger),
	)
	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, c, "客户端不应该为 nil")

	// 执行请求（这会失败，但我们主要是测试日志）
	ctx := context.Background()
	params := map[string]string{
		"domain": "example.com",
	}

	_, _ = c.DoRequest(ctx, "getDomainInfo", params)

	// 验证日志
	t.Run("验证请求日志", func(t *testing.T) {
		assert.NotEmpty(t, testLogger.DebugLogs, "应该有 Debug 日志记录")

		requestLog := ""
		for _, log := range testLogger.DebugLogs {
			if strings.Contains(log, "NameSilo API Request") {
				requestLog = log
				break
			}
		}

		assert.NotEmpty(t, requestLog, "应该找到请求日志")

		// 验证日志包含必要信息（新格式）
		assert.Contains(t, requestLog, "Trace ID", "请求日志应该包含 Trace ID")
		assert.Contains(t, requestLog, "Operation", "请求日志应该包含 Operation")
		assert.Contains(t, requestLog, "getDomainInfo", "请求日志应该包含操作名称")
		assert.Contains(t, requestLog, "Request URL", "请求日志应该包含 Request URL")
		assert.Contains(t, requestLog, "Query Params", "请求日志应该包含 Query Params")

		t.Logf("请求日志: %s", requestLog)
	})

	t.Run("验证响应日志", func(t *testing.T) {
		responseLog := ""
		for _, log := range testLogger.DebugLogs {
			if strings.Contains(log, "NameSilo API Request/Response") {
				responseLog = log
				break
			}
		}

		if responseLog != "" {
			// 验证响应日志包含必要信息
			assert.Contains(t, responseLog, "Status", "响应日志应该包含 Status")
			assert.Contains(t, responseLog, "Headers", "响应日志应该包含 Headers")
			assert.Contains(t, responseLog, "Body Length", "响应日志应该包含 Body Length")
			assert.Contains(t, responseLog, "Body", "响应日志应该包含 Body")

			t.Logf("响应日志: %s", responseLog)
		}
	})
}

// TestDoRequestWithEmptyLogger 测试使用空日志记录器
func TestDoRequestWithEmptyLogger(t *testing.T) {
	// 创建客户端，使用空日志记录器
	c, err := New("test-api-key",
		WithDebug(true),
		WithLogger(NoLogger),
	)
	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, c, "客户端不应该为 nil")
	assert.NotNil(t, c.logger, "logger 应该被设置")

	// 执行请求（这会失败，但不应该 panic）
	ctx := context.Background()
	params := map[string]string{
		"domain": "example.com",
	}

	assert.NotPanics(t, func() {
		_, _ = c.DoRequest(ctx, "getDomainInfo", params)
	}, "使用空日志记录器不应该 panic")

	t.Log("使用空日志记录器测试通过")
}

// TestDoRequestWithoutLogger 测试不设置日志记录器的情况
func TestDoRequestWithoutLogger(t *testing.T) {
	// 创建客户端，不设置日志记录器
	c, err := New("test-api-key",
		WithDebug(true),
	)
	assert.NoError(t, err, "创建客户端应该成功")
	assert.NotNil(t, c, "客户端不应该为 nil")

	// 验证 logger 被自动设置为 NoLogger
	assert.NotNil(t, c.logger, "logger 应该自动设置为 NoLogger")

	// 执行请求（这会失败，但不应该 panic）
	ctx := context.Background()
	params := map[string]string{
		"domain": "example.com",
	}

	assert.NotPanics(t, func() {
		_, _ = c.DoRequest(ctx, "getDomainInfo", params)
	}, "不设置日志记录器不应该 panic")

	t.Log("不设置日志记录器测试通过")
}
