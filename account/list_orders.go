/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:00:00
 * @FilePath: \go-namesilo\account\list_orders.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package account

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// ListOrders 列出订单
// 返回完整的账户订单历史
// Docs: https://www.namesilo.com/api-reference#account/list-orders
func (s *Service) ListOrders(ctx context.Context, req *ListOrdersRequest) (*ListOrdersResponse, error) {
	params := httpx.NewParams().
		SetIf(req.DateFrom != "", "date_from", req.DateFrom).
		SetIf(req.DateTo != "", "date_to", req.DateTo).
		Build()

	data, err := s.client.DoRequest(ctx, "listOrders", params)
	if err != nil {
		return nil, err
	}

	var response ListOrdersResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
