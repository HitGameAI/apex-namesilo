/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:55:15
 * @FilePath: \go-namesilo\domains\whois.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Whois 查询域名 Whois 信息
// API: whoisInfo
// Docs: https://www.namesilo.com/api-reference#domains/whois
func (s *Service) Whois(ctx context.Context, req *WhoisRequest) (*WhoisResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "whoisInfo", params)
	if err != nil {
		return nil, err
	}

	var resp WhoisResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%s", resp.Reply.Error())
	}

	return &resp, nil
}
