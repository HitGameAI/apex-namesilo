/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\transfer\retrieve_auth_code.go
 * @Description: 获取域名授权码
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package transfer

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// RetrieveAuthCode 获取域名的授权码(EPP Code)
// 检索当前在您的账户中的域名的授权码
// Docs: https://www.namesilo.com/api-reference#transfers/retrieve-auth-code
func (s *Service) RetrieveAuthCode(ctx context.Context, req *RetrieveAuthCodeRequest) (*RetrieveAuthCodeResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "retrieveAuthCode", params)
	if err != nil {
		return nil, err
	}

	var resp RetrieveAuthCodeResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
