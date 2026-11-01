/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:30:00
 * @FilePath: \go-namesilo\domains\push.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Push 推送域名
// 将域名推送给其他用户
// Docs: https://www.namesilo.com/api-reference#domains/domain-push
func (s *Service) Push(ctx context.Context, req *DomainPushRequest) (*DomainPushResponse, error) {
	if req.RecipientLogin == "" {
		return nil, ErrRecipientLoginRequired
	}
	if len(req.Domains) == 0 {
		return nil, ErrDomainsRequired
	}

	params := httpx.NewParams().
		Set("recipientLogin", req.RecipientLogin)

	// 添加域名列表
	for _, domain := range req.Domains {
		params.Set("domains[]", domain)
	}

	data, err := s.client.DoRequest(ctx, "domainPush", params.Build())
	if err != nil {
		return nil, err
	}

	var response DomainPushResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
