/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:35:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:08:51
 * @FilePath: \go-namesilo\nameserver\list.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package nameserver

import (
	"context"
	"fmt"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// List 列出已注册的域名服务器
// API: listRegisteredNameServers
// Docs: https://www.namesilo.com/api-reference#nameserver/list-registered-nameservers
func (s *Service) List(ctx context.Context, req *ListRegisteredNameServersRequest) (*ListRegisteredNameServersResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "listRegisteredNameServers", params)
	if err != nil {
		return nil, err
	}

	var resp ListRegisteredNameServersResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%s", resp.Reply.Error())
	}

	return &resp, nil
}
