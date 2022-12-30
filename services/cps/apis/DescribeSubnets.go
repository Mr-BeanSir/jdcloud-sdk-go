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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeSubnetsRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 可用区，精确匹配 (Optional) */
    Az *string `json:"az"`

    /* 子网名称 (Optional) */
    Name *string `json:"name"`

    /* 私有网络ID，精确匹配 (Optional) */
    VpcId *string `json:"vpcId"`

    /* subnetId - 子网ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSubnetsRequest(
    regionId string,
) *DescribeSubnetsRequest {

	return &DescribeSubnetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param az: 可用区，精确匹配 (Optional)
 * param name: 子网名称 (Optional)
 * param vpcId: 私有网络ID，精确匹配 (Optional)
 * param filters: subnetId - 子网ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeSubnetsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    az *string,
    name *string,
    vpcId *string,
    filters []common.Filter,
) *DescribeSubnetsRequest {

    return &DescribeSubnetsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Az: az,
        Name: name,
        VpcId: vpcId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSubnetsRequestWithoutParam() *DescribeSubnetsRequest {

    return &DescribeSubnetsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeSubnetsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeSubnetsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeSubnetsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param az: 可用区，精确匹配(Optional) */
func (r *DescribeSubnetsRequest) SetAz(az string) {
    r.Az = &az
}

/* param name: 子网名称(Optional) */
func (r *DescribeSubnetsRequest) SetName(name string) {
    r.Name = &name
}

/* param vpcId: 私有网络ID，精确匹配(Optional) */
func (r *DescribeSubnetsRequest) SetVpcId(vpcId string) {
    r.VpcId = &vpcId
}

/* param filters: subnetId - 子网ID，精确匹配，支持多个
(Optional) */
func (r *DescribeSubnetsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSubnetsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSubnetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSubnetsResult `json:"result"`
}

type DescribeSubnetsResult struct {
    Subnets []cps.Subnet `json:"subnets"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}