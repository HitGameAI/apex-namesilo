/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:30:00
 * @FilePath: \go-namesilo\domains\transfer.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"
	"net/url"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// Transfer 转移域名
// 从其他注册商将域名转移到您的账户
// 注意：您必须至少有一张已验证的信用卡或足够的账户资金才能使用此命令
// Docs: https://www.namesilo.com/api-reference#domains/transfer-domain
func (s *Service) Transfer(ctx context.Context, req *TransferDomainRequest) (*TransferDomainResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		SetIf(req.Auth != "", "auth", url.QueryEscape(req.Auth)).
		SetIf(req.PaymentID != "", "payment_id", req.PaymentID).
		SetIf(req.Private, "private", "1").
		SetIf(req.AutoRenew, "auto_renew", "1").
		SetIf(req.Portfolio != "", "portfolio", url.QueryEscape(req.Portfolio)).
		SetIf(req.Coupon != "", "coupon", req.Coupon).
		SetIf(req.ContactID != "", "contact_id", req.ContactID)

	// 添加 WHOIS 联系人信息
	if req.FirstName != "" {
		params.Set("fn", url.QueryEscape(req.FirstName))
	}
	if req.LastName != "" {
		params.Set("ln", url.QueryEscape(req.LastName))
	}
	if req.Address != "" {
		params.Set("ad", url.QueryEscape(req.Address))
	}
	if req.City != "" {
		params.Set("cy", url.QueryEscape(req.City))
	}
	if req.State != "" {
		params.Set("st", url.QueryEscape(req.State))
	}
	if req.Zip != "" {
		params.Set("zp", url.QueryEscape(req.Zip))
	}
	if req.Country != "" {
		params.Set("ct", url.QueryEscape(req.Country))
	}
	if req.Email != "" {
		params.Set("em", url.QueryEscape(req.Email))
	}
	if req.Phone != "" {
		params.Set("ph", url.QueryEscape(req.Phone))
	}
	if req.Nickname != "" {
		params.Set("nn", url.QueryEscape(req.Nickname))
	}
	if req.Company != "" {
		params.Set("cp", url.QueryEscape(req.Company))
	}
	if req.Address2 != "" {
		params.Set("ad2", url.QueryEscape(req.Address2))
	}
	if req.Fax != "" {
		params.Set("fx", url.QueryEscape(req.Fax))
	}
	if req.USNexusCategory != "" {
		params.Set("usnc", url.QueryEscape(req.USNexusCategory))
	}
	if req.USAppPurpose != "" {
		params.Set("usap", url.QueryEscape(req.USAppPurpose))
	}

	// 添加域名服务器
	for i, ns := range req.NS {
		if i >= 13 {
			break // 最多13个域名服务器
		}
		params.Set(fmt.Sprintf("ns%d", i+1), ns)
	}

	data, err := s.client.DoRequest(ctx, "transferDomain", params.Build())
	if err != nil {
		return nil, err
	}

	var response TransferDomainResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
