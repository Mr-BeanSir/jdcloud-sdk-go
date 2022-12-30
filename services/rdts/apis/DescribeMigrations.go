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
    rdts "github.com/jdcloud-api/jdcloud-sdk-go/services/rdts/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeMigrationsRequest struct {

    core.JDCloudRequest

    /* 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1)  */
    RegionId string `json:"regionId"`

    /* 页码：取值范围[1,∞)，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小：取值范围[10,100]，默认为10 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 过滤条件：
- instanceIds：迁移任务ID，精确匹配，可选择多个
- migrationName：迁移任务名称，模糊匹配
- migrationStatuses：迁移任务状态，精确匹配，可选择多个（具体可参考迁移任务详情中的迁移状态）
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* 排序属性：
- createdTime：按创建时间排序(asc表示按时间正序，desc表示按时间倒序)
 (Optional) */
    Sorts []common.Sort `json:"sorts"`
}

/*
 * param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMigrationsRequest(
    regionId string,
) *DescribeMigrationsRequest {

	return &DescribeMigrationsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1) (Required)
 * param pageNumber: 页码：取值范围[1,∞)，默认为1 (Optional)
 * param pageSize: 分页大小：取值范围[10,100]，默认为10 (Optional)
 * param filters: 过滤条件：
- instanceIds：迁移任务ID，精确匹配，可选择多个
- migrationName：迁移任务名称，模糊匹配
- migrationStatuses：迁移任务状态，精确匹配，可选择多个（具体可参考迁移任务详情中的迁移状态）
 (Optional)
 * param sorts: 排序属性：
- createdTime：按创建时间排序(asc表示按时间正序，desc表示按时间倒序)
 (Optional)
 */
func NewDescribeMigrationsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
    sorts []common.Sort,
) *DescribeMigrationsRequest {

    return &DescribeMigrationsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
        Sorts: sorts,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMigrationsRequestWithoutParam() *DescribeMigrationsRequest {

    return &DescribeMigrationsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1)(Required) */
func (r *DescribeMigrationsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码：取值范围[1,∞)，默认为1(Optional) */
func (r *DescribeMigrationsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小：取值范围[10,100]，默认为10(Optional) */
func (r *DescribeMigrationsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 过滤条件：
- instanceIds：迁移任务ID，精确匹配，可选择多个
- migrationName：迁移任务名称，模糊匹配
- migrationStatuses：迁移任务状态，精确匹配，可选择多个（具体可参考迁移任务详情中的迁移状态）
(Optional) */
func (r *DescribeMigrationsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

/* param sorts: 排序属性：
- createdTime：按创建时间排序(asc表示按时间正序，desc表示按时间倒序)
(Optional) */
func (r *DescribeMigrationsRequest) SetSorts(sorts []common.Sort) {
    r.Sorts = sorts
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMigrationsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMigrationsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMigrationsResult `json:"result"`
}

type DescribeMigrationsResult struct {
    Instances []rdts.ListItem `json:"instances"`
    TotalCount int `json:"totalCount"`
}