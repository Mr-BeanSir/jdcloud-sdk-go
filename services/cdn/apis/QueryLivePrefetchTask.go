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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryLivePrefetchTaskRequest struct {

    core.JDCloudRequest

    /* 预热的URL (Optional) */
    UrlList []string `json:"urlList"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryLivePrefetchTaskRequest(
) *QueryLivePrefetchTaskRequest {

	return &QueryLivePrefetchTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/task:queryLivePrefetchTask",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param urlList: 预热的URL (Optional)
 */
func NewQueryLivePrefetchTaskRequestWithAllParams(
    urlList []string,
) *QueryLivePrefetchTaskRequest {

    return &QueryLivePrefetchTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:queryLivePrefetchTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        UrlList: urlList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryLivePrefetchTaskRequestWithoutParam() *QueryLivePrefetchTaskRequest {

    return &QueryLivePrefetchTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:queryLivePrefetchTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param urlList: 预热的URL(Optional) */
func (r *QueryLivePrefetchTaskRequest) SetUrlList(urlList []string) {
    r.UrlList = urlList
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryLivePrefetchTaskRequest) GetRegionId() string {
    return ""
}

type QueryLivePrefetchTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryLivePrefetchTaskResult `json:"result"`
}

type QueryLivePrefetchTaskResult struct {
    Result []cdn.QueryLivePrefetchItem `json:"result"`
}