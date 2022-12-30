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


type AzIdSpec struct {

    /* AZ指定方式，SpecifyByReplicaGroup表示按副本组指定，SpecifyByCluster表示按整个集群指定 (Optional) */
    AzSpecifyType *string `json:"azSpecifyType"`

    /* 为集群指定的AZ范围，按集群指定AZ时生效 (Optional) */
    AzsForCluster []string `json:"azsForCluster"`

    /* 缓存Redis主实例所在的可用区ID，按副本组指定AZ时生效  */
    Master string `json:"master"`

    /* 缓存Redis从实例所在的可用区ID，按副本组指定AZ时生效  */
    Slave string `json:"slave"`
}
