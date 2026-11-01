/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\transfer\update_change_epp_code.go
 * @Description: 更改域名转移的 EPP 授权码
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package transfer

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// UpdateChangeEPPCode 更改待转移域名的 EPP 授权码
// 更新正在进行转移的域名的 EPP 授权码
// Docs: https://www.namesilo.com/api-reference#transfers/transfer-update-change-epp-code
func (s *Service) UpdateChangeEPPCode(ctx context.Context, req *TransferUpdateChangeEPPCodeRequest) (*TransferUpdateChangeEPPCodeResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.EPPCode == "" {
		return nil, ErrAuthCodeRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("auth", req.EPPCode).
		Build()

	data, err := s.client.DoRequest(ctx, "transferUpdateChangeEPPCode", params)
	if err != nil {
		return nil, err
	}

	var resp TransferUpdateChangeEPPCodeResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
