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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DescribeDefaultConfigRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 计算节点规格 (Optional) */
    ChNodeClass *string `json:"chNodeClass"`

    /* 分片数 (Optional) */
    ShardNum *int `json:"shardNum"`

    /* 副本数 (Optional) */
    ReplicaNum *int `json:"replicaNum"`
}

/*
 * param regionId: 地域代码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDefaultConfigRequest(
    regionId string,
) *DescribeDefaultConfigRequest {

	return &DescribeDefaultConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:describeDefaultConfig",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param chNodeClass: 计算节点规格 (Optional)
 * param shardNum: 分片数 (Optional)
 * param replicaNum: 副本数 (Optional)
 */
func NewDescribeDefaultConfigRequestWithAllParams(
    regionId string,
    chNodeClass *string,
    shardNum *int,
    replicaNum *int,
) *DescribeDefaultConfigRequest {

    return &DescribeDefaultConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeDefaultConfig",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ChNodeClass: chNodeClass,
        ShardNum: shardNum,
        ReplicaNum: replicaNum,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDefaultConfigRequestWithoutParam() *DescribeDefaultConfigRequest {

    return &DescribeDefaultConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeDefaultConfig",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeDefaultConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param chNodeClass: 计算节点规格(Optional) */
func (r *DescribeDefaultConfigRequest) SetChNodeClass(chNodeClass string) {
    r.ChNodeClass = &chNodeClass
}

/* param shardNum: 分片数(Optional) */
func (r *DescribeDefaultConfigRequest) SetShardNum(shardNum int) {
    r.ShardNum = &shardNum
}

/* param replicaNum: 副本数(Optional) */
func (r *DescribeDefaultConfigRequest) SetReplicaNum(replicaNum int) {
    r.ReplicaNum = &replicaNum
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDefaultConfigRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDefaultConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDefaultConfigResult `json:"result"`
}

type DescribeDefaultConfigResult struct {
    ZkNodeClass string `json:"zkNodeClass"`
    ZkNodeNum int `json:"zkNodeNum"`
    ZkNodeCPU int `json:"zkNodeCPU"`
    ZkNodeMemoryGB int `json:"zkNodeMemoryGB"`
    ZkNodeStorageGB int `json:"zkNodeStorageGB"`
    MonitorNodeClass string `json:"monitorNodeClass"`
    MonitorNodeCPU int `json:"monitorNodeCPU"`
    MonitorNodeMemoryGB int `json:"monitorNodeMemoryGB"`
    MonitorNodeStorageGB int `json:"monitorNodeStorageGB"`
}