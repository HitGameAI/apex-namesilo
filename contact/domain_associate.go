/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 22:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:30:00
 * @FilePath: \go-namesilo\contact\domain_associate.go
 * @Description: 关联联系人到域名
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package contact

import (
	"context"

	"github.com/HitGameAI/apex-namesilo"
	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// AssociateContactToDomain 关联联系人配置文件到域名
// 将联系人配置文件的数据分配给域名的角色
// 至少需要提供一个角色的联系人ID（Registrant、Administrative、Billing 或 Technical）
// Docs: https://www.namesilo.com/api-reference#contact/contact-domain-associate
func (s *Service) AssociateContactToDomain(ctx context.Context, req *ContactDomainAssociateRequest) (*ContactDomainAssociateResponse, error) {
	if req.Domain == "" {
		return nil, ErrDomainRequired
	}
	if req.Registrant == "" && req.Administrative == "" && req.Billing == "" && req.Technical == "" {
		return nil, ErrContactRoleRequired
	}

	params := httpx.NewParams().
		Set("domain", req.Domain).
		SetIf(req.Registrant != "", "registrant", req.Registrant).
		SetIf(req.Administrative != "", "administrative", req.Administrative).
		SetIf(req.Billing != "", "billing", req.Billing).
		SetIf(req.Technical != "", "technical", req.Technical).
		Build()

	data, err := s.client.DoRequest(ctx, "contactDomainAssociate", params)
	if err != nil {
		return nil, err
	}

	var resp ContactDomainAssociateResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, namesilo.NewAPIError("associate contact to domain", resp.Reply.Error())
	}

	return &resp, nil
}
