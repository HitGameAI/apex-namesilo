/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:55:15
 * @FilePath: \go-namesilo\domains\check_availability.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"
	"strings"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// CheckAvailability 检查域名注册可用性
// API: checkRegisterAvailability
// Docs: https://www.namesilo.com/api-reference#domains/check-register-availability
func (s *Service) CheckAvailability(ctx context.Context, req *CheckRegisterAvailabilityRequest) (*CheckRegisterAvailabilityResponse, error) {
	params := httpx.NewParams().
		Set("domains", strings.Join(req.Domains, ",")).
		Build()

	data, err := s.client.DoRequest(ctx, "checkRegisterAvailability", params)
	if err != nil {
		return nil, err
	}

	var resp CheckRegisterAvailabilityResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%s", resp.Reply.Error())
	}

	return &resp, nil
}
