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

type ReleaseAuthorizationTokenRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 注册表名称  */
    RegistryName string `json:"registryName"`

    /* 准备释放的 token ID，功能为指定token释放。 (Optional) */
    AuthorizationToken *string `json:"authorizationToken"`

    /* true 表示强制删除用户当前registry下所有有效token的标志；false 表示删除所有有效token。 (Optional) */
    ForceAll *bool `json:"forceAll"`
}

/*
 * param regionId: 地域 ID (Required)
 * param registryName: 注册表名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewReleaseAuthorizationTokenRequest(
    regionId string,
    registryName string,
) *ReleaseAuthorizationTokenRequest {

	return &ReleaseAuthorizationTokenRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries/{registryName}:releaseAuthorizationToken",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RegistryName: registryName,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param registryName: 注册表名称 (Required)
 * param authorizationToken: 准备释放的 token ID，功能为指定token释放。 (Optional)
 * param forceAll: true 表示强制删除用户当前registry下所有有效token的标志；false 表示删除所有有效token。 (Optional)
 */
func NewReleaseAuthorizationTokenRequestWithAllParams(
    regionId string,
    registryName string,
    authorizationToken *string,
    forceAll *bool,
) *ReleaseAuthorizationTokenRequest {

    return &ReleaseAuthorizationTokenRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}:releaseAuthorizationToken",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryName: registryName,
        AuthorizationToken: authorizationToken,
        ForceAll: forceAll,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewReleaseAuthorizationTokenRequestWithoutParam() *ReleaseAuthorizationTokenRequest {

    return &ReleaseAuthorizationTokenRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}:releaseAuthorizationToken",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *ReleaseAuthorizationTokenRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryName: 注册表名称(Required) */
func (r *ReleaseAuthorizationTokenRequest) SetRegistryName(registryName string) {
    r.RegistryName = registryName
}

/* param authorizationToken: 准备释放的 token ID，功能为指定token释放。(Optional) */
func (r *ReleaseAuthorizationTokenRequest) SetAuthorizationToken(authorizationToken string) {
    r.AuthorizationToken = &authorizationToken
}

/* param forceAll: true 表示强制删除用户当前registry下所有有效token的标志；false 表示删除所有有效token。(Optional) */
func (r *ReleaseAuthorizationTokenRequest) SetForceAll(forceAll bool) {
    r.ForceAll = &forceAll
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ReleaseAuthorizationTokenRequest) GetRegionId() string {
    return r.RegionId
}

type ReleaseAuthorizationTokenResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ReleaseAuthorizationTokenResult `json:"result"`
}

type ReleaseAuthorizationTokenResult struct {
}