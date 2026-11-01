/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:15:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:05:21
 * @FilePath: \go-namesilo\forwarding\aliases.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package forwarding

import (
	namesilo "github.com/HitGameAI/apex-namesilo"
	"github.com/HitGameAI/apex-namesilo/types"
)

// ============================================================================
// 接口定义
// ============================================================================

// ClientInterface 客户端接口别名
type ClientInterface = types.ClientInterface

// ============================================================================
// 错误类型
// ============================================================================

var (
	// 通用错误
	ErrDomainRequired = namesilo.ErrDomainRequired

	// Forwarding 相关错误
	ErrURLRequired            = namesilo.ErrURLRequired
	ErrEmailRequired          = namesilo.ErrEmailRequired
	ErrSubdomainRequired      = namesilo.ErrSubdomainRequired
	ErrForwardNotFound        = namesilo.ErrForwardNotFound
	ErrInvalidURL             = namesilo.ErrInvalidURL
	ErrInvalidEmail           = namesilo.ErrInvalidEmail
	ErrForwardListExceedLimit = namesilo.ErrForwardListExceedLimit
)

// ============== Type Aliases ==============

// ForwardDomainRequest 设置域名转发请求
type ForwardDomainRequest = types.ForwardDomainRequest

// ForwardDomainResponse 设置域名转发响应
type ForwardDomainResponse = types.ForwardDomainResponse

// ForwardSubdomainRequest 设置子域名转发请求
type ForwardSubdomainRequest = types.ForwardSubdomainRequest

// ForwardSubdomainResponse 设置子域名转发响应
type ForwardSubdomainResponse = types.ForwardSubdomainResponse

// DeleteForwardRequest 删除域名/子域名转发请求
type DeleteForwardRequest = types.DeleteForwardRequest

// DeleteForwardResponse 删除转发响应
type DeleteForwardResponse = types.DeleteForwardResponse

// ListEmailForwardsRequest 列出邮件转发请求
type ListEmailForwardsRequest = types.ListEmailForwardsRequest

// ListEmailForwardsResponse 列出邮件转发响应
type ListEmailForwardsResponse = types.ListEmailForwardsResponse

// EmailForward 邮件转发信息
type EmailForward = types.EmailForward

// ConfigureEmailForwardRequest 配置邮件转发请求
type ConfigureEmailForwardRequest = types.ConfigureEmailForwardRequest

// ConfigureEmailForwardResponse 配置邮件转发响应
type ConfigureEmailForwardResponse = types.ConfigureEmailForwardResponse

// DeleteEmailForwardRequest 删除邮件转发请求
type DeleteEmailForwardRequest = types.DeleteEmailForwardRequest

// DeleteEmailForwardResponse 删除邮件转发响应
type DeleteEmailForwardResponse = types.DeleteEmailForwardResponse
