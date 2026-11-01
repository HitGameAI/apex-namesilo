/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:39:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:39:00
 * @FilePath: \go-namesilo\forwarding\list_email_forwards.go
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

// ListEmailForwards 列出邮件转发
// API: listEmailForwards
// Docs: https://www.namesilo.com/api-reference#forwarding/list-email-forwards
func (s *Service) ListEmailForwards(ctx context.Context, req *ListEmailForwardsRequest) (*ListEmailForwardsResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "listEmailForwards", params)
	if err != nil {
		return nil, err
	}

	var resp ListEmailForwardsResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if resp.Reply.Code != 300 && resp.Reply.Code != 0 {
		return nil, fmt.Errorf("API error: code=%d, detail=%s", resp.Reply.Code, resp.Reply.Detail)
	}

	return &resp, nil
}
