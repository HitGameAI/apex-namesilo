/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\transfer\check_transfer_status.go
 * @Description: 检查域名转移状态
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package transfer

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// CheckTransferStatus 检查域名转移状态
// 检查转入您账户的域名转移的当前状态
// Docs: https://www.namesilo.com/api-reference#transfers/check-transfer-status
func (s *Service) CheckTransferStatus(ctx context.Context, req *CheckTransferStatusRequest) (*CheckTransferStatusResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "checkTransferStatus", params)
	if err != nil {
		return nil, err
	}

	var resp CheckTransferStatusResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
