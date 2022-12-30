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

type ListDomainsRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 请求  */
    Req *waf.ListDomains `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例Id (Required)
 * param req: 请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListDomainsRequest(
    regionId string,
    wafInstanceId string,
    req *waf.ListDomains,
) *ListDomainsRequest {

	return &ListDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/domain:list",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        WafInstanceId: wafInstanceId,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例Id (Required)
 * param req: 请求 (Required)
 */
func NewListDomainsRequestWithAllParams(
    regionId string,
    wafInstanceId string,
    req *waf.ListDomains,
) *ListDomainsRequest {

    return &ListDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/domain:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        WafInstanceId: wafInstanceId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListDomainsRequestWithoutParam() *ListDomainsRequest {

    return &ListDomainsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/domain:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *ListDomainsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param wafInstanceId: 实例Id(Required) */
func (r *ListDomainsRequest) SetWafInstanceId(wafInstanceId string) {
    r.WafInstanceId = wafInstanceId
}

/* param req: 请求(Required) */
func (r *ListDomainsRequest) SetReq(req *waf.ListDomains) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListDomainsRequest) GetRegionId() string {
    return r.RegionId
}

type ListDomainsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListDomainsResult `json:"result"`
}

type ListDomainsResult struct {
    WafInstanceId string `json:"wafInstanceId"`
    List []string `json:"list"`
    PageIndex int `json:"pageIndex"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    MaxLimit int `json:"maxLimit"`
}