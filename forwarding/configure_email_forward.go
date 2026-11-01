/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:40:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:40:00
 * @FilePath: \go-namesilo\forwarding\configure_email_forward.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package forwarding

import (
	"context"
	"fmt"
	"strings"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// ConfigureEmailForward 配置邮件转发
// API: configureEmailForward
// Docs: https://www.namesilo.com/api-reference#forwarding/configure-email-forwards
func (s *Service) ConfigureEmailForward(ctx context.Context, req *ConfigureEmailForwardRequest) (*ConfigureEmailForwardResponse, error) {
	if len(req.Forward) > 5 {
		return nil, ErrForwardListExceedLimit
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("email", req.Email)

	// 添加转发目标列表（最多 5 个）
	for i, addr := range req.Forward {
		params.Set(fmt.Sprintf("forward%d", i+1), strings.TrimSpace(addr))
	}

	data, err := s.client.DoRequest(ctx, "configureEmailForward", params.Build())
	if err != nil {
		return nil, err
	}

	var resp ConfigureEmailForwardResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
