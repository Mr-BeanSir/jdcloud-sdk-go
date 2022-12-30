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
	"github.com/Mr-BeanSir/jdcloud-sdk-go/core"
	starshield "github.com/Mr-BeanSir/jdcloud-sdk-go/services/starshield/models"
)

type CreateSSLConfigurationRequest struct {
	core.JDCloudRequest

	/*   */
	Zone_identifier string `json:"zone_identifier"`

	/* 域的SSL证书或证书以及中间层 (Optional) */
	Certificate *string `json:"certificate"`

	/* 域的私钥 (Optional) */
	Private_key *string `json:"private_key"`

	/* SSL泛捆绑在各处有着最高的概率被验证，甚至能被使用过时的或不寻常的信任存储的客户端验证。
	最佳捆绑使用最短的认证链和最新的中间证书。
	而强制捆绑会验证证书链，但不以其他方式修改证书链。
	 (Optional) */
	Bundle_method *string `json:"bundle_method"`

	/*  (Optional) */
	Geo_restrictions *starshield.Geo_restrictions `json:"geo_restrictions"`

	/* “legacy_custom”类型支持在TLS握手中不包含SNI的传统客户端。 (Optional) */
	Ty_pe *string `json:"ty_pe"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSSLConfigurationRequest(
	zone_identifier string,
) *CreateSSLConfigurationRequest {

	return &CreateSSLConfigurationRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/custom_certificates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		Zone_identifier: zone_identifier,
	}
}

/*
  - param zone_identifier:  (Required)
  - param certificate: 域的SSL证书或证书以及中间层 (Optional)
  - param private_key: 域的私钥 (Optional)
  - param bundle_method: SSL泛捆绑在各处有着最高的概率被验证，甚至能被使用过时的或不寻常的信任存储的客户端验证。

最佳捆绑使用最短的认证链和最新的中间证书。
而强制捆绑会验证证书链，但不以其他方式修改证书链。

	(Optional)
	* param geo_restrictions:  (Optional)
	* param ty_pe: “legacy_custom”类型支持在TLS握手中不包含SNI的传统客户端。 (Optional)
*/
func NewCreateSSLConfigurationRequestWithAllParams(
	zone_identifier string,
	certificate *string,
	private_key *string,
	bundle_method *string,
	geo_restrictions *starshield.Geo_restrictions,
	ty_pe *string,
) *CreateSSLConfigurationRequest {

	return &CreateSSLConfigurationRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/custom_certificates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		Zone_identifier:  zone_identifier,
		Certificate:      certificate,
		Private_key:      private_key,
		Bundle_method:    bundle_method,
		Geo_restrictions: geo_restrictions,
		Ty_pe:            ty_pe,
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSSLConfigurationRequestWithoutParam() *CreateSSLConfigurationRequest {

	return &CreateSSLConfigurationRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/custom_certificates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/* param zone_identifier: (Required) */
func (r *CreateSSLConfigurationRequest) SetZone_identifier(zone_identifier string) {
	r.Zone_identifier = zone_identifier
}

/* param certificate: 域的SSL证书或证书以及中间层(Optional) */
func (r *CreateSSLConfigurationRequest) SetCertificate(certificate string) {
	r.Certificate = &certificate
}

/* param private_key: 域的私钥(Optional) */
func (r *CreateSSLConfigurationRequest) SetPrivate_key(private_key string) {
	r.Private_key = &private_key
}

/*
	param bundle_method: SSL泛捆绑在各处有着最高的概率被验证，甚至能被使用过时的或不寻常的信任存储的客户端验证。

最佳捆绑使用最短的认证链和最新的中间证书。
而强制捆绑会验证证书链，但不以其他方式修改证书链。
(Optional)
*/
func (r *CreateSSLConfigurationRequest) SetBundle_method(bundle_method string) {
	r.Bundle_method = &bundle_method
}

/* param geo_restrictions: (Optional) */
func (r *CreateSSLConfigurationRequest) SetGeo_restrictions(geo_restrictions *starshield.Geo_restrictions) {
	r.Geo_restrictions = geo_restrictions
}

/* param ty_pe: “legacy_custom”类型支持在TLS握手中不包含SNI的传统客户端。(Optional) */
func (r *CreateSSLConfigurationRequest) SetTy_pe(ty_pe string) {
	r.Ty_pe = &ty_pe
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSSLConfigurationRequest) GetRegionId() string {
	return ""
}

type CreateSSLConfigurationResponse struct {
	RequestID string                       `json:"requestId"`
	Error     core.ErrorResponse           `json:"error"`
	Result    CreateSSLConfigurationResult `json:"result"`
}

type CreateSSLConfigurationResult struct {
	Data starshield.CustomSSL `json:"data"`
}
