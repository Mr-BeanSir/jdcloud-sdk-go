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


type SegmentsItem struct {

    /* 断句开始时间，单位秒 (Optional) */
    StartTime int `json:"startTime"`

    /* 断句结束时间，单位秒 (Optional) */
    EndTime int `json:"endTime"`

    /* 音频数据所在断句语音识别原文内容，支持返回异常数据所在断句内容或全部原文内容 (Optional) */
    Content string `json:"content"`

    /* 分类信息，100：色情，200：广告，260：广告法，300：暴恐，400：违禁，500：涉政，600：谩骂，900：其他，1100：涉价值观 (Optional) */
    Label int `json:"label"`

    /* 分类级别，0：通过，1：嫌疑，2：不通过 (Optional) */
    Level int `json:"level"`

    /* 线索详细信息 (Optional) */
    HintList []string `json:"hintList"`
}
