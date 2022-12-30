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

type ShareImageRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 镜像ID。  */
    ImageId string `json:"imageId"`

    /* 共享的目标京东云帐户列表。  */
    Pins []string `json:"pins"`

}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 * param pins: 共享的目标京东云帐户列表。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewShareImageRequest(
    regionId string,
    imageId string,
    pins []string,
) *ShareImageRequest {

	return &ShareImageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images/{imageId}:share",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ImageId: imageId,
        Pins: pins,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 * param pins: 共享的目标京东云帐户列表。 (Required)
 */
func NewShareImageRequestWithAllParams(
    regionId string,
    imageId string,
    pins []string,
) *ShareImageRequest {

    return &ShareImageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}:share",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageId: imageId,
        Pins: pins,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewShareImageRequestWithoutParam() *ShareImageRequest {

    return &ShareImageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}:share",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *ShareImageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageId: 镜像ID。(Required) */
func (r *ShareImageRequest) SetImageId(imageId string) {
    r.ImageId = imageId
}

/* param pins: 共享的目标京东云帐户列表。(Required) */
func (r *ShareImageRequest) SetPins(pins []string) {
    r.Pins = pins
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ShareImageRequest) GetRegionId() string {
    return r.RegionId
}

type ShareImageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ShareImageResult `json:"result"`
}

type ShareImageResult struct {
}