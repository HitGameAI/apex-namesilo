/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:00:00
 * @FilePath: \go-namesilo\portfolio\delete.go
 * @Description: 删除域名组合
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package portfolio

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Delete 删除域名组合
// 从账户中删除指定的域名组合
// Docs: https://www.namesilo.com/api-reference#portfolio/portfolio-delete
func (s *Service) Delete(ctx context.Context, req *PortfolioDeleteRequest) (*PortfolioDeleteResponse, error) {
	if req.Portfolio == "" {
		return nil, ErrPortfolioRequired
	}

	params := httpx.NewParams().
		Set("portfolio", req.Portfolio).
		Build()

	data, err := s.client.DoRequest(ctx, "portfolioDelete", params)
	if err != nil {
		return nil, err
	}

	var resp PortfolioDeleteResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
