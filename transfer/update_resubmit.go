/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\transfer\update_resubmit.go
 * @Description: 重新提交域名转移到注册局
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package transfer

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// UpdateResubmit 重新提交域名转移到注册局
// 当域名转移被拒绝时,可以使用此方法重新提交转移请求
// Docs: https://www.namesilo.com/api-reference#transfers/transfer-update-resubmit-registry
func (s *Service) UpdateResubmit(ctx context.Context, req *TransferUpdateResubmitRequest) (*TransferUpdateResubmitResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "transferUpdateResubmitToRegistry", params)
	if err != nil {
		return nil, err
	}

	var resp TransferUpdateResubmitResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
