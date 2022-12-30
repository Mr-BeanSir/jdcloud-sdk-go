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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
)

type GrantRdsPrivilegeRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 账号名，在同一个实例中账号名不能重复  */
    AccountName string `json:"accountName"`

    /* 账号的访问权限  */
    AccountPrivileges []rds.AccountPrivilege `json:"accountPrivileges"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param accountName: 账号名，在同一个实例中账号名不能重复 (Required)
 * param accountPrivileges: 账号的访问权限 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGrantRdsPrivilegeRequest(
    regionId string,
    instanceId string,
    accountName string,
    accountPrivileges []rds.AccountPrivilege,
) *GrantRdsPrivilegeRequest {

	return &GrantRdsPrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/accounts/{accountName}:grantPrivilege",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AccountName: accountName,
        AccountPrivileges: accountPrivileges,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param accountName: 账号名，在同一个实例中账号名不能重复 (Required)
 * param accountPrivileges: 账号的访问权限 (Required)
 */
func NewGrantRdsPrivilegeRequestWithAllParams(
    regionId string,
    instanceId string,
    accountName string,
    accountPrivileges []rds.AccountPrivilege,
) *GrantRdsPrivilegeRequest {

    return &GrantRdsPrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/accounts/{accountName}:grantPrivilege",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        AccountName: accountName,
        AccountPrivileges: accountPrivileges,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGrantRdsPrivilegeRequestWithoutParam() *GrantRdsPrivilegeRequest {

    return &GrantRdsPrivilegeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydRdsInstances/{instanceId}/accounts/{accountName}:grantPrivilege",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *GrantRdsPrivilegeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *GrantRdsPrivilegeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param accountName: 账号名，在同一个实例中账号名不能重复(Required) */
func (r *GrantRdsPrivilegeRequest) SetAccountName(accountName string) {
    r.AccountName = accountName
}

/* param accountPrivileges: 账号的访问权限(Required) */
func (r *GrantRdsPrivilegeRequest) SetAccountPrivileges(accountPrivileges []rds.AccountPrivilege) {
    r.AccountPrivileges = accountPrivileges
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GrantRdsPrivilegeRequest) GetRegionId() string {
    return r.RegionId
}

type GrantRdsPrivilegeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GrantRdsPrivilegeResult `json:"result"`
}

type GrantRdsPrivilegeResult struct {
}