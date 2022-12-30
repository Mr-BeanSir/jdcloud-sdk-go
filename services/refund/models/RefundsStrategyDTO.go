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


type RefundsStrategyDTO struct {

    /* 退款策略主键id(修改必传，新增不传) (Optional) */
    Id int `json:"id"`

    /* 策略来源类型 0-运营后台退款策略  1-控制台退款策略 2-openapi退款策略 3-外部服务上云退款策略 4-产品线退款 5-出货失败退款策略 (Optional) */
    StrategyType int `json:"strategyType"`

    /* 退款方式策略 1 原支付方式返回  2 退款至余额 (Optional) */
    RefundChannelSwitch int `json:"refundChannelSwitch"`

    /* 退款类型策略 0-退款退货 1-退款不退货 (Optional) */
    RefundWaySwitch int `json:"refundWaySwitch"`

    /* 退款范围策略  0-仅退现金  1-全部退款（含代金券） (Optional) */
    RefundAreaSwitch int `json:"refundAreaSwitch"`

    /* 发票允许的状态   0-未开票 1-已退票  2-欠票补扣 (Optional) */
    InvoiceStatusSwitch int `json:"invoiceStatusSwitch"`

    /* 审批策略 0-关闭 1-开启 (Optional) */
    ApproveSwitch int `json:"approveSwitch"`

    /* 退款订单默认策略   0-退全部可退订单  1-退最后一笔订单 2-仅退新购订单 3-仅退续费订单 (Optional) */
    RefundOrderSwitch int `json:"refundOrderSwitch"`

    /* 业务产品线(如vm) (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 优先级 0-兜底策略  1-策略类型策略 2-service_code策略 (Optional) */
    Priority int `json:"priority"`

    /* 状态  0-无效 1-生效 (Optional) */
    Status int `json:"status"`

    /* 退款类型 1-订单退款 2-充值单退款 3-线下退款（人工退款） (Optional) */
    RefundType int `json:"refundType"`

    /* 权限列表,逗号分隔. erp,pin (Optional) */
    PermissionList string `json:"permissionList"`
}
