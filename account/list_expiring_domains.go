/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:00:00
 * @FilePath: \go-namesilo\account\list_expiring_domains.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package account

import (
	"context"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
	"github.com/HitGameAI/apex-toolbox/pkg/mathx"
)

// ListExpiringDomains 列出即将过期的域名
// 返回即将过期的域名列表
// Docs: https://www.namesilo.com/api-reference#account/list-expiring-domains
func (s *Service) ListExpiringDomains(ctx context.Context, req *ListExpiringDomainsRequest) (*ListExpiringDomainsResponse, error) {
	if req.DaysCount <= 0 {
		return nil, ErrDaysCountInvalid
	}

	params := httpx.NewParams().
		Set("daysCount", strconv.Itoa(req.DaysCount)).
		Set("page", strconv.Itoa(mathx.IfNotZero(req.Page, DefaultPage))).
		Set("page_size", strconv.Itoa(mathx.IfNotZero(req.PageSize, DefaultPageSize))).
		Build()

	data, err := s.client.DoRequest(ctx, "listExpiringDomains", params)
	if err != nil {
		return nil, err
	}

	var response ListExpiringDomainsResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
