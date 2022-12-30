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


type DeviceGroupInfo struct {

    /* 组ID (Optional) */
    DeviceGroupId string `json:"deviceGroupId"`

    /* 设备ID (Optional) */
    DeviceId string `json:"deviceId"`

    /* 设备名称 (Optional) */
    Name string `json:"name"`

    /* 父级设备Id (Optional) */
    ParentId string `json:"parentId"`

    /* 设备状态，0-未激活，1-激活离线，2-激活在线 (Optional) */
    Status int `json:"status"`

    /* 设备标识符 (Optional) */
    Identifier string `json:"identifier"`

    /* 设备描述 (Optional) */
    Description string `json:"description"`

    /* 激活时间 (Optional) */
    ActiveTime int64 `json:"activeTime"`

    /* 最后连接时间 (Optional) */
    LastConnectTime int64 `json:"lastConnectTime"`

    /* 注册时间 (Optional) */
    CreateTime int64 `json:"createTime"`

    /* 修改时间 (Optional) */
    UpdateTime int64 `json:"updateTime"`

    /* 产品名称 (Optional) */
    ProductName string `json:"productName"`

    /* 设备型号 (Optional) */
    Model string `json:"model"`

    /* 设备厂商 (Optional) */
    Manufacturer string `json:"manufacturer"`

    /* 设备采集器类型 (Optional) */
    DeviceCollectorType string `json:"deviceCollectorType"`

    /* 最后离线时间 (Optional) */
    LastDisconnectTime int64 `json:"lastDisconnectTime"`

    /* 订单号 (Optional) */
    OrderId int64 `json:"orderId"`

    /* 设备组名称 (Optional) */
    GroupName string `json:"groupName"`

    /* 设备组标签 (Optional) */
    GroupTag string `json:"groupTag"`
}
