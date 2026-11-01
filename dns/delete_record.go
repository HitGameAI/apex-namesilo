/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:20:00
 * @FilePath: \go-namesilo\dns\delete_record.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package dns

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// DeleteRecord 删除 DNS 记录
// 删除现有的 DNS 资源记录
// Docs: https://www.namesilo.com/api-reference#dns/dns-delete-record
func (s *Service) DeleteRecord(ctx context.Context, req *DNSDeleteRecordRequest) (*DNSDeleteRecordResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.RRID == "" {
		return nil, ErrRecordIDRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("rrid", req.RRID).
		Build()

	data, err := s.client.DoRequest(ctx, "dnsDeleteRecord", params)
	if err != nil {
		return nil, err
	}

	var resp DNSDeleteRecordResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
