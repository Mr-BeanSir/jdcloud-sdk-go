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

type DeleteLiveStreamTranslateRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 流名称  */
    StreamName string `json:"streamName"`

    /* 翻译模板  */
    Template string `json:"template"`
}

/*
 * param publishDomain: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param template: 翻译模板 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveStreamTranslateRequest(
    publishDomain string,
    appName string,
    streamName string,
    template string,
) *DeleteLiveStreamTranslateRequest {

	return &DeleteLiveStreamTranslateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/translateStream/{publishDomain}/appNames/{appName}/streams/{streamName}/templates/{template}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        Template: template,
	}
}

/*
 * param publishDomain: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param template: 翻译模板 (Required)
 */
func NewDeleteLiveStreamTranslateRequestWithAllParams(
    publishDomain string,
    appName string,
    streamName string,
    template string,
) *DeleteLiveStreamTranslateRequest {

    return &DeleteLiveStreamTranslateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/translateStream/{publishDomain}/appNames/{appName}/streams/{streamName}/templates/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveStreamTranslateRequestWithoutParam() *DeleteLiveStreamTranslateRequest {

    return &DeleteLiveStreamTranslateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/translateStream/{publishDomain}/appNames/{appName}/streams/{streamName}/templates/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Required) */
func (r *DeleteLiveStreamTranslateRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 应用名称(Required) */
func (r *DeleteLiveStreamTranslateRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param streamName: 流名称(Required) */
func (r *DeleteLiveStreamTranslateRequest) SetStreamName(streamName string) {
    r.StreamName = streamName
}

/* param template: 翻译模板(Required) */
func (r *DeleteLiveStreamTranslateRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveStreamTranslateRequest) GetRegionId() string {
    return ""
}

type DeleteLiveStreamTranslateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveStreamTranslateResult `json:"result"`
}

type DeleteLiveStreamTranslateResult struct {
}