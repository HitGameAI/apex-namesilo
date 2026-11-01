/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:00:00
 * @FilePath: \go-namesilo\portfolio\list.go
 * @Description: 列出所有域名组合
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package portfolio

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// List 列出所有域名组合
// 返回账户中所有的域名组合列表
// Docs: https://www.namesilo.com/api-reference#portfolio/portfolio-list
func (s *Service) List(ctx context.Context, req *PortfolioListRequest) (*PortfolioListResponse, error) {
	params := httpx.NewParams().Build()

	data, err := s.client.DoRequest(ctx, "portfolioList", params)
	if err != nil {
		return nil, err
	}

	var resp PortfolioListResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
