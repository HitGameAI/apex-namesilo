/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:00:00
 * @FilePath: \go-namesilo\portfolio\aliases.go
 * @Description: Portfolio 组合管理模块类型别名
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package portfolio

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

	// Portfolio 相关错误
	ErrPortfolioRequired      = namesilo.ErrPortfolioRequired
	ErrPortfolioNotFound      = namesilo.ErrPortfolioNotFound
	ErrPortfolioAlreadyExists = namesilo.ErrPortfolioAlreadyExists
	ErrPortfolioInUse         = namesilo.ErrPortfolioInUse
)

// ============================================================================
// Portfolio 请求类型
// ============================================================================

type (
	PortfolioListRequest            = types.PortfolioListRequest
	PortfolioAddRequest             = types.PortfolioAddRequest
	PortfolioDeleteRequest          = types.PortfolioDeleteRequest
	PortfolioDomainAssociateRequest = types.PortfolioDomainAssociateRequest
)

// ============================================================================
// Portfolio 响应类型
// ============================================================================

type (
	PortfolioListResponse            = types.PortfolioListResponse
	PortfolioAddResponse             = types.PortfolioAddResponse
	PortfolioDeleteResponse          = types.PortfolioDeleteResponse
	PortfolioDomainAssociateResponse = types.PortfolioDomainAssociateResponse
	Portfolio                        = types.Portfolio
)
