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
	"github.com/Mr-BeanSir/jdcloud-sdk-go/core"
	starshield "github.com/Mr-BeanSir/jdcloud-sdk-go/services/starshield/models"
)

type RuleGroupDetailsRequest struct {
	core.JDCloudRequest

	/*   */
	Zone_identifier string `json:"zone_identifier"`

	/*   */
	Package_identifier string `json:"package_identifier"`

	/*   */
	Identifier string `json:"identifier"`
}

/*
 * param zone_identifier:  (Required)
 * param package_identifier:  (Required)
 * param identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRuleGroupDetailsRequest(
	zone_identifier string,
	package_identifier string,
	identifier string,
) *RuleGroupDetailsRequest {

	return &RuleGroupDetailsRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/firewall$$waf$$packages/{package_identifier}/groups/{identifier}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		Zone_identifier:    zone_identifier,
		Package_identifier: package_identifier,
		Identifier:         identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param package_identifier:  (Required)
 * param identifier:  (Required)
 */
func NewRuleGroupDetailsRequestWithAllParams(
	zone_identifier string,
	package_identifier string,
	identifier string,
) *RuleGroupDetailsRequest {

	return &RuleGroupDetailsRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/firewall$$waf$$packages/{package_identifier}/groups/{identifier}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		Zone_identifier:    zone_identifier,
		Package_identifier: package_identifier,
		Identifier:         identifier,
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewRuleGroupDetailsRequestWithoutParam() *RuleGroupDetailsRequest {

	return &RuleGroupDetailsRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/firewall$$waf$$packages/{package_identifier}/groups/{identifier}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/* param zone_identifier: (Required) */
func (r *RuleGroupDetailsRequest) SetZone_identifier(zone_identifier string) {
	r.Zone_identifier = zone_identifier
}

/* param package_identifier: (Required) */
func (r *RuleGroupDetailsRequest) SetPackage_identifier(package_identifier string) {
	r.Package_identifier = package_identifier
}

/* param identifier: (Required) */
func (r *RuleGroupDetailsRequest) SetIdentifier(identifier string) {
	r.Identifier = identifier
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RuleGroupDetailsRequest) GetRegionId() string {
	return ""
}

type RuleGroupDetailsResponse struct {
	RequestID string                 `json:"requestId"`
	Error     core.ErrorResponse     `json:"error"`
	Result    RuleGroupDetailsResult `json:"result"`
}

type RuleGroupDetailsResult struct {
	Data starshield.WAFRuleGroup `json:"data"`
}
