/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:00:00
 * @FilePath: \go-namesilo\portfolio\add.go
 * @Description: 添加域名组合
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package portfolio

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Add 创建新的域名组合
// 在账户中添加一个新的域名组合
// Docs: https://www.namesilo.com/api-reference#portfolio/portfolio-add
func (s *Service) Add(ctx context.Context, req *PortfolioAddRequest) (*PortfolioAddResponse, error) {
	if req.Portfolio == "" {
		return nil, ErrPortfolioRequired
	}

	params := httpx.NewParams().
		Set("portfolio", req.Portfolio).
		Build()

	data, err := s.client.DoRequest(ctx, "portfolioAdd", params)
	if err != nil {
		return nil, err
	}

	var resp PortfolioAddResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
