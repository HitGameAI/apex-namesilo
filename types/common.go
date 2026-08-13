/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-12-30 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-08-12 15:13:59
 * @FilePath: \apex-core-service\go-namesilo\types\common.go
 * @Description:
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ResponseType string

const (
	ResponseTypeJSON ResponseType = "json"
	ResponseTypeXML  ResponseType = "xml"
)

func (rt ResponseType) String() string {
	return string(rt)
}

// Config NameSilo API 配置
type Config struct {
	ID           uint64        // 配置 ID
	APIKey       string        // API 密钥
	APIVersion   string        // API 版本，默认为 "1"
	BaseURL      string        // API 基础 URL，默认为 https://www.namesilo.com/api
	PublicURL    string        // 公共访问 URL，默认为 https://www.namesilo.com
	ResponseType ResponseType  // 响应类型，json 或 xml，默认 xml
	Timeout      time.Duration // 请求超时时间，默认 30 秒
	Debug        bool          // 是否开启调试模式
	Logger       Logger        // 日志记录器
}

// BaseResponse API 基础响应结构
type BaseResponse struct {
	XMLName xml.Name       `xml:"namesilo" json:"-"`
	Request RequestSection `xml:"request" json:"request"`
	Reply   CommonReply    `xml:"reply" json:"reply"`
}

// RequestSection 请求信息部分
type RequestSection struct {
	Operation string `xml:"operation" json:"operation"` // API 操作名称
	IP        string `xml:"ip" json:"ip"`               // 请求 IP
}

// CommonReply 通用响应 Reply 部分
type CommonReply struct {
	Code        int                    `xml:"code" json:"code"`                           // 状态码
	Detail      string                 `xml:"detail" json:"detail"`                       // 详细信息
	Message     string                 `xml:"message" json:"message,omitempty"`           // 消息（可选）
	Domain      string                 `xml:"domain" json:"domain,omitempty"`             // 域名（部分接口返回）
	OrderAmount float64                `xml:"order_amount" json:"order_amount,omitempty"` // 订单金额（部分接口返回）
	RecordID    string                 `xml:"record_id" json:"record_id,omitempty"`       // 记录ID（DNS相关接口返回）
	Extra       map[string]interface{} `json:"-" xml:"-"`                                 // 额外字段
}

// UnmarshalJSON 自定义 JSON 解析，处理 code 可能是字符串或数字的情况
func (c *CommonReply) UnmarshalJSON(data []byte) error {
	type alias struct {
		Code        json.RawMessage `json:"code"`
		Detail      string          `json:"detail"`
		Message     string          `json:"message"`
		Domain      string          `json:"domain"`
		OrderAmount json.RawMessage `json:"order_amount"`
		RecordID    string          `json:"record_id"`
	}

	var aux alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// 解析 code 字段（可能是字符串或数字）
	if len(aux.Code) > 0 {
		var codeInt int
		if err := json.Unmarshal(aux.Code, &codeInt); err == nil {
			c.Code = codeInt
		} else {
			var codeStr string
			if err := json.Unmarshal(aux.Code, &codeStr); err != nil {
				return fmt.Errorf("failed to parse code field: %w", err)
			}
			codeStr = strings.TrimSpace(codeStr)
			if codeStr == "" {
				return fmt.Errorf("code field is empty")
			}
			code, err := strconv.Atoi(codeStr)
			if err != nil {
				return fmt.Errorf("failed to parse code field: %w", err)
			}
			c.Code = code
		}
	}

	c.Detail = aux.Detail
	c.Message = aux.Message
	c.Domain = aux.Domain
	c.RecordID = aux.RecordID

	// 解析 order_amount 字段（可能是字符串或数字）
	if len(aux.OrderAmount) > 0 {
		var amountFloat float64
		if err := json.Unmarshal(aux.OrderAmount, &amountFloat); err == nil {
			c.OrderAmount = amountFloat
		} else {
			var amountStr string
			if err := json.Unmarshal(aux.OrderAmount, &amountStr); err == nil {
				if amountStr != "" {
					if amount, err := strconv.ParseFloat(amountStr, 64); err == nil {
						c.OrderAmount = amount
					}
				}
			}
		}
	}

	// 解析额外字段
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	knownKeys := map[string]struct{}{
		"code": {}, "detail": {}, "message": {}, "domain": {}, "order_amount": {}, "record_id": {},
	}

	c.Extra = make(map[string]interface{})
	for key, value := range raw {
		if _, ok := knownKeys[key]; ok {
			continue
		}
		var v interface{}
		if err := json.Unmarshal(value, &v); err != nil {
			c.Extra[key] = string(value)
		} else {
			c.Extra[key] = v
		}
	}

	return nil
}

// Success 判断响应是否成功
func (c *CommonReply) Success() bool {
	return c.Code == StatusSuccess ||
		c.Code == StatusSuccessWithPartialFailure ||
		c.Code == StatusSuccessWithContactIssue
}

// Error 实现 error 接口
func (c *CommonReply) Error() string {
	if c.Success() {
		return ""
	}
	if c.Detail != "" {
		return fmt.Sprintf("NameSilo API error (code %d): %s", c.Code, c.Detail)
	}
	return fmt.Sprintf("NameSilo API error (code %d): %s", c.Code, c.Message)
}

// NameSilo API 状态码常量
const (
	// 成功状态码
	StatusSuccess                   = 300 // API 操作成功
	StatusSuccessWithPartialFailure = 301 // 注册成功，但部分主机无效
	StatusSuccessWithContactIssue   = 302 // 订单成功，但联系信息有误

	// 一般错误 (100-199)
	StatusNoHTTPS               = 101 // 未使用 HTTPS
	StatusNoVersion             = 102 // 未指定版本
	StatusInvalidAPIVersion     = 103 // 无效的 API 版本
	StatusNoType                = 104 // 未指定类型
	StatusInvalidAPIType        = 105 // 无效的 API 类型
	StatusNoOperation           = 106 // 未指定操作
	StatusInvalidAPIOperation   = 107 // 无效的 API 操作
	StatusMissingParameter      = 108 // 缺少指定操作的参数
	StatusNoAPIKey              = 109 // 未指定 API 密钥
	StatusInvalidAPIKey         = 110 // 无效的 API 密钥
	StatusInvalidUser           = 111 // 无效用户
	StatusAPINotForSubAccount   = 112 // API 不适用于子账户
	StatusIPNotAllowed          = 113 // 您的 IP 无法访问此 API 帐户
	StatusInvalidDomainSyntax   = 114 // 无效的域语法
	StatusRegistryNotRespond    = 115 // 中央注册处未响应
	StatusInvalidSandboxAccount = 116 // 无效的沙盒帐户
	StatusNoCreditCard          = 117 // 信用卡资料不存在或未关联
	StatusCreditCardNotVerified = 118 // 信用卡资料尚未验证
	StatusInsufficientFunds     = 119 // 账户资金不足
	StatusAPIKeyMustBeGET       = 120 // API 密钥必须以 GET 形式传递

	// 域名相关错误 (200-299)
	StatusDomainNotActive         = 200 // 域未激活或不属于用户
	StatusInternalError           = 201 // 内部系统错误
	StatusGeneralError            = 210 // 一般错误
	StatusAutoRenewEnabled        = 250 // 域已设置为自动续订
	StatusAutoRenewDisabled       = 251 // 域已设置为不自动续订
	StatusDomainLocked            = 252 // 域已被锁定
	StatusDomainUnlocked          = 253 // 域名已解锁
	StatusCannotUpdateNS          = 254 // 无法更新名称服务器
	StatusDomainAlreadyPrivate    = 255 // 域名已经是私有的
	StatusDomainNotPrivate        = 256 // 域名已经不是私有的
	StatusDomainProcessingError   = 261 // 域处理错误
	StatusDomainAlreadyActive     = 262 // 域名已处于活动状态
	StatusInvalidYears            = 263 // 年数无效或未提供
	StatusCannotRenew             = 264 // 域名无法续订指定年限
	StatusCannotTransfer          = 265 // 域名目前无法转移
	StatusNoTransferExists        = 266 // 域中不存在此用户的域名转移
	StatusInvalidOrUnsupportedTLD = 267 // 域名无效或不支持 TLD
	StatusDNSModifyError          = 280 // DNS 修改错误

	// 其他错误
	StatusRequestInProgress = 400 // 现有 API 请求仍在处理中
)

// StatusCodeMessage 返回状态码对应的消息
func StatusCodeMessage(code int) string {
	messages := map[int]string{
		StatusSuccess:                   "API 操作成功",
		StatusSuccessWithPartialFailure: "注册成功，但部分主机无效",
		StatusSuccessWithContactIssue:   "订单成功，但联系信息有误",
		StatusNoHTTPS:                   "未使用 HTTPS",
		StatusNoVersion:                 "未指定版本",
		StatusInvalidAPIVersion:         "无效的 API 版本",
		StatusNoType:                    "未指定类型",
		StatusInvalidAPIType:            "无效的 API 类型",
		StatusNoOperation:               "未指定操作",
		StatusInvalidAPIOperation:       "无效的 API 操作",
		StatusMissingParameter:          "缺少指定操作的参数",
		StatusNoAPIKey:                  "未指定 API 密钥",
		StatusInvalidAPIKey:             "无效的 API 密钥",
		StatusInvalidUser:               "无效用户",
		StatusAPINotForSubAccount:       "API 不适用于子账户",
		StatusIPNotAllowed:              "您的 IP 无法访问此 API 帐户",
		StatusInvalidDomainSyntax:       "无效的域语法",
		StatusRegistryNotRespond:        "中央注册处未响应",
		StatusInvalidSandboxAccount:     "无效的沙盒帐户",
		StatusNoCreditCard:              "信用卡资料不存在或未关联",
		StatusCreditCardNotVerified:     "信用卡资料尚未验证",
		StatusInsufficientFunds:         "账户资金不足",
		StatusAPIKeyMustBeGET:           "API 密钥必须以 GET 形式传递",
		StatusDomainNotActive:           "域未激活或不属于用户",
		StatusInternalError:             "内部系统错误",
		StatusGeneralError:              "一般错误",
		StatusAutoRenewEnabled:          "域已设置为自动续订",
		StatusAutoRenewDisabled:         "域已设置为不自动续订",
		StatusDomainLocked:              "域已被锁定",
		StatusDomainUnlocked:            "域名已解锁",
		StatusCannotUpdateNS:            "无法更新名称服务器",
		StatusDomainAlreadyPrivate:      "域名已经是私有的",
		StatusDomainNotPrivate:          "域名已经不是私有的",
		StatusDomainProcessingError:     "域处理错误",
		StatusDomainAlreadyActive:       "域名已处于活动状态",
		StatusInvalidYears:              "年数无效或未提供",
		StatusCannotRenew:               "域名无法续订指定年限",
		StatusCannotTransfer:            "域名目前无法转移",
		StatusNoTransferExists:          "域中不存在此用户的域名转移",
		StatusInvalidOrUnsupportedTLD:   "域名无效或不支持 TLD",
		StatusDNSModifyError:            "DNS 修改错误",
		StatusRequestInProgress:         "现有 API 请求仍在处理中",
	}

	if msg, ok := messages[code]; ok {
		return msg
	}
	return fmt.Sprintf("未知状态码: %d", code)
}

// API 常量
const (
	DefaultAPIVersion = "1"
	DefaultType       = ResponseTypeXML // NameSilo API 默认返回 XML
	DefaultBaseURL    = "https://www.namesilo.com/api"
	DefaultPublicURL  = "https://www.namesilo.com/public/api"
	DefaultTimeout    = 10 * time.Second
)

// 分页默认值
const (
	DefaultPage     = 1
	DefaultPageSize = 10
)

var (
	// DefaultAllowedTLDs 默认允许的 TLD 白名单
	// 域名搜索时若未指定 IncludeTLDs，则使用此白名单过滤
	// 关键字中包含的 TLD 必须在此白名单内才会被搜索
	DefaultAllowedTLDs = []string{
		"com", "cc", "me", "xyz", "top", "click", "buzz", "club", "win", "life", "net", "ai", "org",
		"cfd", "wiki", "io", "sbs", "one", "help", "art", "design",
		"cyou", "online", "shop", "de", "info", "website",
		"co", "lat", "space", "lol", "icu", "biz", "site", "cam",
		"cloud", "ac", "fun", "autos",
	}
)
