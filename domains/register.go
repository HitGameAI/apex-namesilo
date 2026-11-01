/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:08:25
 * @FilePath: \go-namesilo\domains\register.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Register 注册域名
// API: registerDomain
// Docs: https://www.namesilo.com/api-reference#domains/register-domain
func (s *Service) Register(ctx context.Context, req *RegisterDomainRequest) (*RegisterDomainResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("years", strconv.Itoa(req.Years)).
		SetIf(req.AutoRenew, "auto_renew", "1").
		SetIf(req.Private, "private", "1").
		SetIf(req.Coupon != "", "coupon", req.Coupon).
		Build()

	data, err := s.client.DoRequest(ctx, "registerDomain", params)
	if err != nil {
		return nil, err
	}

	var resp RegisterDomainResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("register domain failed: %s", resp.Reply.Error())
	}

	return &resp, nil
}
