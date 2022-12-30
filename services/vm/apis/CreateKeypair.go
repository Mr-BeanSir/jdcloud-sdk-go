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

type CreateKeypairRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 密钥对名称，需要全局唯一。
只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
  */
    KeyName string `json:"keyName"`

}

/*
 * param regionId: 地域ID。 (Required)
 * param keyName: 密钥对名称，需要全局唯一。
只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateKeypairRequest(
    regionId string,
    keyName string,
) *CreateKeypairRequest {

	return &CreateKeypairRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/keypairs",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        KeyName: keyName,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param keyName: 密钥对名称，需要全局唯一。
只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
 (Required)
 */
func NewCreateKeypairRequestWithAllParams(
    regionId string,
    keyName string,
) *CreateKeypairRequest {

    return &CreateKeypairRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        KeyName: keyName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateKeypairRequestWithoutParam() *CreateKeypairRequest {

    return &CreateKeypairRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *CreateKeypairRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param keyName: 密钥对名称，需要全局唯一。
只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
(Required) */
func (r *CreateKeypairRequest) SetKeyName(keyName string) {
    r.KeyName = keyName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateKeypairRequest) GetRegionId() string {
    return r.RegionId
}

type CreateKeypairResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateKeypairResult `json:"result"`
}

type CreateKeypairResult struct {
    KeyName string `json:"keyName"`
    PrivateKey string `json:"privateKey"`
    KeyFingerprint string `json:"keyFingerprint"`
}