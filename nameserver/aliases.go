/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:20:15
 * @FilePath: \go-namesilo\nameserver\aliases.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package nameserver

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

	// Nameserver 相关错误
	ErrNameserverRequired   = namesilo.ErrNameserverRequired
	ErrNameserverNotFound   = namesilo.ErrNameserverNotFound
	ErrInvalidNameserver    = namesilo.ErrInvalidNameserver
	ErrNameserverInUse      = namesilo.ErrNameserverInUse
	ErrInvalidIPAddress     = namesilo.ErrInvalidIPAddress
	ErrIPAddressExceedLimit = namesilo.ErrIPAddressExceedLimit
)

// ============================================================================
// 状态码常量
// ============================================================================

var StatusSuccess = types.StatusSuccess

// ============================================================================
// 域名服务器请求类型
// ============================================================================

type (
	ChangeNameServersRequest          = types.ChangeNameServersRequest
	ListRegisteredNameServersRequest  = types.ListRegisteredNameServersRequest
	AddRegisteredNameServerRequest    = types.AddRegisteredNameServerRequest
	ModifyRegisteredNameServerRequest = types.ModifyRegisteredNameServerRequest
	DeleteRegisteredNameServerRequest = types.DeleteRegisteredNameServerRequest
)

// ============================================================================
// 域名服务器响应类型
// ============================================================================

type (
	RegisteredNameServer               = types.RegisteredNameServer
	ChangeNameServersResponse          = types.ChangeNameServersResponse
	ListRegisteredNameServersResponse  = types.ListRegisteredNameServersResponse
	AddRegisteredNameServerResponse    = types.AddRegisteredNameServerResponse
	ModifyRegisteredNameServerResponse = types.ModifyRegisteredNameServerResponse
	DeleteRegisteredNameServerResponse = types.DeleteRegisteredNameServerResponse
)
