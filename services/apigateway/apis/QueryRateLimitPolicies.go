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
    apigateway "github.com/jdcloud-api/jdcloud-sdk-go/services/apigateway/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type QueryRateLimitPoliciesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* policyName - 策略名称，模糊匹配
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryRateLimitPoliciesRequest(
    regionId string,
) *QueryRateLimitPoliciesRequest {

	return &QueryRateLimitPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rateLimitPolicies",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: policyName - 策略名称，模糊匹配
 (Optional)
 */
func NewQueryRateLimitPoliciesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *QueryRateLimitPoliciesRequest {

    return &QueryRateLimitPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rateLimitPolicies",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryRateLimitPoliciesRequestWithoutParam() *QueryRateLimitPoliciesRequest {

    return &QueryRateLimitPoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rateLimitPolicies",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QueryRateLimitPoliciesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *QueryRateLimitPoliciesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *QueryRateLimitPoliciesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: policyName - 策略名称，模糊匹配
(Optional) */
func (r *QueryRateLimitPoliciesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryRateLimitPoliciesRequest) GetRegionId() string {
    return r.RegionId
}

type QueryRateLimitPoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryRateLimitPoliciesResult `json:"result"`
}

type QueryRateLimitPoliciesResult struct {
    RateLimitPolicys []apigateway.RateLimitPolicy `json:"rateLimitPolicys"`
    TotalCount int `json:"totalCount"`
}