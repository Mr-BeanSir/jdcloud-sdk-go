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
    waf "github.com/jdcloud-api/jdcloud-sdk-go/services/waf/models"
)

type GetQpsDataRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 请求  */
    Req *waf.GetChartReq `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param req: 请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetQpsDataRequest(
    regionId string,
    req *waf.GetChartReq,
) *GetQpsDataRequest {

	return &GetQpsDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/chart:getQpsData",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param req: 请求 (Required)
 */
func NewGetQpsDataRequestWithAllParams(
    regionId string,
    req *waf.GetChartReq,
) *GetQpsDataRequest {

    return &GetQpsDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/chart:getQpsData",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetQpsDataRequestWithoutParam() *GetQpsDataRequest {

    return &GetQpsDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/chart:getQpsData",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *GetQpsDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param req: 请求(Required) */
func (r *GetQpsDataRequest) SetReq(req *waf.GetChartReq) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetQpsDataRequest) GetRegionId() string {
    return r.RegionId
}

type GetQpsDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetQpsDataResult `json:"result"`
}

type GetQpsDataResult struct {
    TimeScope []int64 `json:"timeScope"`
    Qps waf.Qps `json:"qps"`
}