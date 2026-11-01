/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:30:00
 * @FilePath: \go-namesilo\domains\forward_domain.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"net/url"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// ForwardDomain 转发域名
// 配置您的域名使用我们的任一转发选项进行转发
// Docs: https://www.namesilo.com/api-reference#domains/domain-forward
func (s *Service) ForwardDomain(ctx context.Context, req *DomainForwardRequest) (*DomainForwardResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.Protocol == "" {
		return nil, ErrProtocolRequired
	}
	if req.Address == "" {
		return nil, ErrAddressForwardRequired
	}
	if req.Method == "" {
		return nil, ErrMethodRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("protocol", req.Protocol).
		Set("address", req.Address).
		Set("method", req.Method).
		SetIf(req.MetaTitle != "", "meta_title", url.QueryEscape(req.MetaTitle)).
		SetIf(req.MetaDescription != "", "meta_description", url.QueryEscape(req.MetaDescription)).
		SetIf(req.MetaKeywords != "", "meta_keywords", url.QueryEscape(req.MetaKeywords)).
		Build()

	data, err := s.client.DoRequest(ctx, "domainForward", params)
	if err != nil {
		return nil, err
	}

	var response DomainForwardResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
