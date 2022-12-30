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

type UpdateDomainGroupRequest struct {

    core.JDCloudRequest

    /* 域名组id  */
    Id int `json:"id"`

    /* 域名组内域名，包含主域名 (Optional) */
    Domains []string `json:"domains"`

    /* 主域名,开启共享缓存时必传 (Optional) */
    PrimaryDomain *string `json:"primaryDomain"`

    /* 是否共享内存，共享缓存仅对中国境内加速域名生效 (Optional) */
    ShareCache *string `json:"shareCache"`

    /*  (Optional) */
    DomainGroupName *string `json:"domainGroupName"`
}

/*
 * param id: 域名组id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateDomainGroupRequest(
    id int,
) *UpdateDomainGroupRequest {

	return &UpdateDomainGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domainGroup/{id}:update",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Id: id,
	}
}

/*
 * param id: 域名组id (Required)
 * param domains: 域名组内域名，包含主域名 (Optional)
 * param primaryDomain: 主域名,开启共享缓存时必传 (Optional)
 * param shareCache: 是否共享内存，共享缓存仅对中国境内加速域名生效 (Optional)
 * param domainGroupName:  (Optional)
 */
func NewUpdateDomainGroupRequestWithAllParams(
    id int,
    domains []string,
    primaryDomain *string,
    shareCache *string,
    domainGroupName *string,
) *UpdateDomainGroupRequest {

    return &UpdateDomainGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup/{id}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Id: id,
        Domains: domains,
        PrimaryDomain: primaryDomain,
        ShareCache: shareCache,
        DomainGroupName: domainGroupName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateDomainGroupRequestWithoutParam() *UpdateDomainGroupRequest {

    return &UpdateDomainGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup/{id}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param id: 域名组id(Required) */
func (r *UpdateDomainGroupRequest) SetId(id int) {
    r.Id = id
}
/* param domains: 域名组内域名，包含主域名(Optional) */
func (r *UpdateDomainGroupRequest) SetDomains(domains []string) {
    r.Domains = domains
}
/* param primaryDomain: 主域名,开启共享缓存时必传(Optional) */
func (r *UpdateDomainGroupRequest) SetPrimaryDomain(primaryDomain string) {
    r.PrimaryDomain = &primaryDomain
}
/* param shareCache: 是否共享内存，共享缓存仅对中国境内加速域名生效(Optional) */
func (r *UpdateDomainGroupRequest) SetShareCache(shareCache string) {
    r.ShareCache = &shareCache
}
/* param domainGroupName: (Optional) */
func (r *UpdateDomainGroupRequest) SetDomainGroupName(domainGroupName string) {
    r.DomainGroupName = &domainGroupName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateDomainGroupRequest) GetRegionId() string {
    return ""
}

type UpdateDomainGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateDomainGroupResult `json:"result"`
}

type UpdateDomainGroupResult struct {
}