/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:57:59
 * @FilePath: \go-namesilo\domains\renew.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Renew 续费域名
// 续费您账户中的域名指定年数
// 注意：如果域名处于恢复期，renewDomain 将执行恢复操作。您可以通过账户中的 API Manager 页面覆盖此行为。
// Docs: https://www.namesilo.com/api-reference#domains/renew-domain
func (s *Service) Renew(ctx context.Context, req *RenewDomainRequest) (*RenewDomainResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.Years < 1 || req.Years > 10 {
		return nil, ErrYearsOutOfRange
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("years", strconv.Itoa(req.Years)).
		SetIf(req.PaymentID != "", "payment_id", req.PaymentID).
		SetIf(req.Coupon != "", "coupon", req.Coupon).
		Build()

	data, err := s.client.DoRequest(ctx, "renewDomain", params)
	if err != nil {
		return nil, err
	}

	var resp RenewDomainResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
