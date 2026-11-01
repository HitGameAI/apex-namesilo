/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-06 11:55:05
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 11:55:15
 * @FilePath: \go-namesilo\domains\list.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"
	"strconv"

	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
	"github.com/HitGameAI/apex-toolbox/pkg/mathx"
)

// List 列出域名
// API: listDomains
// Docs: https://www.namesilo.com/api-reference#domains/list-domains
func (s *Service) List(ctx context.Context, req *ListDomainsRequest) (*ListDomainsResponse, error) {
	params := httpx.NewParams().
		Set("page", strconv.Itoa(mathx.IfNotZero(req.Page, DefaultPage))).
		Set("page_size", strconv.Itoa(mathx.IfNotZero(req.PageSize, DefaultPageSize)))

	data, err := s.client.DoRequest(ctx, "listDomains", params.Build())
	if err != nil {
		return nil, err
	}

	var resp ListDomainsResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, fmt.Errorf("%s", resp.Reply.Error())
	}

	return &resp, nil
}
