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

type ModifyReplicationRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 复制任务ID  */
    TaskId string `json:"taskId"`

    /* 目标实例备注说明 (Optional) */
    TargetComment *string `json:"targetComment"`

    /* 目标类型为TiDB或MySQL时，连接目标实例的用户名 (Optional) */
    TargetUser *string `json:"targetUser"`

    /* 目标类型为TiDB或MySQL时，连接目标实例的密码 (Optional) */
    TargetPassword *string `json:"targetPassword"`

    /* Kafka的Topic (Optional) */
    KafkaTopic *string `json:"kafkaTopic"`

    /* Kafka的版本 (Optional) */
    KafkaVersion *string `json:"kafkaVersion"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param taskId: 复制任务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyReplicationRequest(
    regionId string,
    instanceId string,
    taskId string,
) *ModifyReplicationRequest {

	return &ModifyReplicationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/replications/{taskId}:modifyReplication",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        TaskId: taskId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param taskId: 复制任务ID (Required)
 * param targetComment: 目标实例备注说明 (Optional)
 * param targetUser: 目标类型为TiDB或MySQL时，连接目标实例的用户名 (Optional)
 * param targetPassword: 目标类型为TiDB或MySQL时，连接目标实例的密码 (Optional)
 * param kafkaTopic: Kafka的Topic (Optional)
 * param kafkaVersion: Kafka的版本 (Optional)
 */
func NewModifyReplicationRequestWithAllParams(
    regionId string,
    instanceId string,
    taskId string,
    targetComment *string,
    targetUser *string,
    targetPassword *string,
    kafkaTopic *string,
    kafkaVersion *string,
) *ModifyReplicationRequest {

    return &ModifyReplicationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/replications/{taskId}:modifyReplication",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        TaskId: taskId,
        TargetComment: targetComment,
        TargetUser: targetUser,
        TargetPassword: targetPassword,
        KafkaTopic: kafkaTopic,
        KafkaVersion: kafkaVersion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyReplicationRequestWithoutParam() *ModifyReplicationRequest {

    return &ModifyReplicationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/replications/{taskId}:modifyReplication",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *ModifyReplicationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 实例ID(Required) */
func (r *ModifyReplicationRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param taskId: 复制任务ID(Required) */
func (r *ModifyReplicationRequest) SetTaskId(taskId string) {
    r.TaskId = taskId
}
/* param targetComment: 目标实例备注说明(Optional) */
func (r *ModifyReplicationRequest) SetTargetComment(targetComment string) {
    r.TargetComment = &targetComment
}
/* param targetUser: 目标类型为TiDB或MySQL时，连接目标实例的用户名(Optional) */
func (r *ModifyReplicationRequest) SetTargetUser(targetUser string) {
    r.TargetUser = &targetUser
}
/* param targetPassword: 目标类型为TiDB或MySQL时，连接目标实例的密码(Optional) */
func (r *ModifyReplicationRequest) SetTargetPassword(targetPassword string) {
    r.TargetPassword = &targetPassword
}
/* param kafkaTopic: Kafka的Topic(Optional) */
func (r *ModifyReplicationRequest) SetKafkaTopic(kafkaTopic string) {
    r.KafkaTopic = &kafkaTopic
}
/* param kafkaVersion: Kafka的版本(Optional) */
func (r *ModifyReplicationRequest) SetKafkaVersion(kafkaVersion string) {
    r.KafkaVersion = &kafkaVersion
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyReplicationRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyReplicationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyReplicationResult `json:"result"`
}

type ModifyReplicationResult struct {
}