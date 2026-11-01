/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 01:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:30:00
 * @FilePath: \go-namesilo\domains\recommend_domains_test.go
 * @Description:
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */
package domains

import (
	"context"
	"testing"

	"github.com/HitGameAI/apex-namesilo/client"
	"github.com/stretchr/testify/assert"
)

func TestRecommendDomains(t *testing.T) {
	// 创建 Mock 客户端
	mockClient := client.NewMockClient()

	// Mock 数据：所有域名均为 example + TLD 格式，便于防御性过滤
	// example.xyz 标记为 premium 且高价，用于验证价格过滤
	availableContent := `<available>
	<domain price="8.99" premium="0">example.com</domain>
	<domain price="12.99" premium="0">example.net</domain>
	<domain price="1.99" premium="0">example.cc</domain>
	<domain price="15.99" premium="1">example.xyz</domain>
</available>
<unavailable>
	<domain>example.me</domain>
	<domain>example.top</domain>
</unavailable>`

	// 设置 Mock 响应
	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", availableContent))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("默认参数推荐", func(t *testing.T) {
		// 不包含匹配域名时，仅生成随机变体域名
		// Mock 不包含随机变体，防御性过滤后结果为空
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword: "example",
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		assert.GreaterOrEqual(t, len(resp.Recommended), 0, "推荐域名数应该 >= 0")
		assert.Len(t, resp.Matched, 0, "默认不包含匹配域名")
		assert.Len(t, resp.Unavailable, 0, "防御性过滤后不可用域名为空")
	})

	t.Run("包含匹配域名", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 应该包含匹配域名（原始关键词 + TLD）
		totalDomains := len(resp.Recommended) + len(resp.Matched)
		assert.Greater(t, totalDomains, 0, "应该有可用域名")

		// 验证匹配域名格式（应该是 example.xxx）
		for _, domain := range resp.Matched {
			assert.Contains(t, domain.Domain, "example", "匹配域名应该包含关键词")
		}

		// 防御性过滤后，不可用域名仅为请求中存在的
		assert.Len(t, resp.Unavailable, 2, "应该有 2 个不可用域名")
	})

	t.Run("自定义TLD", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			TLDs:           []string{"com", "net"},
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证所有域名都是指定的 TLD（防御性过滤保证了这一点）
		allDomains := append(resp.Recommended, resp.Matched...)
		for _, domain := range allDomains {
			tld := extractTLD(domain.Domain)
			assert.Contains(t, []string{"com", "net"}, tld, "域名 %s 的 TLD 应该是 com 或 net", domain.Domain)
		}
	})

	t.Run("自定义最大价格", func(t *testing.T) {
		maxPrice := 5.0
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			MaxPrice:       maxPrice,
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证所有域名价格都在限制内
		for _, domain := range resp.Recommended {
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
		}
		for _, domain := range resp.Matched {
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
		}
	})

	t.Run("组合参数", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			TLDs:           []string{"com", "net"},
			MaxPrice:       10.0,
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证所有条件
		allDomains := append(resp.Recommended, resp.Matched...)
		for _, domain := range allDomains {
			// 验证价格
			assert.LessOrEqual(t, domain.Price, 10.0, "域名 %s 价格应该 <= 10.0", domain.Domain)

			// 验证 TLD
			tld := extractTLD(domain.Domain)
			assert.Contains(t, []string{"com", "net"}, tld, "域名 %s 的 TLD 应该是 com 或 net", domain.Domain)
		}

		// 验证总数量限制
		assert.LessOrEqual(t, len(allDomains), 10, "总域名数应该 <= 10")
	})

	t.Run("过滤高价域名", func(t *testing.T) {
		maxPrice := 3.0
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			MaxPrice:       maxPrice,
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 确保没有高价域名
		allDomains := append(resp.Recommended, resp.Matched...)
		for _, domain := range allDomains {
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
		}
	})
}

func TestRecommendDomainsKeywordWithTLD(t *testing.T) {
	mockClient := client.NewMockClient()

	// Mock 数据：包含 example + 各种 TLD 的域名（dev 不在 DefaultAllowedTLDs 白名单中）
	availableContent := `<available>
	<domain price="8.99" premium="0">example.com</domain>
	<domain price="1.99" premium="0">example.dev</domain>
</available>
<unavailable>
	<domain>example.cc</domain>
</unavailable>`

	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", availableContent))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("关键字包含非白名单TLD", func(t *testing.T) {
		// "example.dev" → keyword="example"，dev 不在 DefaultAllowedTLDs 白名单中，会被忽略不加入 TLD 列表
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example.dev",
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// example.dev 不应该出现在任何结果中（dev 不在白名单，未被加入 TLD 列表）
		for _, domain := range resp.Matched {
			assert.NotEqual(t, "example.dev", domain.Domain, "不应该包含 example.dev")
		}
		for _, domain := range resp.Recommended {
			assert.NotEqual(t, "example.dev", domain.Domain, "不应该包含 example.dev")
		}

		// example.com 应该在匹配结果中（com 在白名单中）
		foundCom := false
		for _, domain := range resp.Matched {
			if domain.Domain == "example.com" {
				foundCom = true
				break
			}
		}
		assert.True(t, foundCom, "应该包含 example.com")
	})

	t.Run("关键字包含已有TLD", func(t *testing.T) {
		// "example.com" → keyword="example", TLD "com" 已在 DefaultAllowedTLDs 白名单中
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example.com",
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// example.com 应该在匹配结果中
		found := false
		for _, domain := range resp.Matched {
			if domain.Domain == "example.com" {
				found = true
				break
			}
		}
		assert.True(t, found, "应该包含 example.com")

		// example.dev 不应该出现（dev 不在 DefaultAllowedTLDs 白名单中，关键字中也没有 .dev）
		for _, domain := range resp.Matched {
			assert.NotEqual(t, "example.dev", domain.Domain, "不应该包含 example.dev")
		}
	})

	t.Run("关键字包含多级TLD", func(t *testing.T) {
		// "example.dev.com" → keyword="example"，dev 不在白名单被忽略，com 在白名单中
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example.dev.com",
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// example.com 应该在匹配结果中（com 在白名单中）
		foundCom := false
		for _, domain := range resp.Matched {
			if domain.Domain == "example.com" {
				foundCom = true
			}
		}
		assert.True(t, foundCom, "应该包含 example.com")

		// example.dev 不应该出现（dev 不在白名单中，未被加入 TLD 列表）
		for _, domain := range resp.Matched {
			assert.NotEqual(t, "example.dev", domain.Domain, "不应该包含 example.dev")
		}
	})
}

func TestRecommendDomainsEmptyKeyword(t *testing.T) {
	mockClient := client.NewMockClient()

	// 设置空响应
	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", `<available></available><unavailable></unavailable>`))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("空关键词", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword: "",
		})

		// 应该能够处理空关键词（虽然可能没有结果）
		if err == nil {
			assert.NotNil(t, resp, "响应不应该为空")
			assert.Len(t, resp.Recommended, 0, "空关键词应该没有推荐域名")
		}
	})
}

func TestRecommendDomainsWithUnavailable(t *testing.T) {
	mockClient := client.NewMockClient()

	// 大部分域名都不可用的情况（所有域名均为 example + TLD 格式）
	availableContent := `<available>
	<domain price="1.99" premium="0">example.click</domain>
</available>
<unavailable>
	<domain>example.com</domain>
	<domain>example.net</domain>
	<domain>example.cc</domain>
	<domain>example.me</domain>
	<domain>example.xyz</domain>
</unavailable>`

	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", availableContent))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("大部分域名不可用", func(t *testing.T) {
		resp, err := service.RecommendDomains(ctx, &RecommendDomainsRequest{
			Keyword:        "example",
			IncludeMatched: true,
		})

		assert.NoError(t, err, "RecommendDomains 应该成功")
		assert.NotNil(t, resp, "响应不应该为空")

		// 验证不可用域名列表
		assert.GreaterOrEqual(t, len(resp.Unavailable), 1, "应该有不可用域名")

		// 验证可用域名数量较少
		totalAvailable := len(resp.Recommended) + len(resp.Matched)
		assert.LessOrEqual(t, totalAvailable, len(resp.Unavailable), "可用域名应该少于不可用域名")
	})
}
