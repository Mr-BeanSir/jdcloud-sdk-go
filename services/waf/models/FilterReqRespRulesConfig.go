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


type FilterReqRespRulesConfig struct {

    /* id为0时为新增，否则为更新 (Optional) */
    Id int `json:"id"`

    /* 规则名称 (Optional) */
    Name string `json:"name"`

    /* 操作对象，["req"/"resp"] (Optional) */
    OpObject string `json:"opObject"`

    /* 匹配类型，["cookie"/"header"] (Optional) */
    MatchType string `json:"matchType"`

    /* 匹配模式，["del"/"set"] (Optional) */
    MatchOp string `json:"matchOp"`

    /* filter key (Optional) */
    Key string `json:"key"`

    /* 匹配值，当匹配模式为"del"时val应为空 (Optional) */
    Val string `json:"val"`

    /* 匹配值类型，"str"(字符型)，"php"(脚本类型) (Optional) */
    ValType string `json:"valType"`

    /* 更新时间，s (Optional) */
    UpdateTime int `json:"updateTime"`

    /* 0-使用中 1-禁用 (Optional) */
    Disable int `json:"disable"`

    /* 0-用户添加 1-运营后台强制添加 (Optional) */
    IsForceApply int `json:"isForceApply"`
}
