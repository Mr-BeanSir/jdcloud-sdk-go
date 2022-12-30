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


type AntispamResult struct {

    /* 数据唯一标识，能够根据该值定位到该条数据 (Optional) */
    DataId string `json:"dataId"`

    /* 本次请求数据标识，可以根据该标识查询数据最新结果 (Optional) */
    TaskId string `json:"taskId"`

    /* 文本检测状态码，定义为：0：检测成功，1：检测失败 (Optional) */
    Status int `json:"status"`

    /* 审核模式，0：纯机审，1：机审+部分人审，2：机审+全量人审 (Optional) */
    CensonrType int `json:"censonrType"`

    /* 检测结果，0：通过，1：嫌疑，2：不通过 (Optional) */
    Action int `json:"action"`

    /* 语种代码数组，详见 语种代码表 (Optional) */
    Lang []string `json:"lang"`

    /* 是否关联检测命中，true：关联检测命中，flase：原文命中 (Optional) */
    IsRelatedHit bool `json:"isRelatedHit"`

    /* 分类信息 (Optional) */
    Labels []LabelItem `json:"labels"`
}
