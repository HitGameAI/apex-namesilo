/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:00:00
 * @FilePath: \go-namesilo\account\get_balance.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package account

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// GetBalance 获取账户余额
// 查看当前账户资金余额
// Docs: https://www.namesilo.com/api-reference#account/get-account-balance
func (s *Service) GetBalance(ctx context.Context, req *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	params := httpx.NewParams().Build()

	data, err := s.client.DoRequest(ctx, "getAccountBalance", params)
	if err != nil {
		return nil, err
	}

	var response GetAccountBalanceResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
