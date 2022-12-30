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

type CreateAdminProductRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* IoT Engine实例ID信息  */
    InstanceId string `json:"instanceId"`

    /* 产品名称，名称不可为空，3-30个字符，只支持汉字、英文字母、数字、下划线“_”及中划线“-”，必须以汉字、英文字母及数字开头结尾  */
    ProductName string `json:"productName"`

    /* 节点类型，取值：
0：设备。设备不能挂载子设备。可以直连物联网平台，也可以作为网关的子设备连接物联网平台
1：网关。网关可以挂载子设备，具有子设备管理模块，维持子设备的拓扑关系，和将拓扑关系同步到物联网平台
  */
    ProductType int `json:"productType"`

    /* 产品描述，80字符以内 (Optional) */
    ProductDescription *string `json:"productDescription"`

    /* 物模型模板ID，内部参数，用户不可见，默认为自定义 (Optional) */
    TemplateId *string `json:"templateId"`

    /* 内部标签，内部参数，用户不可见，隐藏标签：hidden:true (Optional) */
    InternalTags *interface{} `json:"internalTags"`

    /* 产品名下所有设备的采集器类型  */
    CollDeviceType string `json:"collDeviceType"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoT Engine实例ID信息 (Required)
 * param productName: 产品名称，名称不可为空，3-30个字符，只支持汉字、英文字母、数字、下划线“_”及中划线“-”，必须以汉字、英文字母及数字开头结尾 (Required)
 * param productType: 节点类型，取值：
0：设备。设备不能挂载子设备。可以直连物联网平台，也可以作为网关的子设备连接物联网平台
1：网关。网关可以挂载子设备，具有子设备管理模块，维持子设备的拓扑关系，和将拓扑关系同步到物联网平台
 (Required)
 * param collDeviceType: 产品名下所有设备的采集器类型 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAdminProductRequest(
    regionId string,
    instanceId string,
    productName string,
    productType int,
    collDeviceType string,
) *CreateAdminProductRequest {

	return &CreateAdminProductRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/productsAdmin",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ProductName: productName,
        ProductType: productType,
        CollDeviceType: collDeviceType,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoT Engine实例ID信息 (Required)
 * param productName: 产品名称，名称不可为空，3-30个字符，只支持汉字、英文字母、数字、下划线“_”及中划线“-”，必须以汉字、英文字母及数字开头结尾 (Required)
 * param productType: 节点类型，取值：
0：设备。设备不能挂载子设备。可以直连物联网平台，也可以作为网关的子设备连接物联网平台
1：网关。网关可以挂载子设备，具有子设备管理模块，维持子设备的拓扑关系，和将拓扑关系同步到物联网平台
 (Required)
 * param productDescription: 产品描述，80字符以内 (Optional)
 * param templateId: 物模型模板ID，内部参数，用户不可见，默认为自定义 (Optional)
 * param internalTags: 内部标签，内部参数，用户不可见，隐藏标签：hidden:true (Optional)
 * param collDeviceType: 产品名下所有设备的采集器类型 (Required)
 */
func NewCreateAdminProductRequestWithAllParams(
    regionId string,
    instanceId string,
    productName string,
    productType int,
    productDescription *string,
    templateId *string,
    internalTags *interface{},
    collDeviceType string,
) *CreateAdminProductRequest {

    return &CreateAdminProductRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/productsAdmin",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ProductName: productName,
        ProductType: productType,
        ProductDescription: productDescription,
        TemplateId: templateId,
        InternalTags: internalTags,
        CollDeviceType: collDeviceType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAdminProductRequestWithoutParam() *CreateAdminProductRequest {

    return &CreateAdminProductRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loongrayinstances/{instanceId}/productsAdmin",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateAdminProductRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: IoT Engine实例ID信息(Required) */
func (r *CreateAdminProductRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param productName: 产品名称，名称不可为空，3-30个字符，只支持汉字、英文字母、数字、下划线“_”及中划线“-”，必须以汉字、英文字母及数字开头结尾(Required) */
func (r *CreateAdminProductRequest) SetProductName(productName string) {
    r.ProductName = productName
}

/* param productType: 节点类型，取值：
0：设备。设备不能挂载子设备。可以直连物联网平台，也可以作为网关的子设备连接物联网平台
1：网关。网关可以挂载子设备，具有子设备管理模块，维持子设备的拓扑关系，和将拓扑关系同步到物联网平台
(Required) */
func (r *CreateAdminProductRequest) SetProductType(productType int) {
    r.ProductType = productType
}

/* param productDescription: 产品描述，80字符以内(Optional) */
func (r *CreateAdminProductRequest) SetProductDescription(productDescription string) {
    r.ProductDescription = &productDescription
}

/* param templateId: 物模型模板ID，内部参数，用户不可见，默认为自定义(Optional) */
func (r *CreateAdminProductRequest) SetTemplateId(templateId string) {
    r.TemplateId = &templateId
}

/* param internalTags: 内部标签，内部参数，用户不可见，隐藏标签：hidden:true(Optional) */
func (r *CreateAdminProductRequest) SetInternalTags(internalTags interface{}) {
    r.InternalTags = &internalTags
}

/* param collDeviceType: 产品名下所有设备的采集器类型(Required) */
func (r *CreateAdminProductRequest) SetCollDeviceType(collDeviceType string) {
    r.CollDeviceType = collDeviceType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAdminProductRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAdminProductResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAdminProductResult `json:"result"`
}

type CreateAdminProductResult struct {
    ProductKey string `json:"productKey"`
}