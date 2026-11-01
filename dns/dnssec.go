/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 13:20:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 13:20:00
 * @FilePath: \go-namesilo\dns\dnssec.go
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

// ListDNSSecRecords 列出 DNSSEC 记录
// 查看域名的 DNSSEC DS 记录
// Docs: https://www.namesilo.com/api-reference#dns/dns-seclist-records
func (s *Service) ListDNSSecRecords(ctx context.Context, req *DNSSecListRecordsRequest) (*DNSSecListRecordsResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().Set("domain", req.Domain).Build()

	data, err := s.client.DoRequest(ctx, "dnsSecListRecords", params)
	if err != nil {
		return nil, err
	}

	var resp DNSSecListRecordsResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddDNSSecRecord 添加 DNSSEC 记录
// 为域名添加 DNSSEC DS 记录
// Docs: https://www.namesilo.com/api-reference#dns/dns-secadd-records
func (s *Service) AddDNSSecRecord(ctx context.Context, req *DNSSecAddRecordRequest) (*DNSSecAddRecordResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.Digest == "" {
		return nil, ErrDigestRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("digest", req.Digest).
		Set("keyTag", strconv.Itoa(req.KeyTag)).
		Set("digestType", strconv.Itoa(req.DigestType)).
		Set("alg", strconv.Itoa(req.Algorithm)).
		Build()

	data, err := s.client.DoRequest(ctx, "dnsSecAddRecord", params)
	if err != nil {
		return nil, err
	}

	var resp DNSSecAddRecordResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteDNSSecRecord 删除 DNSSEC 记录
// 删除域名的 DNSSEC DS 记录
// Docs: https://www.namesilo.com/api-reference#dns/dns-secdelete-record
func (s *Service) DeleteDNSSecRecord(ctx context.Context, req *DNSSecDeleteRecordRequest) (*DNSSecDeleteRecordResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.Digest == "" {
		return nil, ErrDigestRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("digest", req.Digest).
		Set("keyTag", strconv.Itoa(req.KeyTag)).
		Set("digestType", strconv.Itoa(req.DigestType)).
		Set("alg", strconv.Itoa(req.Algorithm)).
		Build()

	data, err := s.client.DoRequest(ctx, "dnsSecDeleteRecord", params)
	if err != nil {
		return nil, err
	}

	var resp DNSSecDeleteRecordResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
