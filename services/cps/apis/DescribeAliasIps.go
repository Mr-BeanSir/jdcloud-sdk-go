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

type DescribeAliasIpsRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 子网ID (Optional) */
    SubnetId *string `json:"subnetId"`

    /* 实例ID (Optional) */
    InstanceId *string `json:"instanceId"`

    /* CIDR段，模糊搜索 (Optional) */
    Cidr *string `json:"cidr"`

    /* aliasIpId - 别名IP id<br/>
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAliasIpsRequest(
    regionId string,
) *DescribeAliasIpsRequest {

	return &DescribeAliasIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/aliasIps",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param subnetId: 子网ID (Optional)
 * param instanceId: 实例ID (Optional)
 * param cidr: CIDR段，模糊搜索 (Optional)
 * param filters: aliasIpId - 别名IP id<br/>
 (Optional)
 */
func NewDescribeAliasIpsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    subnetId *string,
    instanceId *string,
    cidr *string,
    filters []common.Filter,
) *DescribeAliasIpsRequest {

    return &DescribeAliasIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/aliasIps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        SubnetId: subnetId,
        InstanceId: instanceId,
        Cidr: cidr,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAliasIpsRequestWithoutParam() *DescribeAliasIpsRequest {

    return &DescribeAliasIpsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/aliasIps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegions）获取云物理服务器支持的地域(Required) */
func (r *DescribeAliasIpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeAliasIpsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeAliasIpsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param subnetId: 子网ID(Optional) */
func (r *DescribeAliasIpsRequest) SetSubnetId(subnetId string) {
    r.SubnetId = &subnetId
}

/* param instanceId: 实例ID(Optional) */
func (r *DescribeAliasIpsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}

/* param cidr: CIDR段，模糊搜索(Optional) */
func (r *DescribeAliasIpsRequest) SetCidr(cidr string) {
    r.Cidr = &cidr
}

/* param filters: aliasIpId - 别名IP id<br/>
(Optional) */
func (r *DescribeAliasIpsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAliasIpsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAliasIpsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAliasIpsResult `json:"result"`
}

type DescribeAliasIpsResult struct {
    AliasIps []cps.AliasIp `json:"aliasIps"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}