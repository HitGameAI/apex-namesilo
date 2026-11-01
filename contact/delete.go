/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 22:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:30:00
 * @FilePath: \go-namesilo\contact\delete.go
 * @Description: 删除联系人
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package contact

import (
	"context"

	namesilo "github.com/HitGameAI/apex-namesilo"
	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// DeleteContact 删除联系人配置文件
// 删除账户中的联系人配置文件
// 注意：只能删除非默认且未关联到任何活动域名或订单配置文件的联系人
// Docs: https://www.namesilo.com/api-reference#contact/contact-delete
func (s *Service) DeleteContact(ctx context.Context, req *ContactDeleteRequest) (*ContactDeleteResponse, error) {
	if req.ContactID == "" {
		return nil, ErrContactIDRequired
	}

	params := httpx.NewParams().
		Set("contact_id", req.ContactID).
		Build()

	data, err := s.client.DoRequest(ctx, "contactDelete", params)
	if err != nil {
		return nil, err
	}

	var resp ContactDeleteResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, namesilo.NewAPIError("delete contact", resp.Reply.Error())
	}

	return &resp, nil
}
