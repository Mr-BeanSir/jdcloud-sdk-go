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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type GetHTTP2EdgePrioritizationSettingRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetHTTP2EdgePrioritizationSettingRequest(
    zone_identifier string,
) *GetHTTP2EdgePrioritizationSettingRequest {

	return &GetHTTP2EdgePrioritizationSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/settings$$h2_prioritization",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 */
func NewGetHTTP2EdgePrioritizationSettingRequestWithAllParams(
    zone_identifier string,
) *GetHTTP2EdgePrioritizationSettingRequest {

    return &GetHTTP2EdgePrioritizationSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/settings$$h2_prioritization",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetHTTP2EdgePrioritizationSettingRequestWithoutParam() *GetHTTP2EdgePrioritizationSettingRequest {

    return &GetHTTP2EdgePrioritizationSettingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/settings$$h2_prioritization",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *GetHTTP2EdgePrioritizationSettingRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetHTTP2EdgePrioritizationSettingRequest) GetRegionId() string {
    return ""
}

type GetHTTP2EdgePrioritizationSettingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetHTTP2EdgePrioritizationSettingResult `json:"result"`
}

type GetHTTP2EdgePrioritizationSettingResult struct {
    Data starshield.HTTP2EdgePrioritization `json:"data"`
}