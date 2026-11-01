/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:15:15
 * @FilePath: \go-namesilo\domains\aliases.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */

package domains

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

	// Domain 相关错误
	ErrDomainNotFound       = namesilo.ErrDomainNotFound
	ErrDomainAlreadyExists  = namesilo.ErrDomainAlreadyExists
	ErrDomainLocked         = namesilo.ErrDomainLocked
	ErrDomainNotAvailable   = namesilo.ErrDomainNotAvailable
	ErrInvalidDomain        = namesilo.ErrInvalidDomain
	ErrDomainTransferDenied = namesilo.ErrDomainTransferDenied
	ErrInsufficientBalance  = namesilo.ErrInsufficientBalance
	// Domain 操作相关错误
	ErrSubDomainRequired      = namesilo.ErrSubDomainRequired
	ErrProtocolRequired       = namesilo.ErrProtocolRequired
	ErrAddressForwardRequired = namesilo.ErrAddressForwardRequired
	ErrMethodRequired         = namesilo.ErrMethodRequired
	ErrDomainsRequired        = namesilo.ErrDomainsRequired
	ErrYearsOutOfRange        = namesilo.ErrYearsOutOfRange
	ErrDomainsExceedLimit     = namesilo.ErrDomainsExceedLimit
	ErrRecipientLoginRequired = namesilo.ErrRecipientLoginRequired
)

// ============================================================================
// 状态码常量
// ============================================================================

var StatusSuccess = types.StatusSuccess

// ============================================================================
// 分页默认值
// ============================================================================
var (
	DefaultPage     = types.DefaultPage
	DefaultPageSize = types.DefaultPageSize
)

var (
	DefaultAllowedTLDs = types.DefaultAllowedTLDs
)

// ============================================================================
// 域名请求类型
// ============================================================================

type (
	AddAutoRenewalRequest               = types.AddAutoRenewalRequest
	RemoveAutoRenewalRequest            = types.RemoveAutoRenewalRequest
	CheckRegisterAvailabilityRequest    = types.CheckRegisterAvailabilityRequest
	RegisterDomainRequest               = types.RegisterDomainRequest
	RenewDomainRequest                  = types.RenewDomainRequest
	DomainLockRequest                   = types.DomainLockRequest
	DomainUnlockRequest                 = types.DomainUnlockRequest
	WhoisRequest                        = types.WhoisRequest
	ListDomainsRequest                  = types.ListDomainsRequest
	GetDomainInfoRequest                = types.GetDomainInfoRequest
	TransferDomainRequest               = types.TransferDomainRequest
	RegisterDomainDropRequest           = types.RegisterDomainDropRequest
	DomainForwardRequest                = types.DomainForwardRequest
	DomainForwardSubDomainRequest       = types.DomainForwardSubDomainRequest
	DeleteDomainForwardSubDomainRequest = types.DeleteDomainForwardSubDomainRequest
	DomainPushRequest                   = types.DomainPushRequest
	CheckTransferAvailabilityRequest    = types.CheckTransferAvailabilityRequest
	RecommendDomainsRequest             = types.RecommendDomainsRequest
	CheckAvailabilityFilterRequest      = types.CheckAvailabilityFilterRequest
)

// ============================================================================
// 域名响应类型
// ============================================================================

type (
	Domain                               = types.Domain
	AvailableDomain                      = types.AvailableDomain
	ListDomainsResponse                  = types.ListDomainsResponse
	GetDomainInfoResponse                = types.GetDomainInfoResponse
	CheckRegisterAvailabilityResponse    = types.CheckRegisterAvailabilityResponse
	RegisterDomainResponse               = types.RegisterDomainResponse
	RenewDomainResponse                  = types.RenewDomainResponse
	DomainLockResponse                   = types.DomainLockResponse
	DomainUnlockResponse                 = types.DomainUnlockResponse
	AddAutoRenewalResponse               = types.AddAutoRenewalResponse
	RemoveAutoRenewalResponse            = types.RemoveAutoRenewalResponse
	WhoisResponse                        = types.WhoisResponse
	DomainInfoReply                      = types.DomainInfoReply
	NameserverEntry                      = types.NameserverEntry
	ContactIDs                           = types.ContactIDs
	TransferDomainResponse               = types.TransferDomainResponse
	RegisterDomainDropResponse           = types.RegisterDomainDropResponse
	RegisterDomainClaimsResponse         = types.RegisterDomainClaimsResponse
	TrademarkClaim                       = types.TrademarkClaim
	TrademarkClaimInfo                   = types.TrademarkClaimInfo
	TrademarkContact                     = types.TrademarkContact
	DomainForwardResponse                = types.DomainForwardResponse
	DomainForwardSubDomainResponse       = types.DomainForwardSubDomainResponse
	DeleteDomainForwardSubDomainResponse = types.DeleteDomainForwardSubDomainResponse
	DomainPushResponse                   = types.DomainPushResponse
	DomainPushResult                     = types.DomainPushResult
	DomainPushStatus                     = types.DomainPushStatus
	CheckTransferAvailabilityResponse    = types.CheckTransferAvailabilityResponse
	TransferAvailableDomain              = types.TransferAvailableDomain
	TransferUnavailableDomain            = types.TransferUnavailableDomain
	RecommendDomainsResponse             = types.RecommendDomainsResponse
	CheckAvailabilityFilterResponse      = types.CheckAvailabilityFilterResponse
	DomainFilter                         = types.DomainFilter
	FilterType                           = types.FilterType
)

// 过滤器类型常量
var (
	FilterByPrice  = types.FilterByPrice
	FilterByTLD    = types.FilterByTLD
	FilterByLength = types.FilterByLength
)
