/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 15:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 15:30:00
 * @FilePath: \go-namesilo\domains\register_drop.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// RegisterDrop 使用 Drop-Catching 注册域名
// 使用我们的免费 drop-catching API 流程注册指定年数和指定属性的新域名
// 注意：
// - 此 API 功能只能在 10:45am PT 和 12:15pm PT 之间使用
// - 唯一的支付选项是使用 NameSilo 账户资金（不支持信用卡或 PayPal）
// - Drop-catching 目前仅限于 .com 和 .net 域名
// - 所有 WHOIS 角色将使用您的账户默认联系人资料
// - 将使用我们的默认域名服务器
// - 注意：您必须使用 /apibatch 进行这些操作
// Docs: https://www.namesilo.com/api-reference#domains/register-domain-drop
func (s *Service) RegisterDrop(ctx context.Context, req *RegisterDomainDropRequest) (*RegisterDomainDropResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.Years < 1 || req.Years > 10 {
		return nil, ErrYearsOutOfRange
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		Set("years", strconv.Itoa(req.Years)).
		SetIf(req.Private, "private", "1").
		SetIf(req.AutoRenew, "auto_renew", "1").
		Build()

	// 注意：实际使用时应该使用 /apibatch 端点
	data, err := s.client.DoRequest(ctx, "registerDomainDrop", params)
	if err != nil {
		return nil, err
	}

	var response RegisterDomainDropResponse
	if err := s.client.ParseResponse(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
