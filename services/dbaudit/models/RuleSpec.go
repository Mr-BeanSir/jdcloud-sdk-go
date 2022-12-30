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


type RuleSpec struct {

    /* 规则名称，长度限制32字节 (Optional) */
    RuleName string `json:"ruleName"`

    /* 风险级别: 0->无风险，1->低风险，2->中风险，3->高风险 (Optional) */
    RiskLevel int `json:"riskLevel"`

    /* 规则定义-正则表达式 (Optional) */
    RuleContent string `json:"ruleContent"`

    /* 规则描述，长度限制128字节 (Optional) */
    RuleDesc string `json:"ruleDesc"`
}
