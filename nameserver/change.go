/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:35:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:50:15
 * @FilePath: \go-namesilo\nameserver\change.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package nameserver

import (
	"context"
	"fmt"

	"github.com/HitGameAI/apex-namesilo/types"
	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Change 更改域名服务器
// API: changeNameServers
// Docs: https://www.namesilo.com/api-reference#nameserver/change-nameserver
func (s *Service) Change(ctx context.Context, req *types.ChangeNameServersRequest) (*types.ChangeNameServersResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain)

	// 添加域名服务器参数 (ns1, ns2, ..., ns13)
	for i, ns := range req.Nameservers {
		if i < 13 { // 最多 13 个域名服务器
			params.Set(fmt.Sprintf("ns%d", i+1), ns)
		}
	}

	data, err := s.client.DoRequest(ctx, "changeNameServers", params.Build())
	if err != nil {
		return nil, err
	}

	var resp types.ChangeNameServersResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
