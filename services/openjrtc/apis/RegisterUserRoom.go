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

type RegisterUserRoomRequest struct {

    core.JDCloudRequest

    /* 业务接入方定义的房间号 (Optional) */
    UserRoomId *string `json:"userRoomId"`

    /* 房间名称 (Optional) */
    RoomName *string `json:"roomName"`

    /* 应用ID (Optional) */
    AppId *string `json:"appId"`

    /* 房间类型 1-小房间(音频单流订阅) 2-大房间(音频固定订阅),默认取控制台APP对应的房间类型 (Optional) */
    RoomType *int `json:"roomType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRegisterUserRoomRequest(
) *RegisterUserRoomRequest {

	return &RegisterUserRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/registerUserRoom",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param userRoomId: 业务接入方定义的房间号 (Optional)
 * param roomName: 房间名称 (Optional)
 * param appId: 应用ID (Optional)
 * param roomType: 房间类型 1-小房间(音频单流订阅) 2-大房间(音频固定订阅),默认取控制台APP对应的房间类型 (Optional)
 */
func NewRegisterUserRoomRequestWithAllParams(
    userRoomId *string,
    roomName *string,
    appId *string,
    roomType *int,
) *RegisterUserRoomRequest {

    return &RegisterUserRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/registerUserRoom",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        UserRoomId: userRoomId,
        RoomName: roomName,
        AppId: appId,
        RoomType: roomType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRegisterUserRoomRequestWithoutParam() *RegisterUserRoomRequest {

    return &RegisterUserRoomRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/registerUserRoom",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param userRoomId: 业务接入方定义的房间号(Optional) */
func (r *RegisterUserRoomRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = &userRoomId
}

/* param roomName: 房间名称(Optional) */
func (r *RegisterUserRoomRequest) SetRoomName(roomName string) {
    r.RoomName = &roomName
}

/* param appId: 应用ID(Optional) */
func (r *RegisterUserRoomRequest) SetAppId(appId string) {
    r.AppId = &appId
}

/* param roomType: 房间类型 1-小房间(音频单流订阅) 2-大房间(音频固定订阅),默认取控制台APP对应的房间类型(Optional) */
func (r *RegisterUserRoomRequest) SetRoomType(roomType int) {
    r.RoomType = &roomType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RegisterUserRoomRequest) GetRegionId() string {
    return ""
}

type RegisterUserRoomResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RegisterUserRoomResult `json:"result"`
}

type RegisterUserRoomResult struct {
    RoomId int64 `json:"roomId"`
    UserRoomId string `json:"userRoomId"`
    RoomName string `json:"roomName"`
    RoomType int `json:"roomType"`
    AppId string `json:"appId"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}