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

type CheckAuthExistRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    AccessKey string `json:"accessKey"`

    /*   */
    AuthUserType string `json:"authUserType"`
}

/*
 * param regionId: 地域ID (Required)
 * param accessKey:  (Required)
 * param authUserType:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckAuthExistRequest(
    regionId string,
    accessKey string,
    authUserType string,
) *CheckAuthExistRequest {

	return &CheckAuthExistRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/accessAuths:checkAccessKeyExist",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AccessKey: accessKey,
        AuthUserType: authUserType,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param accessKey:  (Required)
 * param authUserType:  (Required)
 */
func NewCheckAuthExistRequestWithAllParams(
    regionId string,
    accessKey string,
    authUserType string,
) *CheckAuthExistRequest {

    return &CheckAuthExistRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessAuths:checkAccessKeyExist",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AccessKey: accessKey,
        AuthUserType: authUserType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckAuthExistRequestWithoutParam() *CheckAuthExistRequest {

    return &CheckAuthExistRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessAuths:checkAccessKeyExist",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CheckAuthExistRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param accessKey: (Required) */
func (r *CheckAuthExistRequest) SetAccessKey(accessKey string) {
    r.AccessKey = accessKey
}

/* param authUserType: (Required) */
func (r *CheckAuthExistRequest) SetAuthUserType(authUserType string) {
    r.AuthUserType = authUserType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckAuthExistRequest) GetRegionId() string {
    return r.RegionId
}

type CheckAuthExistResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckAuthExistResult `json:"result"`
}

type CheckAuthExistResult struct {
    AccessAuthId string `json:"accessAuthId"`
}