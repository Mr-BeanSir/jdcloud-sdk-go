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

type RegisterUserRequest struct {

    core.JDCloudRequest

    /* 应用ID (Optional) */
    AppId *string `json:"appId"`

    /* 用户名称 (Optional) */
    UserName *string `json:"userName"`

    /* 业务接入方用户体系定义的userId (Optional) */
    UserId *string `json:"userId"`

    /* 是否临时用户 (Optional) */
    Temporary *bool `json:"temporary"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRegisterUserRequest(
) *RegisterUserRequest {

	return &RegisterUserRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/registerUser",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param appId: 应用ID (Optional)
 * param userName: 用户名称 (Optional)
 * param userId: 业务接入方用户体系定义的userId (Optional)
 * param temporary: 是否临时用户 (Optional)
 */
func NewRegisterUserRequestWithAllParams(
    appId *string,
    userName *string,
    userId *string,
    temporary *bool,
) *RegisterUserRequest {

    return &RegisterUserRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/registerUser",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        UserName: userName,
        UserId: userId,
        Temporary: temporary,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRegisterUserRequestWithoutParam() *RegisterUserRequest {

    return &RegisterUserRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/registerUser",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Optional) */
func (r *RegisterUserRequest) SetAppId(appId string) {
    r.AppId = &appId
}

/* param userName: 用户名称(Optional) */
func (r *RegisterUserRequest) SetUserName(userName string) {
    r.UserName = &userName
}

/* param userId: 业务接入方用户体系定义的userId(Optional) */
func (r *RegisterUserRequest) SetUserId(userId string) {
    r.UserId = &userId
}

/* param temporary: 是否临时用户(Optional) */
func (r *RegisterUserRequest) SetTemporary(temporary bool) {
    r.Temporary = &temporary
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RegisterUserRequest) GetRegionId() string {
    return ""
}

type RegisterUserResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RegisterUserResult `json:"result"`
}

type RegisterUserResult struct {
    AppId string `json:"appId"`
    UserId string `json:"userId"`
    PeerId int64 `json:"peerId"`
}