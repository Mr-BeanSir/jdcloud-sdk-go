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
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeNetworkInterfacesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* networkInterfaceIds - 弹性网卡ID列表，支持多个
networkInterfaceNames - 弹性网卡名称列表，支持多个
vpcId - 弹性网卡所属vpc Id，支持单个
subnetId	- 弹性网卡所属子网Id，支持单个
role - 网卡角色，取值范围：Primary（主网卡）、Secondary（辅助网卡）、Managed （受管网卡），支持单个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeNetworkInterfacesRequest(
    regionId string,
) *DescribeNetworkInterfacesRequest {

	return &DescribeNetworkInterfacesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydNetworkInterfaces",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: networkInterfaceIds - 弹性网卡ID列表，支持多个
networkInterfaceNames - 弹性网卡名称列表，支持多个
vpcId - 弹性网卡所属vpc Id，支持单个
subnetId	- 弹性网卡所属子网Id，支持单个
role - 网卡角色，取值范围：Primary（主网卡）、Secondary（辅助网卡）、Managed （受管网卡），支持单个
 (Optional)
 */
func NewDescribeNetworkInterfacesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeNetworkInterfacesRequest {

    return &DescribeNetworkInterfacesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydNetworkInterfaces",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeNetworkInterfacesRequestWithoutParam() *DescribeNetworkInterfacesRequest {

    return &DescribeNetworkInterfacesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydNetworkInterfaces",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeNetworkInterfacesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeNetworkInterfacesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeNetworkInterfacesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: networkInterfaceIds - 弹性网卡ID列表，支持多个
networkInterfaceNames - 弹性网卡名称列表，支持多个
vpcId - 弹性网卡所属vpc Id，支持单个
subnetId	- 弹性网卡所属子网Id，支持单个
role - 网卡角色，取值范围：Primary（主网卡）、Secondary（辅助网卡）、Managed （受管网卡），支持单个
(Optional) */
func (r *DescribeNetworkInterfacesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeNetworkInterfacesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeNetworkInterfacesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeNetworkInterfacesResult `json:"result"`
}

type DescribeNetworkInterfacesResult struct {
    NetworkInterfaces []vpc.NetworkInterface `json:"networkInterfaces"`
    TotalCount int `json:"totalCount"`
}