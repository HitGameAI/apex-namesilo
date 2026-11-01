/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:15:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:15:00
 * @FilePath: \go-namesilo\forwarding\delete_email_forward.go
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

// DeleteEmailForward 删除邮件转发
// API: deleteEmailForward
// Docs: https://www.namesilo.com/api-reference#forwarding/delete-email-forwards
func (s *Service) DeleteEmailForward(ctx context.Context, req *DeleteEmailForwardRequest) (*DeleteEmailForwardResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("email", req.Email).
		Build()

	data, err := s.client.DoRequest(ctx, "deleteEmailForward", params)
	if err != nil {
		return nil, err
	}

	var resp DeleteEmailForwardResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
