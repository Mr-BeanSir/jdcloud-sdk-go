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
    sts "github.com/jdcloud-api/jdcloud-sdk-go/services/sts/models"
)

type AssumeRoleRequest struct {

    core.JDCloudRequest

    /* 扮演角色参数  */
    AssumeRoleInfo *sts.AssumeRoleInfo `json:"assumeRoleInfo"`
}

/*
 * param assumeRoleInfo: 扮演角色参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAssumeRoleRequest(
    assumeRoleInfo *sts.AssumeRoleInfo,
) *AssumeRoleRequest {

	return &AssumeRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sessionToken:assumeRole",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AssumeRoleInfo: assumeRoleInfo,
	}
}

/*
 * param assumeRoleInfo: 扮演角色参数 (Required)
 */
func NewAssumeRoleRequestWithAllParams(
    assumeRoleInfo *sts.AssumeRoleInfo,
) *AssumeRoleRequest {

    return &AssumeRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sessionToken:assumeRole",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AssumeRoleInfo: assumeRoleInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAssumeRoleRequestWithoutParam() *AssumeRoleRequest {

    return &AssumeRoleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sessionToken:assumeRole",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param assumeRoleInfo: 扮演角色参数(Required) */
func (r *AssumeRoleRequest) SetAssumeRoleInfo(assumeRoleInfo *sts.AssumeRoleInfo) {
    r.AssumeRoleInfo = assumeRoleInfo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AssumeRoleRequest) GetRegionId() string {
    return ""
}

type AssumeRoleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AssumeRoleResult `json:"result"`
}

type AssumeRoleResult struct {
    Credentials sts.Credentials `json:"credentials"`
}