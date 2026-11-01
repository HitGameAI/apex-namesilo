/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:22:15
 * @FilePath: \go-namesilo\contact\aliases.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package contact

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
	ErrAPIKeyRequired = namesilo.ErrAPIKeyRequired
	ErrDomainRequired = namesilo.ErrDomainRequired

	// Contact 相关错误
	ErrContactIDRequired   = namesilo.ErrContactIDRequired
	ErrFirstNameRequired   = namesilo.ErrFirstNameRequired
	ErrLastNameRequired    = namesilo.ErrLastNameRequired
	ErrAddressRequired     = namesilo.ErrAddressRequired
	ErrCityRequired        = namesilo.ErrCityRequired
	ErrStateRequired       = namesilo.ErrStateRequired
	ErrZipRequired         = namesilo.ErrZipRequired
	ErrCountryRequired     = namesilo.ErrCountryRequired
	ErrEmailRequired       = namesilo.ErrEmailRequired
	ErrPhoneRequired       = namesilo.ErrPhoneRequired
	ErrContactRoleRequired = namesilo.ErrContactRoleRequired
	ErrContactNotFound     = namesilo.ErrContactNotFound
	ErrContactInUse        = namesilo.ErrContactInUse
	ErrInvalidContactData  = namesilo.ErrInvalidContactData
)

// ============================================================================
// Contact 请求类型
// ============================================================================

type (
	ContactAddRequest             = types.ContactAddRequest
	ContactUpdateRequest          = types.ContactUpdateRequest
	ContactDomainAssociateRequest = types.ContactDomainAssociateRequest
	ContactDeleteRequest          = types.ContactDeleteRequest
	ContactListRequest            = types.ContactListRequest
)

// ============================================================================
// Contact 响应类型
// ============================================================================

type (
	ContactAddResponse             = types.ContactAddResponse
	ContactUpdateResponse          = types.ContactUpdateResponse
	ContactDomainAssociateResponse = types.ContactDomainAssociateResponse
	ContactDeleteResponse          = types.ContactDeleteResponse
	ContactListResponse            = types.ContactListResponse
)
