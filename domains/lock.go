/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:55:15
 * @FilePath: \go-namesilo\domains\lock.go
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

// Lock 锁定域名
// API: domainLock
// Docs: https://www.namesilo.com/api-reference#domains/domain-lock
func (s *Service) Lock(ctx context.Context, req *DomainLockRequest) (*DomainLockResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "domainLock", params)
	if err != nil {
		return nil, err
	}

	var resp DomainLockResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}

// Unlock 解锁域名
// API: domainUnlock
// Docs: https://www.namesilo.com/api-reference#domains/domain-unlock
func (s *Service) Unlock(ctx context.Context, req *DomainUnlockRequest) (*DomainUnlockResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "domainUnlock", params)
	if err != nil {
		return nil, err
	}

	var resp DomainUnlockResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
