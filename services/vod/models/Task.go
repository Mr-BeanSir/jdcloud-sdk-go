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


type Task struct {

    /* 任务ID (Optional) */
    TaskId int64 `json:"taskId"`

    /* 视频名称 (Optional) */
    Name string `json:"name"`

    /* 分类ID (Optional) */
    CategoryId int `json:"categoryId"`

    /* 分类名称 (Optional) */
    Category string `json:"category"`

    /* 格式 (Optional) */
    Format string `json:"format"`

    /* 文件大小 (Optional) */
    Size int `json:"size"`

    /* 上传状态 (Optional) */
    Status int `json:"status"`
}
