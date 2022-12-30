// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type SetHttpSslRequest struct {

    core.JDCloudRequest

    /* 域名ID  */
    DomainId int `json:"domainId"`

    /* 证书来源。取值范围：default (Optional) */
    Source *string `json:"source"`

    /* 证书标题 (Optional) */
    Title *string `json:"title"`

    /* 证书内容 (Optional) */
    SslCert *string `json:"sslCert"`

    /* 证书私钥 (Optional) */
    SslKey *string `json:"sslKey"`

    /* 跳转类型。取值范围：
default - 采用回源域名的默认协议
http - 强制采用http协议回源
https - 强制采用https协议回源
 (Optional) */
    JumpType *string `json:"jumpType"`

    /* SSL配置启用状态 (Optional) */
    Enabled *bool `json:"enabled"`
}

/*
 * param domainId: 域名ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetHttpSslRequest(
    domainId int,
) *SetHttpSslRequest {

	return &SetHttpSslRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainId}:setHttpSsl",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        DomainId: domainId,
	}
}

/*
 * param domainId: 域名ID (Required)
 * param source: 证书来源。取值范围：default (Optional)
 * param title: 证书标题 (Optional)
 * param sslCert: 证书内容 (Optional)
 * param sslKey: 证书私钥 (Optional)
 * param jumpType: 跳转类型。取值范围：
default - 采用回源域名的默认协议
http - 强制采用http协议回源
https - 强制采用https协议回源
 (Optional)
 * param enabled: SSL配置启用状态 (Optional)
 */
func NewSetHttpSslRequestWithAllParams(
    domainId int,
    source *string,
    title *string,
    sslCert *string,
    sslKey *string,
    jumpType *string,
    enabled *bool,
) *SetHttpSslRequest {

    return &SetHttpSslRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setHttpSsl",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        DomainId: domainId,
        Source: source,
        Title: title,
        SslCert: sslCert,
        SslKey: sslKey,
        JumpType: jumpType,
        Enabled: enabled,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetHttpSslRequestWithoutParam() *SetHttpSslRequest {

    return &SetHttpSslRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setHttpSsl",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainId: 域名ID(Required) */
func (r *SetHttpSslRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

/* param source: 证书来源。取值范围：default(Optional) */
func (r *SetHttpSslRequest) SetSource(source string) {
    r.Source = &source
}

/* param title: 证书标题(Optional) */
func (r *SetHttpSslRequest) SetTitle(title string) {
    r.Title = &title
}

/* param sslCert: 证书内容(Optional) */
func (r *SetHttpSslRequest) SetSslCert(sslCert string) {
    r.SslCert = &sslCert
}

/* param sslKey: 证书私钥(Optional) */
func (r *SetHttpSslRequest) SetSslKey(sslKey string) {
    r.SslKey = &sslKey
}

/* param jumpType: 跳转类型。取值范围：
default - 采用回源域名的默认协议
http - 强制采用http协议回源
https - 强制采用https协议回源
(Optional) */
func (r *SetHttpSslRequest) SetJumpType(jumpType string) {
    r.JumpType = &jumpType
}

/* param enabled: SSL配置启用状态(Optional) */
func (r *SetHttpSslRequest) SetEnabled(enabled bool) {
    r.Enabled = &enabled
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetHttpSslRequest) GetRegionId() string {
    return ""
}

type SetHttpSslResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetHttpSslResult `json:"result"`
}

type SetHttpSslResult struct {
}