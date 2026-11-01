/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:58:36
 * @FilePath: \go-namesilo\dns\update_record.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package dns

import (
	"context"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// UpdateRecord 更新 DNS 记录
// 更新现有的 DNS 资源记录
// Docs: https://www.namesilo.com/api-reference#dns/dns-update-record
func (s *Service) UpdateRecord(ctx context.Context, req *DNSUpdateRecordRequest) (*DNSUpdateRecordResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.RecordID == "" {
		return nil, ErrRecordIDRequired
	}
	if req.Host == "" {
		return nil, ErrRRHostRequired
	}
	if req.Value == "" {
		return nil, ErrRRValueRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("rrid", req.RecordID).
		Set("rrhost", req.Host).
		Set("rrvalue", req.Value).
		SetIf(req.Distance > 0, "rrdistance", strconv.Itoa(req.Distance)).
		SetIf(req.TTL > 0, "rrttl", strconv.Itoa(req.TTL)).
		Build()

	data, err := s.client.DoRequest(ctx, "dnsUpdateRecord", params)
	if err != nil {
		return nil, err
	}

	var resp DNSUpdateRecordResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
