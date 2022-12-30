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


type CardInfoResp struct {

    /* 物联网卡iccid (Optional) */
    Iccid string `json:"iccid"`

    /* 物联网卡套餐名称 (Optional) */
    PackageName string `json:"packageName"`

    /* 物联网卡激活时间 (Optional) */
    ActiveTm string `json:"activeTm"`

    /* 套餐到期时间 (Optional) */
    PackageExpiredTm string `json:"packageExpiredTm"`
}
