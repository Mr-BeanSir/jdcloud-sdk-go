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


type StatusListCfg struct {

    /* 规则名称 (Optional) */
    Name string `json:"name"`

    /* 原有状态码 (Optional) */
    OriStatus string `json:"oriStatus"`

    /* 设置状态码，只能为"200"，"302" (Optional) */
    SetStatus string `json:"setStatus"`

    /* 设置状态码为"200"时，改值为自定义页面名称，"302"时为跳转url (Optional) */
    Val string `json:"val"`

    /* 更新时间，s (Optional) */
    UpdateTime int `json:"updateTime"`

    /* 0-使用中 1-禁用 (Optional) */
    Disable int `json:"disable"`
}
