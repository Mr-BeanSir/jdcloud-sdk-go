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
    openjrtc "github.com/jdcloud-api/jdcloud-sdk-go/services/openjrtc/models"
)

type AddPushStreamRuleRequest struct {

    core.JDCloudRequest

    /* 应用ID (Optional) */
    AppId *string `json:"appId"`

    /* 推流规则 (Optional) */
    Rules []openjrtc.PushStreamRule `json:"rules"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddPushStreamRuleRequest(
) *AddPushStreamRuleRequest {

	return &AddPushStreamRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/addPushStreamRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param appId: 应用ID (Optional)
 * param rules: 推流规则 (Optional)
 */
func NewAddPushStreamRuleRequestWithAllParams(
    appId *string,
    rules []openjrtc.PushStreamRule,
) *AddPushStreamRuleRequest {

    return &AddPushStreamRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/addPushStreamRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        Rules: rules,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddPushStreamRuleRequestWithoutParam() *AddPushStreamRuleRequest {

    return &AddPushStreamRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/addPushStreamRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Optional) */
func (r *AddPushStreamRuleRequest) SetAppId(appId string) {
    r.AppId = &appId
}

/* param rules: 推流规则(Optional) */
func (r *AddPushStreamRuleRequest) SetRules(rules []openjrtc.PushStreamRule) {
    r.Rules = rules
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddPushStreamRuleRequest) GetRegionId() string {
    return ""
}

type AddPushStreamRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddPushStreamRuleResult `json:"result"`
}

type AddPushStreamRuleResult struct {
}