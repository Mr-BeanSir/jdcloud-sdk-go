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


type RecordFile struct {

    /* 录制文件ID
 (Optional) */
    FileId string `json:"fileId"`

    /* 文件格式
 (Optional) */
    Format string `json:"format"`

    /* 视频宽度
- 单位: 像素
 (Optional) */
    Width int `json:"width"`

    /* 视频高度
- 单位: 像素
 (Optional) */
    Height int `json:"height"`

    /* 录制开始时间
 (Optional) */
    StartTime string `json:"startTime"`

    /* 录制结束时间
 (Optional) */
    EndTime string `json:"endTime"`

    /* 视频时长，单位：毫秒
 (Optional) */
    Duration int `json:"duration"`

    /* 文件大小，单位：B
 (Optional) */
    Size int `json:"size"`

    /* 码率
 (Optional) */
    Bitrate int `json:"bitrate"`

    /* 帧率
 (Optional) */
    Fps int `json:"fps"`

    /* 文件地址
 (Optional) */
    FileUrl string `json:"fileUrl"`

    /* 创建时间
 (Optional) */
    CreateTime string `json:"createTime"`
}
