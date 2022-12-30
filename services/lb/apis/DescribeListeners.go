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
    lb "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeListenersRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* listenerNames - 监听器名称列表，支持多个
listenerIds - 监听器Id列表，支持多个
loadBalancerId - 负载均衡器Id，支持单个
loadBalancerType - 负载均衡类型，取值为：alb、nlb、dnlb，默认alb，支持单个
urlMapIds - 【仅alb支持】转发规则组Id列表，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeListenersRequest(
    regionId string,
) *DescribeListenersRequest {

	return &DescribeListenersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: listenerNames - 监听器名称列表，支持多个
listenerIds - 监听器Id列表，支持多个
loadBalancerId - 负载均衡器Id，支持单个
loadBalancerType - 负载均衡类型，取值为：alb、nlb、dnlb，默认alb，支持单个
urlMapIds - 【仅alb支持】转发规则组Id列表，支持多个
 (Optional)
 */
func NewDescribeListenersRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeListenersRequest {

    return &DescribeListenersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/",
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
func NewDescribeListenersRequestWithoutParam() *DescribeListenersRequest {

    return &DescribeListenersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeListenersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeListenersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeListenersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: listenerNames - 监听器名称列表，支持多个
listenerIds - 监听器Id列表，支持多个
loadBalancerId - 负载均衡器Id，支持单个
loadBalancerType - 负载均衡类型，取值为：alb、nlb、dnlb，默认alb，支持单个
urlMapIds - 【仅alb支持】转发规则组Id列表，支持多个
(Optional) */
func (r *DescribeListenersRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeListenersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeListenersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeListenersResult `json:"result"`
}

type DescribeListenersResult struct {
    Listeners []lb.Listener `json:"listeners"`
    TotalCount int `json:"totalCount"`
}