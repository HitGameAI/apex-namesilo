/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:20:00
 * @FilePath: \go-namesilo\dns\list_records.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package dns

import (
	"context"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// ListRecords 列出 DNS 记录
// 查看域名所有的当前 DNS 记录
// Docs: https://www.namesilo.com/api-reference#dns/dns-list-records
func (s *Service) ListRecords(ctx context.Context, req *DNSListRecordsRequest) (*DNSListRecordsResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Build()

	data, err := s.client.DoRequest(ctx, "dnsListRecords", params)
	if err != nil {
		return nil, err
	}

	var resp DNSListRecordsResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
