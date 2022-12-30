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

type OperateRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 物联网卡操作请求类型  */
    RequestType string `json:"requestType"`

    /* 物联网卡操作请求参数json串  */
    RequestParam string `json:"requestParam"`
}

/*
 * param regionId: Region ID (Required)
 * param requestType: 物联网卡操作请求类型 (Required)
 * param requestParam: 物联网卡操作请求参数json串 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOperateRequest(
    regionId string,
    requestType string,
    requestParam string,
) *OperateRequest {

	return &OperateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/operate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RequestType: requestType,
        RequestParam: requestParam,
	}
}

/*
 * param regionId: Region ID (Required)
 * param requestType: 物联网卡操作请求类型 (Required)
 * param requestParam: 物联网卡操作请求参数json串 (Required)
 */
func NewOperateRequestWithAllParams(
    regionId string,
    requestType string,
    requestParam string,
) *OperateRequest {

    return &OperateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/operate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RequestType: requestType,
        RequestParam: requestParam,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOperateRequestWithoutParam() *OperateRequest {

    return &OperateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/operate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *OperateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param requestType: 物联网卡操作请求类型(Required) */
func (r *OperateRequest) SetRequestType(requestType string) {
    r.RequestType = requestType
}

/* param requestParam: 物联网卡操作请求参数json串(Required) */
func (r *OperateRequest) SetRequestParam(requestParam string) {
    r.RequestParam = requestParam
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OperateRequest) GetRegionId() string {
    return r.RegionId
}

type OperateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OperateResult `json:"result"`
}

type OperateResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Result string `json:"result"`
}