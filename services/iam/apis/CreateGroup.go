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
    iam "github.com/jdcloud-api/jdcloud-sdk-go/services/iam/models"
)

type CreateGroupRequest struct {

    core.JDCloudRequest

    /*   */
    CreateGroupInfo *iam.CreateGroupInfo `json:"createGroupInfo"`
}

/*
 * param createGroupInfo:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateGroupRequest(
    createGroupInfo *iam.CreateGroupInfo,
) *CreateGroupRequest {

	return &CreateGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/group",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CreateGroupInfo: createGroupInfo,
	}
}

/*
 * param createGroupInfo:  (Required)
 */
func NewCreateGroupRequestWithAllParams(
    createGroupInfo *iam.CreateGroupInfo,
) *CreateGroupRequest {

    return &CreateGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/group",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CreateGroupInfo: createGroupInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateGroupRequestWithoutParam() *CreateGroupRequest {

    return &CreateGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/group",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param createGroupInfo: (Required) */
func (r *CreateGroupRequest) SetCreateGroupInfo(createGroupInfo *iam.CreateGroupInfo) {
    r.CreateGroupInfo = createGroupInfo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateGroupRequest) GetRegionId() string {
    return ""
}

type CreateGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateGroupResult `json:"result"`
}

type CreateGroupResult struct {
    Group iam.CreateGroupRes `json:"group"`
}