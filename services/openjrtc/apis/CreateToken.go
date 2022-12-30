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

type CreateTokenRequest struct {

    core.JDCloudRequest

    /* appId (Optional) */
    AppId *string `json:"appId"`

    /* appKey (Optional) */
    AppKey *string `json:"appKey"`

    /* 用户id (Optional) */
    UserId *string `json:"userId"`

    /* 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional) */
    UserRoomId *string `json:"userRoomId"`

    /* 时间戳-毫秒 (Optional) */
    Timestamp *int64 `json:"timestamp"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTokenRequest(
) *CreateTokenRequest {

	return &CreateTokenRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/createToken",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param appId: appId (Optional)
 * param appKey: appKey (Optional)
 * param userId: 用户id (Optional)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional)
 * param timestamp: 时间戳-毫秒 (Optional)
 */
func NewCreateTokenRequestWithAllParams(
    appId *string,
    appKey *string,
    userId *string,
    userRoomId *string,
    timestamp *int64,
) *CreateTokenRequest {

    return &CreateTokenRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/createToken",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        AppKey: appKey,
        UserId: userId,
        UserRoomId: userRoomId,
        Timestamp: timestamp,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTokenRequestWithoutParam() *CreateTokenRequest {

    return &CreateTokenRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/createToken",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: appId(Optional) */
func (r *CreateTokenRequest) SetAppId(appId string) {
    r.AppId = &appId
}

/* param appKey: appKey(Optional) */
func (r *CreateTokenRequest) SetAppKey(appKey string) {
    r.AppKey = &appKey
}

/* param userId: 用户id(Optional) */
func (r *CreateTokenRequest) SetUserId(userId string) {
    r.UserId = &userId
}

/* param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号(Optional) */
func (r *CreateTokenRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = &userRoomId
}

/* param timestamp: 时间戳-毫秒(Optional) */
func (r *CreateTokenRequest) SetTimestamp(timestamp int64) {
    r.Timestamp = &timestamp
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTokenRequest) GetRegionId() string {
    return ""
}

type CreateTokenResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTokenResult `json:"result"`
}

type CreateTokenResult struct {
    AppId string `json:"appId"`
    AppKey string `json:"appKey"`
    UserId string `json:"userId"`
    UserRoomId string `json:"userRoomId"`
    Nonce string `json:"nonce"`
    Timestamp int64 `json:"timestamp"`
    Token string `json:"token"`
    Available bool `json:"available"`
}