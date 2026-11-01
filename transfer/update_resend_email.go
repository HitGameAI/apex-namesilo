/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:30:00
 * @FilePath: \go-namesilo\transfer\update_resend_email.go
 * @Description: 重新发送域名转移管理员邮件
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package transfer

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// UpdateResendEmail 重新发送域名转移管理员邮件
// 重新发送转移批准邮件到域名的管理员邮箱
// Docs: https://www.namesilo.com/api-reference#transfers/transfer-update-resend-admin-email
func (s *Service) UpdateResendEmail(ctx context.Context, req *TransferUpdateResendEmailRequest) (*TransferUpdateResendEmailResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "transferUpdateResendAdminEmail", params)
	if err != nil {
		return nil, err
	}

	var resp TransferUpdateResendEmailResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
