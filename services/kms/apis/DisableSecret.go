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

type DisableSecretRequest struct {

    core.JDCloudRequest

    /* 机密ID  */
    SecretId string `json:"secretId"`
}

/*
 * param secretId: 机密ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableSecretRequest(
    secretId string,
) *DisableSecretRequest {

	return &DisableSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/secret/{secretId}:disable",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        SecretId: secretId,
	}
}

/*
 * param secretId: 机密ID (Required)
 */
func NewDisableSecretRequestWithAllParams(
    secretId string,
) *DisableSecretRequest {

    return &DisableSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}:disable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        SecretId: secretId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableSecretRequestWithoutParam() *DisableSecretRequest {

    return &DisableSecretRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/secret/{secretId}:disable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param secretId: 机密ID(Required) */
func (r *DisableSecretRequest) SetSecretId(secretId string) {
    r.SecretId = secretId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableSecretRequest) GetRegionId() string {
    return ""
}

type DisableSecretResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableSecretResult `json:"result"`
}

type DisableSecretResult struct {
}