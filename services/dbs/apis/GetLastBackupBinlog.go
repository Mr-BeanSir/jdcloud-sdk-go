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
    dbs "github.com/jdcloud-api/jdcloud-sdk-go/services/dbs/models"
)

type GetLastBackupBinlogRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划ID  */
    BackupPlanId string `json:"backupPlanId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetLastBackupBinlogRequest(
    regionId string,
    backupPlanId string,
) *GetLastBackupBinlogRequest {

	return &GetLastBackupBinlogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:getLastBackupBinlog",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        BackupPlanId: backupPlanId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 */
func NewGetLastBackupBinlogRequestWithAllParams(
    regionId string,
    backupPlanId string,
) *GetLastBackupBinlogRequest {

    return &GetLastBackupBinlogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:getLastBackupBinlog",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BackupPlanId: backupPlanId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetLastBackupBinlogRequestWithoutParam() *GetLastBackupBinlogRequest {

    return &GetLastBackupBinlogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:getLastBackupBinlog",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *GetLastBackupBinlogRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupPlanId: 备份计划ID(Required) */
func (r *GetLastBackupBinlogRequest) SetBackupPlanId(backupPlanId string) {
    r.BackupPlanId = backupPlanId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetLastBackupBinlogRequest) GetRegionId() string {
    return r.RegionId
}

type GetLastBackupBinlogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetLastBackupBinlogResult `json:"result"`
}

type GetLastBackupBinlogResult struct {
    LastBackupBinlog dbs.BackupBriefInfo `json:"lastBackupBinlog"`
}