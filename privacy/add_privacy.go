/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:00:00
 * @FilePath: \go-namesilo\privacy\add_privacy.go
 * @Description: 添加域名隐私保护
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package privacy

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// AddPrivacy 为域名添加隐私保护
// 为指定域名添加 WHOIS 隐私保护服务
// Docs: https://www.namesilo.com/api-reference#privacy/add-privacy
func (s *Service) AddPrivacy(ctx context.Context, req *AddPrivacyRequest) (*AddPrivacyResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "addPrivacy", params)
	if err != nil {
		return nil, err
	}

	var resp AddPrivacyResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
