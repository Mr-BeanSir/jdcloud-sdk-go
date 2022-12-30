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

type ImageScanV2Request struct {

    core.JDCloudRequest

    /* 指定检测场景 (Optional) */
    Images []censor.ImageTaskV2 `json:"images"`

    /* 接口版本v4.1 (Optional) */
    Version *string `json:"version"`

    /* 业务bizType，请联系客户经理获取 (Optional) */
    BizType *string `json:"bizType"`

    /* 接口指定过检分类，可多选，过检分类列表：100：色情，110：性感低俗，200：广告，210：二维码，260：广告法，300：暴恐，400：违禁，500：涉政，800：恶心类，900：其他，1100：涉价值观响应结果 (Optional) */
    CheckLabels []string `json:"checkLabels"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewImageScanV2Request(
) *ImageScanV2Request {

	return &ImageScanV2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/image:scanv2",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param images: 指定检测场景 (Optional)
 * param version: 接口版本v4.1 (Optional)
 * param bizType: 业务bizType，请联系客户经理获取 (Optional)
 * param checkLabels: 接口指定过检分类，可多选，过检分类列表：100：色情，110：性感低俗，200：广告，210：二维码，260：广告法，300：暴恐，400：违禁，500：涉政，800：恶心类，900：其他，1100：涉价值观响应结果 (Optional)
 */
func NewImageScanV2RequestWithAllParams(
    images []censor.ImageTaskV2,
    version *string,
    bizType *string,
    checkLabels []string,
) *ImageScanV2Request {

    return &ImageScanV2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/image:scanv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Images: images,
        Version: version,
        BizType: bizType,
        CheckLabels: checkLabels,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewImageScanV2RequestWithoutParam() *ImageScanV2Request {

    return &ImageScanV2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/image:scanv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param images: 指定检测场景(Optional) */
func (r *ImageScanV2Request) SetImages(images []censor.ImageTaskV2) {
    r.Images = images
}

/* param version: 接口版本v4.1(Optional) */
func (r *ImageScanV2Request) SetVersion(version string) {
    r.Version = &version
}

/* param bizType: 业务bizType，请联系客户经理获取(Optional) */
func (r *ImageScanV2Request) SetBizType(bizType string) {
    r.BizType = &bizType
}

/* param checkLabels: 接口指定过检分类，可多选，过检分类列表：100：色情，110：性感低俗，200：广告，210：二维码，260：广告法，300：暴恐，400：违禁，500：涉政，800：恶心类，900：其他，1100：涉价值观响应结果(Optional) */
func (r *ImageScanV2Request) SetCheckLabels(checkLabels []string) {
    r.CheckLabels = checkLabels
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ImageScanV2Request) GetRegionId() string {
    return ""
}

type ImageScanV2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ImageScanV2Result `json:"result"`
}

type ImageScanV2Result struct {
    Antispam []censor.ImageAntispamResult `json:"antispam"`
    Ocr []censor.OCRResult `json:"ocr"`
    Face []censor.FaceResult `json:"face"`
    Quality []censor.QualityResult `json:"quality"`
    Logo []censor.LogoResult `json:"logo"`
    Scene []censor.SceneResult `json:"scene"`
}