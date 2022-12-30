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


type PreOrders struct {

    /* 预购订单编号 (Optional) */
    PreOrderId string `json:"preOrderId"`

    /* 预购订单数量 (Optional) */
    PreOrderNum string `json:"preOrderNum"`

    /* 产品Key (Optional) */
    ProductKey string `json:"productKey"`

    /* 产品名称 (Optional) */
    ProductName string `json:"productName"`

    /* 用户Pin (Optional) */
    UserPin string `json:"userPin"`

    /* 用户名称 (Optional) */
    UserName string `json:"userName"`

    /* 已录入数量 (Optional) */
    RecordNum int `json:"recordNum"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 订单状态，0-下单,1-处理完成 (Optional) */
    PreOrderStatus string `json:"preOrderStatus"`
}
