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

type ModifyInstanceSpecRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* data节点规格 (Optional) */
    NodeClass *string `json:"nodeClass"`

    /* data节点磁盘 (Optional) */
    NodeDiskGB *int `json:"nodeDiskGB"`

    /* data节点数 (Optional) */
    NodeCount *int `json:"nodeCount"`

    /* master节点规格 (Optional) */
    MasterClass *string `json:"masterClass"`

    /* coordinating节点规格 (Optional) */
    CoordinatingClass *string `json:"coordinatingClass"`

    /* coordinating节点数 (Optional) */
    CoordinatingCount *int `json:"coordinatingCount"`
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceSpecRequest(
    regionId string,
    instanceId string,
) *ModifyInstanceSpecRequest {

	return &ModifyInstanceSpecRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceSpec",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: regionId (Required)
 * param instanceId: 实例ID (Required)
 * param nodeClass: data节点规格 (Optional)
 * param nodeDiskGB: data节点磁盘 (Optional)
 * param nodeCount: data节点数 (Optional)
 * param masterClass: master节点规格 (Optional)
 * param coordinatingClass: coordinating节点规格 (Optional)
 * param coordinatingCount: coordinating节点数 (Optional)
 */
func NewModifyInstanceSpecRequestWithAllParams(
    regionId string,
    instanceId string,
    nodeClass *string,
    nodeDiskGB *int,
    nodeCount *int,
    masterClass *string,
    coordinatingClass *string,
    coordinatingCount *int,
) *ModifyInstanceSpecRequest {

    return &ModifyInstanceSpecRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceSpec",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        NodeClass: nodeClass,
        NodeDiskGB: nodeDiskGB,
        NodeCount: nodeCount,
        MasterClass: masterClass,
        CoordinatingClass: coordinatingClass,
        CoordinatingCount: coordinatingCount,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceSpecRequestWithoutParam() *ModifyInstanceSpecRequest {

    return &ModifyInstanceSpecRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceSpec",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *ModifyInstanceSpecRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *ModifyInstanceSpecRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param nodeClass: data节点规格(Optional) */
func (r *ModifyInstanceSpecRequest) SetNodeClass(nodeClass string) {
    r.NodeClass = &nodeClass
}

/* param nodeDiskGB: data节点磁盘(Optional) */
func (r *ModifyInstanceSpecRequest) SetNodeDiskGB(nodeDiskGB int) {
    r.NodeDiskGB = &nodeDiskGB
}

/* param nodeCount: data节点数(Optional) */
func (r *ModifyInstanceSpecRequest) SetNodeCount(nodeCount int) {
    r.NodeCount = &nodeCount
}

/* param masterClass: master节点规格(Optional) */
func (r *ModifyInstanceSpecRequest) SetMasterClass(masterClass string) {
    r.MasterClass = &masterClass
}

/* param coordinatingClass: coordinating节点规格(Optional) */
func (r *ModifyInstanceSpecRequest) SetCoordinatingClass(coordinatingClass string) {
    r.CoordinatingClass = &coordinatingClass
}

/* param coordinatingCount: coordinating节点数(Optional) */
func (r *ModifyInstanceSpecRequest) SetCoordinatingCount(coordinatingCount int) {
    r.CoordinatingCount = &coordinatingCount
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceSpecRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceSpecResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceSpecResult `json:"result"`
}

type ModifyInstanceSpecResult struct {
    OrderNum string `json:"orderNum"`
    InstanceId string `json:"instanceId"`
}