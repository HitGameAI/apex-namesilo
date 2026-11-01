/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:36:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:36:00
 * @FilePath: \go-namesilo\forwarding\forward_domain.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package forwarding

import (
	"context"
	"fmt"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// ForwardDomain 设置域名转发
// API: domainForward
// Docs: https://www.namesilo.com/api-reference#forwarding/forward-domain
func (s *Service) ForwardDomain(ctx context.Context, req *ForwardDomainRequest) (*ForwardDomainResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("protocol", req.Protocol).
		Set("address", req.Address).
		SetIf(req.Method != "", "method", req.Method).
		SetIf(req.IncludePath != "", "include_path", req.IncludePath).
		SetIf(req.Wildcard != "", "wildcard", req.Wildcard).
		Build()

	data, err := s.client.DoRequest(ctx, "domainForward", params)
	if err != nil {
		return nil, err
	}

	var resp ForwardDomainResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
