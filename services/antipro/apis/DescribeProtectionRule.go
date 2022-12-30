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
    antipro "github.com/jdcloud-api/jdcloud-sdk-go/services/antipro/models"
)

type DescribeProtectionRuleRequest struct {

    core.JDCloudRequest

    /* 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海  */
    RegionId string `json:"regionId"`

    /* 防护包实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 被防护 IP, 缺省时获取防护包实例的防护规则 (Optional) */
    Ip *string `json:"ip"`
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProtectionRuleRequest(
    regionId string,
    instanceId string,
) *DescribeProtectionRuleRequest {

	return &DescribeProtectionRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:describeProtectionRule",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param ip: 被防护 IP, 缺省时获取防护包实例的防护规则 (Optional)
 */
func NewDescribeProtectionRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    ip *string,
) *DescribeProtectionRuleRequest {

    return &DescribeProtectionRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:describeProtectionRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Ip: ip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProtectionRuleRequestWithoutParam() *DescribeProtectionRuleRequest {

    return &DescribeProtectionRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:describeProtectionRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海(Required) */
func (r *DescribeProtectionRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 防护包实例 Id(Required) */
func (r *DescribeProtectionRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param ip: 被防护 IP, 缺省时获取防护包实例的防护规则(Optional) */
func (r *DescribeProtectionRuleRequest) SetIp(ip string) {
    r.Ip = &ip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProtectionRuleRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeProtectionRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProtectionRuleResult `json:"result"`
}

type DescribeProtectionRuleResult struct {
    Data antipro.ProtectionRule `json:"data"`
}