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


type VideoSourceInfo struct {

    /* 视频ID (Optional) */
    VideoId string `json:"videoId"`

    /* 文件名称 (Optional) */
    FileName string `json:"fileName"`

    /* 文件URL（对象存储URL） (Optional) */
    FileUrl string `json:"fileUrl"`

    /* 文件内容摘要。
对于分片上传的视频，若想获取其MD5，可以在上传时，指定Content-MD5。
 (Optional) */
    Md5 string `json:"md5"`
}
