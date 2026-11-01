/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\transfer\aliases.go
 * @Description: Transfer 转移模块类型别名
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package transfer

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

	// Transfer 相关错误
	ErrTransferNotFound       = namesilo.ErrTransferNotFound
	ErrAuthCodeRequired       = namesilo.ErrAuthCodeRequired
	ErrInvalidAuthCode        = namesilo.ErrInvalidAuthCode
	ErrTransferIDRequired     = namesilo.ErrTransferIDRequired
	ErrDomainTransferDenied   = namesilo.ErrDomainTransferDenied
	ErrTransferAlreadyPending = namesilo.ErrTransferAlreadyPending
	ErrInsufficientBalance    = namesilo.ErrInsufficientBalance
)

// ============================================================================
// Transfer 请求类型
// ============================================================================

type (
	RetrieveAuthCodeRequest           = types.RetrieveAuthCodeRequest
	CheckTransferStatusRequest        = types.CheckTransferStatusRequest
	TransferUpdateResubmitRequest     = types.TransferUpdateResubmitRequest
	TransferUpdateResendEmailRequest  = types.TransferUpdateResendEmailRequest
	TransferUpdateChangeEPPCodeRequest = types.TransferUpdateChangeEPPCodeRequest
)

// ============================================================================
// Transfer 响应类型
// ============================================================================

type (
	RetrieveAuthCodeResponse           = types.RetrieveAuthCodeResponse
	CheckTransferStatusResponse        = types.CheckTransferStatusResponse
	TransferUpdateResubmitResponse     = types.TransferUpdateResubmitResponse
	TransferUpdateResendEmailResponse  = types.TransferUpdateResendEmailResponse
	TransferUpdateChangeEPPCodeResponse = types.TransferUpdateChangeEPPCodeResponse
	TransferStatus                     = types.TransferStatus
)
