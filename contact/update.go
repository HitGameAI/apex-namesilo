/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-16 22:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-16 22:33:49
 * @FilePath: \go-namesilo\contact\update.go
 * @Description: 更新联系人
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package contact

import (
	"context"

	namesilo "github.com/HitGameAI/apex-namesilo"
	"github.com/HitGameAI/apex-toolbox/pkg/httpx"
)

// UpdateContact 更新联系人配置文件
// 更新账户中的联系人配置文件
// 对联系人配置文件所做的任何更改都会更新与该联系人配置文件关联的所有域的注册记录
// Docs: https://www.namesilo.com/api-reference#contact/contact-update
func (s *Service) UpdateContact(ctx context.Context, req *ContactUpdateRequest) (*ContactUpdateResponse, error) {
	if req.ContactID == "" {
		return nil, ErrContactIDRequired
	}
	if req.FirstName == "" {
		return nil, ErrFirstNameRequired
	}
	if req.LastName == "" {
		return nil, ErrLastNameRequired
	}
	if req.Address == "" {
		return nil, ErrAddressRequired
	}
	if req.City == "" {
		return nil, ErrCityRequired
	}
	if req.State == "" {
		return nil, ErrStateRequired
	}
	if req.Zip == "" {
		return nil, ErrZipRequired
	}
	if req.Country == "" {
		return nil, ErrCountryRequired
	}
	if req.Email == "" {
		return nil, ErrEmailRequired
	}
	if req.Phone == "" {
		return nil, ErrPhoneRequired
	}

	params := httpx.NewParams().
		Set("contact_id", req.ContactID).
		Set("fn", req.FirstName).
		Set("ln", req.LastName).
		Set("ad", req.Address).
		Set("cy", req.City).
		Set("st", req.State).
		Set("zp", req.Zip).
		Set("ct", req.Country).
		Set("em", req.Email).
		Set("ph", req.Phone).
		SetIf(req.Nickname != "", "nn", req.Nickname).
		SetIf(req.Company != "", "cp", req.Company).
		SetIf(req.Address2 != "", "ad2", req.Address2).
		SetIf(req.Fax != "", "fx", req.Fax).
		SetIf(req.USNC != "", "usnc", req.USNC).
		SetIf(req.USAP != "", "usap", req.USAP).
		SetIf(req.CALF != "", "calf", req.CALF).
		SetIf(req.CALN != "", "caln", req.CALN).
		SetIf(req.CAAG != "", "caag", req.CAAG).
		SetIf(req.CAWD != "", "cawd", req.CAWD).
		SetIf(req.EUCS != "", "eucs", req.EUCS).
		Build()

	data, err := s.client.DoRequest(ctx, "contactUpdate", params)
	if err != nil {
		return nil, err
	}

	var resp ContactUpdateResponse
	if err := s.client.ParseResponse(data, &resp); err != nil {
		return nil, err
	}

	if !resp.Reply.Success() {
		return nil, namesilo.NewAPIError("update contact", resp.Reply.Error())
	}

	return &resp, nil
}
