/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:30:00
 * @FilePath: \go-namesilo\domains\check_transfer_availability.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"strings"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// CheckTransferAvailability 检查域名转移可用性
// 检查指定的域名是否可以转移到您的 NameSilo 账户
// Docs: https://www.namesilo.com/api-reference#domains/check-transfer-availability
func (s *Service) CheckTransferAvailability(ctx context.Context, req *CheckTransferAvailabilityRequest) (*CheckTransferAvailabilityResponse, error) {
	if len(req.Domains) == 0 {
		return nil, ErrDomainsRequired
	}
	if len(req.Domains) > 200 {
		return nil, ErrDomainsExceedLimit
	}

	params := httpx.NewParams().
		Set("domains", strings.Join(req.Domains, ",")).
		Build()

	data, err := s.client.DoRequest(ctx, "checkTransferAvailability", params)
	if err != nil {
		return nil, err
	}

	var response CheckTransferAvailabilityResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
