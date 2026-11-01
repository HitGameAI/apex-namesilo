/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 19:23:00
 * @FilePath: \go-namesilo\types\common.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package types

import (
	logger "github.com/HitGameAI/apex-toolbox/pkg/logger"
)

// Logger 日志接口，直接复用 go-logger.ILogger
// go-logger.ILogger 提供了 DebugContext/InfoContext/WarnContext/ErrorContext 等方法，
// 与 go-namesilo 的日志调用方式完全兼容
type Logger = logger.ILogger

// NoLogger 直接复用 go-logger 的全局空日志器；未注入 Logger 时使用该实例，不输出任何日志
var NoLogger = logger.NoLogger

// NewEmptyLogger 返回一个空日志记录器（go-logger.NoLogger 单例）
func NewEmptyLogger() Logger {
	return NoLogger
}
