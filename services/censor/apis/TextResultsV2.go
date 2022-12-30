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
    censor "github.com/jdcloud-api/jdcloud-sdk-go/services/censor/models"
)

type TextResultsV2Request struct {

    core.JDCloudRequest

    /* 业务bizType，请联系客户经理获取 (Optional) */
    BizType *string `json:"bizType"`

    /* 接口版本v4 (Optional) */
    Version *string `json:"version"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewTextResultsV2Request(
) *TextResultsV2Request {

	return &TextResultsV2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/text:resultsv2",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param bizType: 业务bizType，请联系客户经理获取 (Optional)
 * param version: 接口版本v4 (Optional)
 */
func NewTextResultsV2RequestWithAllParams(
    bizType *string,
    version *string,
) *TextResultsV2Request {

    return &TextResultsV2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/text:resultsv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        BizType: bizType,
        Version: version,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewTextResultsV2RequestWithoutParam() *TextResultsV2Request {

    return &TextResultsV2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/text:resultsv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param bizType: 业务bizType，请联系客户经理获取(Optional) */
func (r *TextResultsV2Request) SetBizType(bizType string) {
    r.BizType = &bizType
}

/* param version: 接口版本v4(Optional) */
func (r *TextResultsV2Request) SetVersion(version string) {
    r.Version = &version
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r TextResultsV2Request) GetRegionId() string {
    return ""
}

type TextResultsV2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result TextResultsV2Result `json:"result"`
}

type TextResultsV2Result struct {
    Result []censor.TextResultDetailV2 `json:"result"`
}