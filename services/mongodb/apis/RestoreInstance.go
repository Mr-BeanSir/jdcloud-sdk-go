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

type RestoreInstanceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 备份ID  */
    BackupId string `json:"backupId"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param backupId: 备份ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRestoreInstanceRequest(
    regionId string,
    instanceId string,
    backupId string,
) *RestoreInstanceRequest {

	return &RestoreInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/restoreInstance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        BackupId: backupId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param backupId: 备份ID (Required)
 */
func NewRestoreInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
    backupId string,
) *RestoreInstanceRequest {

    return &RestoreInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/restoreInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        BackupId: backupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRestoreInstanceRequestWithoutParam() *RestoreInstanceRequest {

    return &RestoreInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/restoreInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *RestoreInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: Instance ID(Required) */
func (r *RestoreInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param backupId: 备份ID(Required) */
func (r *RestoreInstanceRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestoreInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type RestoreInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestoreInstanceResult `json:"result"`
}

type RestoreInstanceResult struct {
}