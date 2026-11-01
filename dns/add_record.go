/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:58:09
 * @FilePath: \go-namesilo\dns\add_record.go
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

// AddRecord 添加 DNS 记录
// 为指定域名添加新的 DNS 资源记录
// Docs: https://www.namesilo.com/api-reference#dns/dns-add-record
func (s *Service) AddRecord(ctx context.Context, req *DNSAddRecordRequest) (*DNSAddRecordResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.Type == "" {
		return nil, ErrRRTypeRequired
	}
	if req.Host == "" {
		return nil, ErrRRHostRequired
	}
	if req.Value == "" {
		return nil, ErrRRValueRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("rrtype", req.Type).
		Set("rrhost", req.Host).
		Set("rrvalue", req.Value).
		SetIf(req.Distance > 0, "rrdistance", strconv.Itoa(req.Distance)).
		SetIf(req.TTL > 0, "rrttl", strconv.Itoa(req.TTL)).
		Build()

	data, err := s.client.DoRequest(ctx, "dnsAddRecord", params)
	if err != nil {
		return nil, err
	}

	var resp DNSAddRecordResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
