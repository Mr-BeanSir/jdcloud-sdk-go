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


type RemoteQuotaReq struct {

    /*  (Optional) */
    Id int64 `json:"id"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 业务线 (Optional) */
    AppCode string `json:"appCode"`

    /* 资源类型 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 一级资源serviceCode (Optional) */
    ParentCode string `json:"parentCode"`

    /* 所属地域 (Optional) */
    Region string `json:"region"`

    /* 资源id,私有网络 (Optional) */
    ResourceId string `json:"resourceId"`

    /* 修改后配额 (Optional) */
    UserQuota int `json:"userQuota"`

    /* 修改原因 (Optional) */
    Reason string `json:"reason"`

    /* 过期时间 (Optional) */
    ExpireTime string `json:"expireTime"`

    /* 修改配额时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 扩展字段，json格式字符串 (Optional) */
    ExtraInfo string `json:"extraInfo"`
}
