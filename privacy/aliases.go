/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:00:00
 * @FilePath: \go-namesilo\privacy\aliases.go
 * @Description: Privacy 隐私保护模块类型别名
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package privacy

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

	// Privacy 相关错误
	ErrPrivacyAlreadyEnabled  = namesilo.ErrPrivacyAlreadyEnabled
	ErrPrivacyAlreadyDisabled = namesilo.ErrPrivacyAlreadyDisabled
	ErrInsufficientBalance    = namesilo.ErrInsufficientBalance
)

// ============================================================================
// Privacy 请求类型
// ============================================================================

type (
	AddPrivacyRequest    = types.AddPrivacyRequest
	RemovePrivacyRequest = types.RemovePrivacyRequest
)

// ============================================================================
// Privacy 响应类型
// ============================================================================

type (
	AddPrivacyResponse    = types.AddPrivacyResponse
	RemovePrivacyResponse = types.RemovePrivacyResponse
)
