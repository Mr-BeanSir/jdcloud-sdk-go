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
    clickhouse "github.com/jdcloud-api/jdcloud-sdk-go/services/clickhouse/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeSlowlogResultRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 节点名称 (Optional) */
    SegmentName *string `json:"segmentName"`

    /* 查询时间 (Optional) */
    QueryTime *string `json:"queryTime"`

    /*  (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSlowlogResultRequest(
    regionId string,
    instanceId string,
) *DescribeSlowlogResultRequest {

	return &DescribeSlowlogResultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/slowlog:describeSlowlogResult",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口 (Optional)
 * param startTime: 开始时间 (Optional)
 * param endTime: 结束时间 (Optional)
 * param segmentName: 节点名称 (Optional)
 * param queryTime: 查询时间 (Optional)
 * param filters:  (Optional)
 */
func NewDescribeSlowlogResultRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
    startTime *string,
    endTime *string,
    segmentName *string,
    queryTime *string,
    filters []common.Filter,
) *DescribeSlowlogResultRequest {

    return &DescribeSlowlogResultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/slowlog:describeSlowlogResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        StartTime: startTime,
        EndTime: endTime,
        SegmentName: segmentName,
        QueryTime: queryTime,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSlowlogResultRequestWithoutParam() *DescribeSlowlogResultRequest {

    return &DescribeSlowlogResultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/slowlog:describeSlowlogResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeSlowlogResultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *DescribeSlowlogResultRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页;(Optional) */
func (r *DescribeSlowlogResultRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口(Optional) */
func (r *DescribeSlowlogResultRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param startTime: 开始时间(Optional) */
func (r *DescribeSlowlogResultRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间(Optional) */
func (r *DescribeSlowlogResultRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param segmentName: 节点名称(Optional) */
func (r *DescribeSlowlogResultRequest) SetSegmentName(segmentName string) {
    r.SegmentName = &segmentName
}

/* param queryTime: 查询时间(Optional) */
func (r *DescribeSlowlogResultRequest) SetQueryTime(queryTime string) {
    r.QueryTime = &queryTime
}

/* param filters: (Optional) */
func (r *DescribeSlowlogResultRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSlowlogResultRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSlowlogResultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSlowlogResultResult `json:"result"`
}

type DescribeSlowlogResultResult struct {
    SlowlogResult clickhouse.SlowLog `json:"slowlogResult"`
    TotalCount int `json:"totalCount"`
}