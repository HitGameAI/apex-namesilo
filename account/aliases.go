/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:59:15
 * @FilePath: \go-namesilo\account\aliases.go
 * @Description: 账户相关类型别名
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package account

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
	// Account 相关错误
	ErrInsufficientBalance = namesilo.ErrInsufficientBalance
	ErrInvalidAmount       = namesilo.ErrInvalidAmount
	ErrPaymentFailed       = namesilo.ErrPaymentFailed
	ErrOrderNumberRequired = namesilo.ErrOrderNumberRequired
	ErrDaysCountInvalid    = namesilo.ErrDaysCountInvalid
	ErrAmountInvalid       = namesilo.ErrAmountInvalid
)

// ============================================================================
// 分页默认值
// ============================================================================
var (
	DefaultPage     = types.DefaultPage
	DefaultPageSize = types.DefaultPageSize
)

// ============================================================================
// 价格相关请求类型
// ============================================================================

type GetPricesRequest = types.GetPricesRequest

// ============================================================================
// 价格相关响应类型
// ============================================================================

type (
	GetPricesResponse = types.GetPricesResponse
	PricesReply       = types.PricesReply
	PriceDetail       = types.PriceDetail
)

// ============================================================================
// 账户余额相关类型
// ============================================================================

type (
	GetAccountBalanceRequest  = types.GetAccountBalanceRequest
	GetAccountBalanceResponse = types.GetAccountBalanceResponse
)

// ============================================================================
// 账户资金相关类型
// ============================================================================

type (
	AddAccountFundsRequest  = types.AddAccountFundsRequest
	AddAccountFundsResponse = types.AddAccountFundsResponse
)

// ============================================================================
// 订单相关类型
// ============================================================================

type (
	Order                = types.Order
	OrderDetail          = types.OrderDetail
	ListOrdersRequest    = types.ListOrdersRequest
	ListOrdersResponse   = types.ListOrdersResponse
	OrderDetailsRequest  = types.OrderDetailsRequest
	OrderDetailsResponse = types.OrderDetailsResponse
)

// ============================================================================
// 即将过期域名相关类型
// ============================================================================

type (
	ExpiringDomain               = types.ExpiringDomain
	ListExpiringDomainsRequest   = types.ListExpiringDomainsRequest
	ListExpiringDomainsResponse  = types.ListExpiringDomainsResponse
	CountExpiringDomainsRequest  = types.CountExpiringDomainsRequest
	CountExpiringDomainsResponse = types.CountExpiringDomainsResponse
)
