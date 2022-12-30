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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type ElasticIpSpec struct {

    /* 弹性公网IP的限速（单位：Mbps），取值范围为[1-200]  */
    Bandwidth int `json:"bandwidth"`

    /* IP服务商，取值为bgp或no_bgp <br>【cn-north-1】：bgp； <br>【cn-south-1】：bgp； <br>【cn-east-1】：bgp； <br>【cn-east-2】：bgp  */
    Provider string `json:"provider"`

    /* 计费配置 (Optional) */
    ChargeSpec charge.ChargeSpec `json:"chargeSpec"`
}
