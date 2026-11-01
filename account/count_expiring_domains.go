/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:00:00
 * @FilePath: \go-namesilo\account\count_expiring_domains.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package account

import (
	"context"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// CountExpiringDomains 统计即将过期的域名数量
// 返回即将过期的域名数量
// Docs: https://www.namesilo.com/api-reference#account/count-expiring-domains
func (s *Service) CountExpiringDomains(ctx context.Context, req *CountExpiringDomainsRequest) (*CountExpiringDomainsResponse, error) {
	if req.DaysCount <= 0 {
		return nil, ErrDaysCountInvalid
	}

	params := httpx.NewParams().
		Set("daysCount", strconv.Itoa(req.DaysCount)).
		Build()

	data, err := s.client.DoRequest(ctx, "countExpiringDomains", params)
	if err != nil {
		return nil, err
	}

	var response CountExpiringDomainsResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
