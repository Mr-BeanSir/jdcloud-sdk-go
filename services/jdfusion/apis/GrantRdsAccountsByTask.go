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
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/models"
)

type GrantRdsAccountsByTaskRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* RDS实例ID  */
    InstId string `json:"instId"`

    /* 账号名称  */
    AccountName string `json:"accountName"`

    /* RDS账号对数据库的权限信息  */
    Info *jdfusion.DbPrivilegeInfo `json:"info"`
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 * param accountName: 账号名称 (Required)
 * param info: RDS账号对数据库的权限信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGrantRdsAccountsByTaskRequest(
    regionId string,
    instId string,
    accountName string,
    info *jdfusion.DbPrivilegeInfo,
) *GrantRdsAccountsByTaskRequest {

	return &GrantRdsAccountsByTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rds_instances/{instId}/accounts/{accountName}:grantByTask",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstId: instId,
        AccountName: accountName,
        Info: info,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 * param accountName: 账号名称 (Required)
 * param info: RDS账号对数据库的权限信息 (Required)
 */
func NewGrantRdsAccountsByTaskRequestWithAllParams(
    regionId string,
    instId string,
    accountName string,
    info *jdfusion.DbPrivilegeInfo,
) *GrantRdsAccountsByTaskRequest {

    return &GrantRdsAccountsByTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/accounts/{accountName}:grantByTask",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstId: instId,
        AccountName: accountName,
        Info: info,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGrantRdsAccountsByTaskRequestWithoutParam() *GrantRdsAccountsByTaskRequest {

    return &GrantRdsAccountsByTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/accounts/{accountName}:grantByTask",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GrantRdsAccountsByTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instId: RDS实例ID(Required) */
func (r *GrantRdsAccountsByTaskRequest) SetInstId(instId string) {
    r.InstId = instId
}

/* param accountName: 账号名称(Required) */
func (r *GrantRdsAccountsByTaskRequest) SetAccountName(accountName string) {
    r.AccountName = accountName
}

/* param info: RDS账号对数据库的权限信息(Required) */
func (r *GrantRdsAccountsByTaskRequest) SetInfo(info *jdfusion.DbPrivilegeInfo) {
    r.Info = info
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GrantRdsAccountsByTaskRequest) GetRegionId() string {
    return r.RegionId
}

type GrantRdsAccountsByTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GrantRdsAccountsByTaskResult `json:"result"`
}

type GrantRdsAccountsByTaskResult struct {
    Cloud jdfusion.ResourceTFInfo `json:"cloud"`
}