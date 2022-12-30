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
    apigateway "github.com/jdcloud-api/jdcloud-sdk-go/services/apigateway/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeIsDeployApiGroupsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* deployStatus - 发布状态，已发布：1，未发布：0
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIsDeployApiGroupsRequest(
    regionId string,
) *DescribeIsDeployApiGroupsRequest {

	return &DescribeIsDeployApiGroupsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups:isDeploy",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param filters: deployStatus - 发布状态，已发布：1，未发布：0
 (Optional)
 */
func NewDescribeIsDeployApiGroupsRequestWithAllParams(
    regionId string,
    filters []common.Filter,
) *DescribeIsDeployApiGroupsRequest {

    return &DescribeIsDeployApiGroupsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups:isDeploy",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIsDeployApiGroupsRequestWithoutParam() *DescribeIsDeployApiGroupsRequest {

    return &DescribeIsDeployApiGroupsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups:isDeploy",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeIsDeployApiGroupsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param filters: deployStatus - 发布状态，已发布：1，未发布：0
(Optional) */
func (r *DescribeIsDeployApiGroupsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIsDeployApiGroupsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIsDeployApiGroupsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIsDeployApiGroupsResult `json:"result"`
}

type DescribeIsDeployApiGroupsResult struct {
    ApiGroups []apigateway.ApiGroup `json:"apiGroups"`
    TotalCount int `json:"totalCount"`
}