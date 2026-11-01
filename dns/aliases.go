/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:22:15
 * @FilePath: \go-namesilo\dns\aliases.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package dns

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

	// DNS 相关错误
	ErrRecordIDRequired    = namesilo.ErrRecordIDRequired
	ErrRecordTypeRequired  = namesilo.ErrRecordTypeRequired
	ErrRecordValueRequired = namesilo.ErrRecordValueRequired
	ErrInvalidRecordType   = namesilo.ErrInvalidRecordType
	ErrRecordNotFound      = namesilo.ErrRecordNotFound
	ErrInvalidTTL          = namesilo.ErrInvalidTTL
	ErrInvalidPriority     = namesilo.ErrInvalidPriority
	ErrRRTypeRequired      = namesilo.ErrRRTypeRequired
	ErrRRHostRequired      = namesilo.ErrRRHostRequired
	ErrRRValueRequired     = namesilo.ErrRRValueRequired
	ErrDigestRequired      = namesilo.ErrDigestRequired
)

// ============================================================================
// DNS 记录类型常量
// ============================================================================

const (
	RecordTypeA     = types.RecordTypeA
	RecordTypeAAAA  = types.RecordTypeAAAA
	RecordTypeCNAME = types.RecordTypeCNAME
	RecordTypeMX    = types.RecordTypeMX
	RecordTypeTXT   = types.RecordTypeTXT
	RecordTypeSRV   = types.RecordTypeSRV
	RecordTypeCAA   = types.RecordTypeCAA
)

// ============================================================================
// DNS 请求类型
// ============================================================================

type (
	DNSListRecordsRequest     = types.DNSListRecordsRequest
	DNSAddRecordRequest       = types.DNSAddRecordRequest
	DNSUpdateRecordRequest    = types.DNSUpdateRecordRequest
	DNSDeleteRecordRequest    = types.DNSDeleteRecordRequest
	DNSSecListRecordsRequest  = types.DNSSecListRecordsRequest
	DNSSecAddRecordRequest    = types.DNSSecAddRecordRequest
	DNSSecDeleteRecordRequest = types.DNSSecDeleteRecordRequest
)

// ============================================================================
// DNS 响应类型
// ============================================================================

type (
	DNSRecord                  = types.DNSRecord
	DNSListRecordsResponse     = types.DNSListRecordsResponse
	DNSAddRecordResponse       = types.DNSAddRecordResponse
	DNSUpdateRecordResponse    = types.DNSUpdateRecordResponse
	DNSDeleteRecordResponse    = types.DNSDeleteRecordResponse
	DNSSecListRecordsResponse  = types.DNSSecListRecordsResponse
	DNSSecAddRecordResponse    = types.DNSSecAddRecordResponse
	DNSSecDeleteRecordResponse = types.DNSSecDeleteRecordResponse
)
