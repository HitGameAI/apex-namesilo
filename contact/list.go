/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 22:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:30:00
 * @FilePath: \go-namesilo\contact\list.go
 * @Description: 列出联系人
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package contact

import (
	"context"
	"fmt"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// ListContacts 列出联系人
// 获取账户中所有联系人配置文件的列表
// 注意：每个请求最多返回1000个联系人，使用offset参数浏览数据集
// Docs: https://www.namesilo.com/api-reference#contact/contact-list
func (s *Service) ListContacts(ctx context.Context, req *ContactListRequest) (*ContactListResponse, error) {
	params := httpx.NewParams().
		SetIf(req.ContactID != "", "contact_id", req.ContactID).
		SetIf(req.Offset > 0, "offset", strconv.Itoa(req.Offset)).
		Build()

	data, err := s.client.DoRequest(ctx, "contactList", params)
	if err != nil {
		return nil, err
	}

	var resp ContactListResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%s", resp.Reply.Error())
	}

	return &resp, nil
}
