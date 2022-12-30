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

type DetachSubUserPolicyRequest struct {

    core.JDCloudRequest

    /* 子用户名  */
    SubUser string `json:"subUser"`

    /* 策略名称  */
    PolicyName string `json:"policyName"`
}

/*
 * param subUser: 子用户名 (Required)
 * param policyName: 策略名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDetachSubUserPolicyRequest(
    subUser string,
    policyName string,
) *DetachSubUserPolicyRequest {

	return &DetachSubUserPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/subUser/{subUser}:detachSubUserPolicy",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        SubUser: subUser,
        PolicyName: policyName,
	}
}

/*
 * param subUser: 子用户名 (Required)
 * param policyName: 策略名称 (Required)
 */
func NewDetachSubUserPolicyRequestWithAllParams(
    subUser string,
    policyName string,
) *DetachSubUserPolicyRequest {

    return &DetachSubUserPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/subUser/{subUser}:detachSubUserPolicy",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        SubUser: subUser,
        PolicyName: policyName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDetachSubUserPolicyRequestWithoutParam() *DetachSubUserPolicyRequest {

    return &DetachSubUserPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/subUser/{subUser}:detachSubUserPolicy",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param subUser: 子用户名(Required) */
func (r *DetachSubUserPolicyRequest) SetSubUser(subUser string) {
    r.SubUser = subUser
}

/* param policyName: 策略名称(Required) */
func (r *DetachSubUserPolicyRequest) SetPolicyName(policyName string) {
    r.PolicyName = policyName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DetachSubUserPolicyRequest) GetRegionId() string {
    return ""
}

type DetachSubUserPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DetachSubUserPolicyResult `json:"result"`
}

type DetachSubUserPolicyResult struct {
}