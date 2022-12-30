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


type DomainBandTrafficDataItem struct {

    /* 开始时间戳 (Optional) */
    StartTimeStamp string `json:"startTimeStamp"`

    /* 开始时间戳 (Optional) */
    EndTimeStamp string `json:"endTimeStamp"`

    /* 带宽单位Mbps (Optional) */
    Avgbandwidth float64 `json:"avgbandwidth"`

    /* 流量单位MB (Optional) */
    Flow float64 `json:"flow"`

    /* 请求量 (Optional) */
    Pv int64 `json:"pv"`
}
