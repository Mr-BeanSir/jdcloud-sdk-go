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

type EnableWafBlackRulesRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /*  (Optional) */
    Ids []string `json:"ids"`

    /*  (Optional) */
    RuleType *string `json:"ruleType"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableWafBlackRulesRequest(
    domain string,
) *EnableWafBlackRulesRequest {

	return &EnableWafBlackRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/wafBlackRule:enable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param ids:  (Optional)
 * param ruleType:  (Optional)
 */
func NewEnableWafBlackRulesRequestWithAllParams(
    domain string,
    ids []string,
    ruleType *string,
) *EnableWafBlackRulesRequest {

    return &EnableWafBlackRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafBlackRule:enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Ids: ids,
        RuleType: ruleType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableWafBlackRulesRequestWithoutParam() *EnableWafBlackRulesRequest {

    return &EnableWafBlackRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafBlackRule:enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *EnableWafBlackRulesRequest) SetDomain(domain string) {
    r.Domain = domain
}
/* param ids: (Optional) */
func (r *EnableWafBlackRulesRequest) SetIds(ids []string) {
    r.Ids = ids
}
/* param ruleType: (Optional) */
func (r *EnableWafBlackRulesRequest) SetRuleType(ruleType string) {
    r.RuleType = &ruleType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableWafBlackRulesRequest) GetRegionId() string {
    return ""
}

type EnableWafBlackRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableWafBlackRulesResult `json:"result"`
}

type EnableWafBlackRulesResult struct {
}