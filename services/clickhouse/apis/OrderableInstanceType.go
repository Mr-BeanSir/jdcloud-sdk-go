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
    clickhouse "github.com/jdcloud-api/jdcloud-sdk-go/services/clickhouse/models"
)

type OrderableInstanceTypeRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 地域代码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOrderableInstanceTypeRequest(
    regionId string,
) *OrderableInstanceTypeRequest {

	return &OrderableInstanceTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/OrderableInstanceType",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 */
func NewOrderableInstanceTypeRequestWithAllParams(
    regionId string,
) *OrderableInstanceTypeRequest {

    return &OrderableInstanceTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/OrderableInstanceType",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOrderableInstanceTypeRequestWithoutParam() *OrderableInstanceTypeRequest {

    return &OrderableInstanceTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/OrderableInstanceType",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *OrderableInstanceTypeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OrderableInstanceTypeRequest) GetRegionId() string {
    return r.RegionId
}

type OrderableInstanceTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OrderableInstanceTypeResult `json:"result"`
}

type OrderableInstanceTypeResult struct {
    EngineStatus int `json:"engineStatus"`
    OrderableAZs []clickhouse.Az `json:"orderableAZs"`
}