/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:57:00
 * @FilePath: \go-namesilo\domains\auto_renewal.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// AddAutoRenewal 添加自动续费
// API: addAutoRenewal
// Docs: https://www.namesilo.com/api-reference#domains/add-auto-renewal
func (s *Service) AddAutoRenewal(ctx context.Context, req *AddAutoRenewalRequest) (*AddAutoRenewalResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "addAutoRenewal", params)
	if err != nil {
		return nil, err
	}

	var resp AddAutoRenewalResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}

// RemoveAutoRenewal 移除自动续费
// API: removeAutoRenewal
// Docs: https://www.namesilo.com/api-reference#domains/remove-auto-renewal
func (s *Service) RemoveAutoRenewal(ctx context.Context, req *RemoveAutoRenewalRequest) (*RemoveAutoRenewalResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "removeAutoRenewal", params)
	if err != nil {
		return nil, err
	}

	var resp RemoveAutoRenewalResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
