/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:35:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:50:15
 * @FilePath: \go-namesilo\nameserver\add.go
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

// Add 添加已注册的域名服务器
// API: addRegisteredNameServer
// Docs: https://www.namesilo.com/api-reference#nameserver/add-registered-nameserver
func (s *Service) Add(ctx context.Context, req *types.AddRegisteredNameServerRequest) (*types.AddRegisteredNameServerResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("new_host", req.NewHost)

	// 检查 IP 地址数量限制
	if len(req.IPs) > 13 {
		return nil, ErrIPAddressExceedLimit
	}
	// 添加 IP 地址参数 (ip1, ip2, ..., ip13)
	for i, ip := range req.IPs {
		params.Set(fmt.Sprintf("ip%d", i+1), ip)
	}

	data, err := s.client.DoRequest(ctx, "addRegisteredNameServer", params.Build())
	if err != nil {
		return nil, err
	}

	var resp types.AddRegisteredNameServerResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
