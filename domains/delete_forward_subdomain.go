/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:15:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:15:00
 * @FilePath: \go-namesilo\domains\delete_forward_subdomain.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// DeleteForwardSubDomain 删除子域名转发
// 删除子域名的转发配置
// Docs: https://www.namesilo.com/api-reference#domains/domain-forward-sub-domain-delete
func (s *Service) DeleteForwardSubDomain(ctx context.Context, req *DeleteDomainForwardSubDomainRequest) (*DeleteDomainForwardSubDomainResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.SubDomain == "" {
		return nil, ErrSubDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("sub_domain", req.SubDomain).
		Build()

	data, err := s.client.DoRequest(ctx, "domainForwardSubDomainDelete", params)
	if err != nil {
		return nil, err
	}

	var response DeleteDomainForwardSubDomainResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
