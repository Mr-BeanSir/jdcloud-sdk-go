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

type DescribeSSLRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSSLRequest(
    regionId string,
    instanceId string,
) *DescribeSSLRequest {

	return &DescribeSSLRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/ssl",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 */
func NewDescribeSSLRequestWithAllParams(
    regionId string,
    instanceId string,
) *DescribeSSLRequest {

    return &DescribeSSLRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ssl",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSSLRequestWithoutParam() *DescribeSSLRequest {

    return &DescribeSSLRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ssl",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeSSLRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 实例ID(Required) */
func (r *DescribeSSLRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSSLRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSSLResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSSLResult `json:"result"`
}

type DescribeSSLResult struct {
    Available bool `json:"available"`
}