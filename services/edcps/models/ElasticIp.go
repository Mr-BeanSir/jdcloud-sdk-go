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

package models

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type ElasticIp struct {

    /* 地域代码, 如cn-east-tz1 (Optional) */
    Region string `json:"region"`

    /* 弹性公网IPID (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 弹性公网IP (Optional) */
    ElasticIp string `json:"elasticIp"`

    /* 带宽, 单位Mbps (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 额外上行带宽, 单位Mbps (Optional) */
    ExtraUplinkBandwidth int `json:"extraUplinkBandwidth"`

    /* 链路类型 (Optional) */
    LineType string `json:"lineType"`

    /* 状态 (Optional) */
    Status string `json:"status"`

    /* 实例类型 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 绑定的ip地址 (Optional) */
    TargetIp string `json:"targetIp"`

    /* 共享带宽 id (Optional) */
    BandwidthPackageId string `json:"bandwidthPackageId"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`
}
