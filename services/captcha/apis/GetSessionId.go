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
    captcha "github.com/jdcloud-api/jdcloud-sdk-go/services/captcha/models"
)

type GetSessionIdRequest struct {

    core.JDCloudRequest

    /* 应用id (Optional) */
    AppId *int64 `json:"appId"`

    /* 场景id (Optional) */
    SceneId *int64 `json:"sceneId"`

    /* 密钥，从界面获取 (Optional) */
    Secret *string `json:"secret"`

    /* uuid，ios客户端传openudid, android客户端传androidid, m, pc, wxapp客户端此值为空即可 (Optional) */
    Uuid *string `json:"uuid"`

    /* 客户端ip (Optional) */
    Ip *string `json:"ip"`

    /* 客户端userAgent (Optional) */
    UserAgent *string `json:"userAgent"`

    /* 指纹，ios和android客户端(clientType)从sdk获取, m, pc, wxapp客户端此值为空即可 (Optional) */
    FingerPrint *string `json:"fingerPrint"`

    /* 客户端类型, android, ios, pc, wxmapp, m (Optional) */
    ClientType *string `json:"clientType"`

    /* 客户端版本，用户端app版本，可选 (Optional) */
    ClientVersion *string `json:"clientVersion"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSessionIdRequest(
) *GetSessionIdRequest {

	return &GetSessionIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/captcha:getsessionid",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param appId: 应用id (Optional)
 * param sceneId: 场景id (Optional)
 * param secret: 密钥，从界面获取 (Optional)
 * param uuid: uuid，ios客户端传openudid, android客户端传androidid, m, pc, wxapp客户端此值为空即可 (Optional)
 * param ip: 客户端ip (Optional)
 * param userAgent: 客户端userAgent (Optional)
 * param fingerPrint: 指纹，ios和android客户端(clientType)从sdk获取, m, pc, wxapp客户端此值为空即可 (Optional)
 * param clientType: 客户端类型, android, ios, pc, wxmapp, m (Optional)
 * param clientVersion: 客户端版本，用户端app版本，可选 (Optional)
 */
func NewGetSessionIdRequestWithAllParams(
    appId *int64,
    sceneId *int64,
    secret *string,
    uuid *string,
    ip *string,
    userAgent *string,
    fingerPrint *string,
    clientType *string,
    clientVersion *string,
) *GetSessionIdRequest {

    return &GetSessionIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/captcha:getsessionid",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        SceneId: sceneId,
        Secret: secret,
        Uuid: uuid,
        Ip: ip,
        UserAgent: userAgent,
        FingerPrint: fingerPrint,
        ClientType: clientType,
        ClientVersion: clientVersion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSessionIdRequestWithoutParam() *GetSessionIdRequest {

    return &GetSessionIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/captcha:getsessionid",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Optional) */
func (r *GetSessionIdRequest) SetAppId(appId int64) {
    r.AppId = &appId
}

/* param sceneId: 场景id(Optional) */
func (r *GetSessionIdRequest) SetSceneId(sceneId int64) {
    r.SceneId = &sceneId
}

/* param secret: 密钥，从界面获取(Optional) */
func (r *GetSessionIdRequest) SetSecret(secret string) {
    r.Secret = &secret
}

/* param uuid: uuid，ios客户端传openudid, android客户端传androidid, m, pc, wxapp客户端此值为空即可(Optional) */
func (r *GetSessionIdRequest) SetUuid(uuid string) {
    r.Uuid = &uuid
}

/* param ip: 客户端ip(Optional) */
func (r *GetSessionIdRequest) SetIp(ip string) {
    r.Ip = &ip
}

/* param userAgent: 客户端userAgent(Optional) */
func (r *GetSessionIdRequest) SetUserAgent(userAgent string) {
    r.UserAgent = &userAgent
}

/* param fingerPrint: 指纹，ios和android客户端(clientType)从sdk获取, m, pc, wxapp客户端此值为空即可(Optional) */
func (r *GetSessionIdRequest) SetFingerPrint(fingerPrint string) {
    r.FingerPrint = &fingerPrint
}

/* param clientType: 客户端类型, android, ios, pc, wxmapp, m(Optional) */
func (r *GetSessionIdRequest) SetClientType(clientType string) {
    r.ClientType = &clientType
}

/* param clientVersion: 客户端版本，用户端app版本，可选(Optional) */
func (r *GetSessionIdRequest) SetClientVersion(clientVersion string) {
    r.ClientVersion = &clientVersion
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSessionIdRequest) GetRegionId() string {
    return ""
}

type GetSessionIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSessionIdResult `json:"result"`
}

type GetSessionIdResult struct {
    Data captcha.SessionDataResp `json:"data"`
}