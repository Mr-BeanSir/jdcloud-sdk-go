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


type CreateParserSpec struct {

    /*   */
    ParserFields []ParserField `json:"parserFields"`

    /* 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp  */
    ParserMode string `json:"parserMode"`

    /* 解析语法 (Optional) */
    ParserPattern string `json:"parserPattern"`

    /* 日志样例 (Optional) */
    ParserSample string `json:"parserSample"`

    /* 预处理任务列表。按照数组的顺序执行。 (Optional) */
    Pipelines []PipelineSpec `json:"pipelines"`
}
