/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 09:35:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 09:50:32
 * @FilePath: \go-namesilo\nameserver\delete.go
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

// Delete 删除已注册的域名服务器
// API: deleteRegisteredNameServer
// Docs: https://www.namesilo.com/api-reference#nameserver/delete-registered-nameserver
func (s *Service) Delete(ctx context.Context, req *types.DeleteRegisteredNameServerRequest) (*types.DeleteRegisteredNameServerResponse, error) {
	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("current_host", req.CurrentHost).
		Build()

	data, err := s.client.DoRequest(ctx, "deleteRegisteredNameServer", params)
	if err != nil {
		return nil, err
	}

	var resp types.DeleteRegisteredNameServerResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%w", &resp.Reply)
	}

	return &resp, nil
}
