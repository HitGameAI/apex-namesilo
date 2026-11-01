/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:51:17
 * @FilePath: \go-namesilo\nameserver\modify.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package nameserver

import (
	"context"
	"fmt"

	"github.com/HitGameAI/apex-namesilo/types"
	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Modify 修改已注册的域名服务器
// API: modifyRegisteredNameServer
// Docs: https://www.namesilo.com/api-reference#nameserver/modify-registered-nameserver
func (s *Service) Modify(ctx context.Context, req *types.ModifyRegisteredNameServerRequest) (*types.ModifyRegisteredNameServerResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("current_host", req.CurrentHost).
		Set("new_host", req.NewHost)

	// 添加 IP 地址参数 (ip1, ip2, ..., ip13)
	for i, ip := range req.IPs {
		if i < 13 { // 最多 13 个 IP
			params.Set(fmt.Sprintf("ip%d", i+1), ip)
		}
	}

	data, err := s.client.DoRequest(ctx, "modifyRegisteredNameServer", params.Build())
	if err != nil {
		return nil, err
	}

	var resp types.ModifyRegisteredNameServerResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
