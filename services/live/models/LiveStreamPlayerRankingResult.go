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


type LiveStreamPlayerRankingResult struct {

    /* 起始时间点，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Optional) */
    StartTime string `json:"startTime"`

    /* 结束时间点，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Optional) */
    EndTime string `json:"endTime"`

    /* 排行
 (Optional) */
    Ranking int64 `json:"ranking"`

    /* 流名称
 (Optional) */
    StreamName string `json:"streamName"`

    /* 观众数量
 (Optional) */
    PlayerCount int64 `json:"playerCount"`
}
