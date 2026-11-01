/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:00:00
 * @FilePath: \go-namesilo\account\order_details.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package account

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// GetOrderDetails 获取订单详情
// 查看指定订单号的详细信息
// Docs: https://www.namesilo.com/api-reference#account/order-details
func (s *Service) GetOrderDetails(ctx context.Context, req *OrderDetailsRequest) (*OrderDetailsResponse, error) {
	if req.OrderNumber == "" {
		return nil, ErrOrderNumberRequired
	}

	params := httpx.NewParams().
		Set("order_number", req.OrderNumber).
		Build()

	data, err := s.client.DoRequest(ctx, "orderDetails", params)
	if err != nil {
		return nil, err
	}

	var response OrderDetailsResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
