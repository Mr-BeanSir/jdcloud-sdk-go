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
    cr "github.com/jdcloud-api/jdcloud-sdk-go/services/cr/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeAuthorizationTokensRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 注册表名称  */
    RegistryName string `json:"registryName"`

    /* token - 令牌 ID，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域 ID (Required)
 * param registryName: 注册表名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAuthorizationTokensRequest(
    regionId string,
    registryName string,
) *DescribeAuthorizationTokensRequest {

	return &DescribeAuthorizationTokensRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries/{registryName}/tokens",
			Method:  "GET",
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
 * param filters: token - 令牌 ID，支持多个
 (Optional)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 */
func NewDescribeAuthorizationTokensRequestWithAllParams(
    regionId string,
    registryName string,
    filters []common.Filter,
    pageNumber *int,
    pageSize *int,
) *DescribeAuthorizationTokensRequest {

    return &DescribeAuthorizationTokensRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}/tokens",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryName: registryName,
        Filters: filters,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAuthorizationTokensRequestWithoutParam() *DescribeAuthorizationTokensRequest {

    return &DescribeAuthorizationTokensRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}/tokens",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *DescribeAuthorizationTokensRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryName: 注册表名称(Required) */
func (r *DescribeAuthorizationTokensRequest) SetRegistryName(registryName string) {
    r.RegistryName = registryName
}

/* param filters: token - 令牌 ID，支持多个
(Optional) */
func (r *DescribeAuthorizationTokensRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeAuthorizationTokensRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribeAuthorizationTokensRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAuthorizationTokensRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAuthorizationTokensResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAuthorizationTokensResult `json:"result"`
}

type DescribeAuthorizationTokensResult struct {
    AuthorizationTokens []cr.AuthorizationData `json:"authorizationTokens"`
    TotalCount int `json:"totalCount"`
}