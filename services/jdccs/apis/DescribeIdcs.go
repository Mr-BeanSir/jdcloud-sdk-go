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
    jdccs "github.com/jdcloud-api/jdcloud-sdk-go/services/jdccs/models"
)

type DescribeIdcsRequest struct {

    core.JDCloudRequest

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 是否查询全部，默认分页 (Optional) */
    All *int `json:"all"`

    /* 是否强制包含外部机房 yes/no (Optional) */
    IncludeExternalIdc *string `json:"includeExternalIdc"`

}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIdcsRequest(
) *DescribeIdcsRequest {

	return &DescribeIdcsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小，默认为20 (Optional)
 * param all: 是否查询全部，默认分页 (Optional)
 * param includeExternalIdc: 是否强制包含外部机房 yes/no (Optional)
 */
func NewDescribeIdcsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    all *int,
    includeExternalIdc *string,
) *DescribeIdcsRequest {

    return &DescribeIdcsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        All: all,
        IncludeExternalIdc: includeExternalIdc,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIdcsRequestWithoutParam() *DescribeIdcsRequest {

    return &DescribeIdcsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeIdcsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20(Optional) */
func (r *DescribeIdcsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param all: 是否查询全部，默认分页(Optional) */
func (r *DescribeIdcsRequest) SetAll(all int) {
    r.All = &all
}

/* param includeExternalIdc: 是否强制包含外部机房 yes/no(Optional) */
func (r *DescribeIdcsRequest) SetIncludeExternalIdc(includeExternalIdc string) {
    r.IncludeExternalIdc = &includeExternalIdc
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIdcsRequest) GetRegionId() string {
    return ""
}

type DescribeIdcsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIdcsResult `json:"result"`
}

type DescribeIdcsResult struct {
    Idcs []jdccs.Idc `json:"idcs"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}