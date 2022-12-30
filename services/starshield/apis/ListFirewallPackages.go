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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type ListFirewallPackagesRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /* Name of the firewall package (Optional) */
    Name *string `json:"name"`

    /* Page number of paginated results (Optional) */
    Page *int `json:"page"`

    /* 每页的包数 (Optional) */
    Per_page *int `json:"per_page"`

    /* 按字段对包排序 (Optional) */
    Order *string `json:"order"`

    /* asc - 升序；desc - 降序 (Optional) */
    Direction *string `json:"direction"`

    /* 是否匹配所有搜索要求或至少一个（任何） (Optional) */
    Match *string `json:"match"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListFirewallPackagesRequest(
    zone_identifier string,
) *ListFirewallPackagesRequest {

	return &ListFirewallPackagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/firewall$$waf$$packages",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param name: Name of the firewall package (Optional)
 * param page: Page number of paginated results (Optional)
 * param per_page: 每页的包数 (Optional)
 * param order: 按字段对包排序 (Optional)
 * param direction: asc - 升序；desc - 降序 (Optional)
 * param match: 是否匹配所有搜索要求或至少一个（任何） (Optional)
 */
func NewListFirewallPackagesRequestWithAllParams(
    zone_identifier string,
    name *string,
    page *int,
    per_page *int,
    order *string,
    direction *string,
    match *string,
) *ListFirewallPackagesRequest {

    return &ListFirewallPackagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/firewall$$waf$$packages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Name: name,
        Page: page,
        Per_page: per_page,
        Order: order,
        Direction: direction,
        Match: match,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListFirewallPackagesRequestWithoutParam() *ListFirewallPackagesRequest {

    return &ListFirewallPackagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/firewall$$waf$$packages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *ListFirewallPackagesRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

/* param name: Name of the firewall package(Optional) */
func (r *ListFirewallPackagesRequest) SetName(name string) {
    r.Name = &name
}

/* param page: Page number of paginated results(Optional) */
func (r *ListFirewallPackagesRequest) SetPage(page int) {
    r.Page = &page
}

/* param per_page: 每页的包数(Optional) */
func (r *ListFirewallPackagesRequest) SetPer_page(per_page int) {
    r.Per_page = &per_page
}

/* param order: 按字段对包排序(Optional) */
func (r *ListFirewallPackagesRequest) SetOrder(order string) {
    r.Order = &order
}

/* param direction: asc - 升序；desc - 降序(Optional) */
func (r *ListFirewallPackagesRequest) SetDirection(direction string) {
    r.Direction = &direction
}

/* param match: 是否匹配所有搜索要求或至少一个（任何）(Optional) */
func (r *ListFirewallPackagesRequest) SetMatch(match string) {
    r.Match = &match
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListFirewallPackagesRequest) GetRegionId() string {
    return ""
}

type ListFirewallPackagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListFirewallPackagesResult `json:"result"`
}

type ListFirewallPackagesResult struct {
    DataList []starshield.WAFRulePackage `json:"dataList"`
}