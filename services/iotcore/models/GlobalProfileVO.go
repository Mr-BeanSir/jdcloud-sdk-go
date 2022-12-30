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


type GlobalProfileVO struct {

    /* 档案数据类型，1=布尔，2=整型，3=浮点型，4=字符串 (Optional) */
    DataType *int `json:"dataType"`

    /* 是否可修改，1=可修改，0=不可修改 (Optional) */
    Editable *int `json:"editable"`

    /* 是否必填,1=必填，0非必填 (Optional) */
    Mandatory *int `json:"mandatory"`

    /* 档案code (Optional) */
    ProfileCode *string `json:"profileCode"`

    /* 当前特征描述 (Optional) */
    ProfileDesc *string `json:"profileDesc"`

    /* 档案name (Optional) */
    ProfileName *string `json:"profileName"`

    /* 档案值 (Optional) */
    ProfileValue *string `json:"profileValue"`

    /* 档案类型，1=全局设备，2=全局产品，3=产品设备 (Optional) */
    Scope *int `json:"scope"`
}
