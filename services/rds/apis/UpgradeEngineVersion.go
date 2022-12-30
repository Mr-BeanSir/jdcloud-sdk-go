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

type UpgradeEngineVersionRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 计划开始升级的时间，1：立即开始升级，2：维护时间窗口升级，0：取消升级  */
    UpgradeSchedule int `json:"upgradeSchedule"`

    /* 升级到的新版本，默认为当前实例可升级到的最新版本 (Optional) */
    NewVersion *string `json:"newVersion"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param upgradeSchedule: 计划开始升级的时间，1：立即开始升级，2：维护时间窗口升级，0：取消升级 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpgradeEngineVersionRequest(
    regionId string,
    instanceId string,
    upgradeSchedule int,
) *UpgradeEngineVersionRequest {

	return &UpgradeEngineVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:upgradeEngineVersion",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        UpgradeSchedule: upgradeSchedule,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param upgradeSchedule: 计划开始升级的时间，1：立即开始升级，2：维护时间窗口升级，0：取消升级 (Required)
 * param newVersion: 升级到的新版本，默认为当前实例可升级到的最新版本 (Optional)
 */
func NewUpgradeEngineVersionRequestWithAllParams(
    regionId string,
    instanceId string,
    upgradeSchedule int,
    newVersion *string,
) *UpgradeEngineVersionRequest {

    return &UpgradeEngineVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:upgradeEngineVersion",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        UpgradeSchedule: upgradeSchedule,
        NewVersion: newVersion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpgradeEngineVersionRequestWithoutParam() *UpgradeEngineVersionRequest {

    return &UpgradeEngineVersionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:upgradeEngineVersion",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *UpgradeEngineVersionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *UpgradeEngineVersionRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param upgradeSchedule: 计划开始升级的时间，1：立即开始升级，2：维护时间窗口升级，0：取消升级(Required) */
func (r *UpgradeEngineVersionRequest) SetUpgradeSchedule(upgradeSchedule int) {
    r.UpgradeSchedule = upgradeSchedule
}

/* param newVersion: 升级到的新版本，默认为当前实例可升级到的最新版本(Optional) */
func (r *UpgradeEngineVersionRequest) SetNewVersion(newVersion string) {
    r.NewVersion = &newVersion
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpgradeEngineVersionRequest) GetRegionId() string {
    return r.RegionId
}

type UpgradeEngineVersionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpgradeEngineVersionResult `json:"result"`
}

type UpgradeEngineVersionResult struct {
}