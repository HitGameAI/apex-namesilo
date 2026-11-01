/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:38:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:38:00
 * @FilePath: \go-namesilo\forwarding\delete_forward.go
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

// DeleteForward 删除域名或子域名转发
// API: domainForwardSubDomainDelete
// Docs: https://www.namesilo.com/api-reference#forwarding/domain-forward-subdomain-delete
func (s *Service) DeleteForward(ctx context.Context, req *DeleteForwardRequest) (*DeleteForwardResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		SetIf(req.SubDomain != "", "sub_domain", req.SubDomain).
		Build()

	data, err := s.client.DoRequest(ctx, "domainForwardSubDomainDelete", params)
	if err != nil {
		return nil, err
	}

	var resp DeleteForwardResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
