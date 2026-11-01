/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:00:00
 * @FilePath: \go-namesilo\account\add_funds.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package account

import (
	"context"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// AddFunds 添加账户资金
// 增加您的 NameSilo 账户资金余额
// Docs: https://www.namesilo.com/api-reference#account/add-account-funds
func (s *Service) AddFunds(ctx context.Context, req *AddAccountFundsRequest) (*AddAccountFundsResponse, error) {
	if req.Amount <= 0 {
		return nil, ErrAmountInvalid
	}

	params := httpx.NewParams().
		Set("amount", strconv.FormatFloat(req.Amount, 'f', 2, 64)).
		SetIf(req.PaymentID != "", "payment_id", req.PaymentID).
		Build()

	data, err := s.client.DoRequest(ctx, "addAccountFunds", params)
	if err != nil {
		return nil, err
	}

	var response AddAccountFundsResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
