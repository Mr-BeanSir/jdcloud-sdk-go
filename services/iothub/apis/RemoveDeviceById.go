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

type RemoveDeviceByIdRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 设备Id  */
    DeviceId string `json:"deviceId"`
}

/*
 * param regionId: 设备归属的实例所在区域 (Required)
 * param deviceId: 设备Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemoveDeviceByIdRequest(
    regionId string,
    deviceId string,
) *RemoveDeviceByIdRequest {

	return &RemoveDeviceByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/device/{deviceId}:deleteById",
			Method:  "DELETE",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DeviceId: deviceId,
	}
}

/*
 * param regionId: 设备归属的实例所在区域 (Required)
 * param deviceId: 设备Id (Required)
 */
func NewRemoveDeviceByIdRequestWithAllParams(
    regionId string,
    deviceId string,
) *RemoveDeviceByIdRequest {

    return &RemoveDeviceByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/device/{deviceId}:deleteById",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DeviceId: deviceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemoveDeviceByIdRequestWithoutParam() *RemoveDeviceByIdRequest {

    return &RemoveDeviceByIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/device/{deviceId}:deleteById",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *RemoveDeviceByIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param deviceId: 设备Id(Required) */
func (r *RemoveDeviceByIdRequest) SetDeviceId(deviceId string) {
    r.DeviceId = deviceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemoveDeviceByIdRequest) GetRegionId() string {
    return r.RegionId
}

type RemoveDeviceByIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemoveDeviceByIdResult `json:"result"`
}

type RemoveDeviceByIdResult struct {
}