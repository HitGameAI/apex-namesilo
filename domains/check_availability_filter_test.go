/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-01-17 00:30:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-01-17 01:16:26
 * @FilePath: \go-namesilo\domains\check_availability_filter_test.go
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

func TestCheckAvailabilityFilter(t *testing.T) {
	// 创建 Mock 客户端
	mockClient := client.NewMockClient()

	// 设置 Mock 响应内容
	availableContent := `<available>
	<domain price="8.99" premium="0">example.com</domain>
	<domain price="1.99" premium="0">test.net</domain>
	<domain price="15.99" premium="1">premium.com</domain>
	<domain price="2.99" premium="0">short.cc</domain>
	<domain price="3.99" premium="0">very.xyz</domain>
</available>
<unavailable>
	<domain>google.com</domain>
</unavailable>`

	// 设置 Mock 响应
	mockClient.WithMockResponse(client.MockSuccessXMLResponse("checkRegisterAvailability", availableContent))

	service := NewService(mockClient)
	ctx := context.Background()

	t.Run("无过滤条件", func(t *testing.T) {
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com", "test.net", "google.com"},
			Filter:  nil,
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")
		assert.Len(t, resp.Available, 5, "应该有 5 个可用域名")
		assert.Len(t, resp.Unavailable, 1, "应该有 1 个不可用域名")
		assert.Len(t, resp.Filtered, 0, "应该有 0 个被过滤域名")
	})

	t.Run("价格过滤-MaxPrice", func(t *testing.T) {
		maxPrice := 3.0
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com", "test.net"},
			Filter: &DomainFilter{
				MaxPrice: &maxPrice,
			},
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")
		assert.Len(t, resp.Available, 2, "应该有 2 个可用域名（价格 <= 3.0）")
		assert.Len(t, resp.Filtered, 3, "应该有 3 个被过滤域名")

		// 验证价格
		for _, domain := range resp.Available {
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
		}
	})

	t.Run("价格过滤-MinPrice和MaxPrice", func(t *testing.T) {
		minPrice := 2.0
		maxPrice := 5.0
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com"},
			Filter: &DomainFilter{
				MinPrice: &minPrice,
				MaxPrice: &maxPrice,
			},
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")
		assert.Len(t, resp.Available, 2, "应该有 2 个可用域名（2.0 <= 价格 <= 5.0）")

		for _, domain := range resp.Available {
			assert.GreaterOrEqual(t, domain.Price, minPrice, "域名 %s 价格应该 >= %.2f", domain.Domain, minPrice)
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)
		}
	})

	t.Run("TLD过滤-IncludeTLDs", func(t *testing.T) {
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com", "test.net"},
			Filter: &DomainFilter{
				IncludeTLDs: []string{"com", "net"},
			},
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")
		assert.Len(t, resp.Available, 3, "应该有 3 个可用域名（TLD 为 com/net）")

		for _, domain := range resp.Available {
			tld := extractTLD(domain.Domain)
			assert.Contains(t, []string{"com", "net"}, tld, "域名 %s 的 TLD 应该是 com 或 net", domain.Domain)
		}
	})

	t.Run("TLD过滤-ExcludeTLDs", func(t *testing.T) {
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com"},
			Filter: &DomainFilter{
				ExcludeTLDs: []string{"xyz", "cc"},
			},
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")
		assert.GreaterOrEqual(t, len(resp.Filtered), 2, "至少应该有 2 个被过滤域名（xyz/cc）")

		// 验证可用域名中不包含被排除的 TLD
		for _, domain := range resp.Available {
			tld := extractTLD(domain.Domain)
			assert.NotContains(t, []string{"xyz", "cc"}, tld, "域名 %s 不应该包含被排除的 TLD", domain.Domain)
		}
	})

	t.Run("长度过滤", func(t *testing.T) {
		minLen := 4
		maxLen := 10
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com"},
			Filter: &DomainFilter{
				MinLength: &minLen,
				MaxLength: &maxLen,
			},
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")

		for _, domain := range resp.Available {
			name := extractDomainName(domain.Domain)
			nameLen := len(name)
			assert.GreaterOrEqual(t, nameLen, minLen, "域名 %s 长度应该 >= %d", domain.Domain, minLen)
			assert.LessOrEqual(t, nameLen, maxLen, "域名 %s 长度应该 <= %d", domain.Domain, maxLen)
		}
	})

	t.Run("Premium过滤", func(t *testing.T) {
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com"},
			Filter: &DomainFilter{
				ExcludePremium: true,
			},
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")

		// 验证可用域名中不包含 Premium 域名
		for _, domain := range resp.Available {
			assert.NotEqual(t, "premium.com", domain.Domain, "不应该包含 Premium 域名")
		}

		// 检查 Filtered 中是否包含 premium.com
		foundPremium := false
		for _, domain := range resp.Filtered {
			if domain.Domain == "premium.com" {
				foundPremium = true
				break
			}
		}
		assert.True(t, foundPremium, "premium.com 应该在过滤列表中")
	})

	t.Run("组合过滤-价格+TLD+长度", func(t *testing.T) {
		maxPrice := 5.0
		maxLen := 8
		resp, err := service.CheckAvailabilityFilter(ctx, &CheckAvailabilityFilterRequest{
			Domains: []string{"example.com"},
			Filter: &DomainFilter{
				MaxPrice:    &maxPrice,
				IncludeTLDs: []string{"net", "cc"},
				MaxLength:   &maxLen,
			},
		})

		assert.NoError(t, err, "CheckAvailabilityFilter 应该成功")

		// 验证所有条件
		for _, domain := range resp.Available {
			// 验证价格
			assert.LessOrEqual(t, domain.Price, maxPrice, "域名 %s 价格应该 <= %.2f", domain.Domain, maxPrice)

			// 验证 TLD
			tld := extractTLD(domain.Domain)
			assert.Contains(t, []string{"net", "cc"}, tld, "域名 %s 的 TLD 应该是 net 或 cc", domain.Domain)

			// 验证长度
			name := extractDomainName(domain.Domain)
			assert.LessOrEqual(t, len(name), maxLen, "域名 %s 长度应该 <= %d", domain.Domain, maxLen)
		}
	})
}

func TestExtractTLD(t *testing.T) {
	tests := []struct {
		domain   string
		expected string
	}{
		{"example.com", "com"},
		{"test.net", "net"},
		{"sub.example.com", "com"},
		{"example", "example"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			result := extractTLD(tt.domain)
			assert.Equal(t, tt.expected, result, "extractTLD(%q) 应该返回 %q", tt.domain, tt.expected)
		})
	}
}

func TestExtractDomainName(t *testing.T) {
	tests := []struct {
		domain   string
		expected string
	}{
		{"example.com", "example"},
		{"test.net", "test"},
		{"sub.example.com", "sub.example"},
		{"example", "example"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			result := extractDomainName(tt.domain)
			assert.Equal(t, tt.expected, result, "extractDomainName(%q) 应该返回 %q", tt.domain, tt.expected)
		})
	}
}
