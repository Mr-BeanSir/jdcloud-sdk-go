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
	"github.com/Mr-BeanSir/jdcloud-sdk-go/core"
	starshield "github.com/Mr-BeanSir/jdcloud-sdk-go/services/starshield/models"
)

type ChangeAlwaysOnlineSettingRequest struct {
	core.JDCloudRequest

	/*   */
	Zone_identifier string `json:"zone_identifier"`

	/* on - 开启；off - 关闭 (Optional) */
	Value *string `json:"value"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewChangeAlwaysOnlineSettingRequest(
	zone_identifier string,
) *ChangeAlwaysOnlineSettingRequest {

	return &ChangeAlwaysOnlineSettingRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/settings$$always_online",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
		Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param value: on - 开启；off - 关闭 (Optional)
 */
func NewChangeAlwaysOnlineSettingRequestWithAllParams(
	zone_identifier string,
	value *string,
) *ChangeAlwaysOnlineSettingRequest {

	return &ChangeAlwaysOnlineSettingRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/settings$$always_online",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
		Zone_identifier: zone_identifier,
		Value:           value,
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewChangeAlwaysOnlineSettingRequestWithoutParam() *ChangeAlwaysOnlineSettingRequest {

	return &ChangeAlwaysOnlineSettingRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/settings$$always_online",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
	}
}

/* param zone_identifier: (Required) */
func (r *ChangeAlwaysOnlineSettingRequest) SetZone_identifier(zone_identifier string) {
	r.Zone_identifier = zone_identifier
}

/* param value: on - 开启；off - 关闭(Optional) */
func (r *ChangeAlwaysOnlineSettingRequest) SetValue(value string) {
	r.Value = &value
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ChangeAlwaysOnlineSettingRequest) GetRegionId() string {
	return ""
}

type ChangeAlwaysOnlineSettingResponse struct {
	RequestID string                          `json:"requestId"`
	Error     core.ErrorResponse              `json:"error"`
	Result    ChangeAlwaysOnlineSettingResult `json:"result"`
}

type ChangeAlwaysOnlineSettingResult struct {
	Data starshield.AlwaysOnlineMode `json:"data"`
}
