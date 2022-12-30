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

type DescribeExposeTypeRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* vpcId  */
    VpcId string `json:"vpcId"`
}

/*
 * param regionId: 地域代码 (Required)
 * param vpcId: vpcId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeExposeTypeRequest(
    regionId string,
    vpcId string,
) *DescribeExposeTypeRequest {

	return &DescribeExposeTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcs/{vpcId}:describeExposeType",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcId: vpcId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param vpcId: vpcId (Required)
 */
func NewDescribeExposeTypeRequestWithAllParams(
    regionId string,
    vpcId string,
) *DescribeExposeTypeRequest {

    return &DescribeExposeTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}:describeExposeType",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcId: vpcId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeExposeTypeRequestWithoutParam() *DescribeExposeTypeRequest {

    return &DescribeExposeTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}:describeExposeType",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeExposeTypeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpcId: vpcId(Required) */
func (r *DescribeExposeTypeRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeExposeTypeRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeExposeTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeExposeTypeResult `json:"result"`
}

type DescribeExposeTypeResult struct {
    ExposeType []string `json:"exposeType"`
}