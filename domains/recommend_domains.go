/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 13:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 00:56:25
 * @FilePath: \go-namesilo\domains\recommend_domains.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"fmt"
	"strings"

	"github.com/kamalyes/go-toolbox/pkg/mathx"
	"github.com/kamalyes/go-toolbox/pkg/random"
)

// RecommendDomains 推荐域名（基于关键词生成推荐域名列表）
// Docs: https://www.namesilo.com/api-reference#domains/check-register-availability
//
// 根据关键词生成多种变体并检查可用性，同时检查原始关键词的可用性
// 自动过滤超过最大价格的高级域名
func (s *Service) RecommendDomains(ctx context.Context, req *RecommendDomainsRequest) (*RecommendDomainsResponse, error) {
	// 设置默认值
	tlds := mathx.IfEmpty(req.TLDs, DefaultAllowedTLDs)
	includeMatched := req.IncludeMatched // 默认 false，需要显式设置

	// 处理关键字中包含 TLD 的情况（如 12356abc.ai、12356abc.ai.com）
	// 以第一个 . 为界，前半部分作为关键字，后半部分的每段都作为 TLD
	// 仅将白名单（DefaultAllowedTLDs）内的 TLD 加入搜索列表，非白名单 TLD（如 abc.cn 中的 cn）会被忽略，
	// 避免拆分后用非白名单 TLD 生成推荐域名
	keyword := req.Keyword
	if parts := strings.Split(keyword, "."); len(parts) > 1 && parts[0] != "" {
		keyword = parts[0]
		for _, tld := range parts[1:] {
			if tld != "" && mathx.ContainsAnySlice(DefaultAllowedTLDs, tld) && !mathx.ContainsAnySlice(tlds, tld) {
				tlds = append(tlds, tld)
			}
		}
	}

	// 1. 生成推荐域名（基于关键词变体）
	recommendKeywords := random.NewDomainKeywordBuilder(keyword).Generate()
	recommendDomains := random.JoinDomainsWithTLDs(recommendKeywords, tlds)

	// 2. 生成匹配域名（原始关键词）
	var allDomains []string
	var originalDomains []string

	if includeMatched {
		originalDomains = random.JoinDomainsWithTLDs([]string{keyword}, tlds)
		allDomains = append(recommendDomains, originalDomains...)
	} else {
		allDomains = recommendDomains
	}

	// 3. 创建原始域名集合（用于快速查找）
	originalDomainSet := make(map[string]bool)
	for _, domain := range originalDomains {
		originalDomainSet[domain] = true
	}

	// 4. 检查所有域名可用性（不提前截断，让API返回所有结果）
	checkResp, err := s.CheckAvailability(ctx, &CheckRegisterAvailabilityRequest{
		Domains: allDomains,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to check domain availability: %w", err)
	}

	// 5. 创建请求域名集合（防御性过滤，跳过 API 返回的未请求域名）
	requestedDomainSet := make(map[string]bool, len(allDomains))
	for _, d := range allDomains {
		requestedDomainSet[d] = true
	}

	// 6. 分离并过滤域名
	resp := &RecommendDomainsResponse{
		Recommended: make([]AvailableDomain, 0),
		Matched:     make([]AvailableDomain, 0),
	}

	// 防御性过滤不可用域名列表
	for _, d := range checkResp.Reply.Unavailable {
		if requestedDomainSet[d] {
			resp.Unavailable = append(resp.Unavailable, d)
		}
	}

	// 处理可用域名：防御性过滤 + 价格过滤
	for _, domain := range checkResp.Reply.Available {
		// 防御性过滤：跳过 API 返回的未请求域名
		if !requestedDomainSet[domain.Domain] {
			continue
		}
		// 跳过超过最大价格的高级域名
		if req.MaxPrice > 0 && domain.Price > req.MaxPrice {
			continue
		}

		availDomain := AvailableDomain{
			Domain: domain.Domain,
			Price:  domain.Price,
		}
		// 根据域名类型分类
		if originalDomainSet[domain.Domain] {
			resp.Matched = append(resp.Matched, availDomain)
		} else {
			resp.Recommended = append(resp.Recommended, availDomain)
		}
	}

	return resp, nil
}
