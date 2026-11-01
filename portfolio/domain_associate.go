/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:00:00
 * @FilePath: \go-namesilo\portfolio\domain_associate.go
 * @Description: 关联域名到组合
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package portfolio

import (
	"context"
	"strings"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// DomainAssociate 将域名关联到组合
// 将一个或多个域名添加到指定的组合中
// Docs: https://www.namesilo.com/api-reference#portfolio/portfolio-domain-associate
func (s *Service) DomainAssociate(ctx context.Context, req *PortfolioDomainAssociateRequest) (*PortfolioDomainAssociateResponse, error) {
	if len(req.Domains) == 0 {
		return nil, ErrDomainRequired
	}
	if req.Portfolio == "" {
		return nil, ErrPortfolioRequired
	}

	params := httpx.NewParams().
		Set("domains", strings.Join(req.Domains, ",")).
		Set("portfolio", req.Portfolio).
		Build()

	data, err := s.client.DoRequest(ctx, "portfolioDomainAssociate", params)
	if err != nil {
		return nil, err
	}

	var resp PortfolioDomainAssociateResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
