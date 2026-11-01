/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:37:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:37:00
 * @FilePath: \go-namesilo\forwarding\forward_subdomain.go
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

// ForwardSubdomain 设置子域名转发
// API: domainForwardSubDomain
// Docs: https://www.namesilo.com/api-reference#forwarding/forward-domain-sub-domain
func (s *Service) ForwardSubdomain(ctx context.Context, req *ForwardSubdomainRequest) (*ForwardSubdomainResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("sub_domain", req.SubDomain).
		Set("protocol", req.Protocol).
		Set("address", req.Address).
		SetIf(req.Method != "", "method", req.Method).
		SetIf(req.IncludePath != "", "include_path", req.IncludePath).
		Build()

	data, err := s.client.DoRequest(ctx, "domainForwardSubDomain", params)
	if err != nil {
		return nil, err
	}

	var resp ForwardSubdomainResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
