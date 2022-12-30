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

type DeleteUserAccessKeyRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* accessKey  */
    AccessKey string `json:"accessKey"`
}

/*
 * param regionId: Region ID (Required)
 * param accessKey: accessKey (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteUserAccessKeyRequest(
    regionId string,
    accessKey string,
) *DeleteUserAccessKeyRequest {

	return &DeleteUserAccessKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/userAccessKey/{accessKey}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AccessKey: accessKey,
	}
}

/*
 * param regionId: Region ID (Required)
 * param accessKey: accessKey (Required)
 */
func NewDeleteUserAccessKeyRequestWithAllParams(
    regionId string,
    accessKey string,
) *DeleteUserAccessKeyRequest {

    return &DeleteUserAccessKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userAccessKey/{accessKey}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AccessKey: accessKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteUserAccessKeyRequestWithoutParam() *DeleteUserAccessKeyRequest {

    return &DeleteUserAccessKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userAccessKey/{accessKey}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteUserAccessKeyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param accessKey: accessKey(Required) */
func (r *DeleteUserAccessKeyRequest) SetAccessKey(accessKey string) {
    r.AccessKey = accessKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteUserAccessKeyRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteUserAccessKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteUserAccessKeyResult `json:"result"`
}

type DeleteUserAccessKeyResult struct {
}