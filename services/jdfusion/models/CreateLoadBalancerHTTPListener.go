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


type CreateLoadBalancerHTTPListener struct {

    /* 是  负载均衡实例ID (Optional) */
    LoadBalancerId *string `json:"loadBalancerId"`

    /* 是  负载均衡实例前端使用的端口 (Optional) */
    ListenerPort *int `json:"listenerPort"`

    /* 是  负载均衡实例后端使用的端口 (Optional) */
    BackendServerPort *int `json:"backendServerPort"`

    /* 否  服务器组ID (Optional) */
    VserverGroupId *string `json:"vserverGroupId"`

    /* 是  监听的带宽峰值 (Optional) */
    Bandwidth *int `json:"bandwidth"`

    /* 否  健康检查使用的端口 (Optional) */
    HealthCheckConnectPort *int `json:"healthCheckConnectPort"`

    /* 否  健康检查连续成功多少次后，将后端服务器的健康检查状态由fail判定为success。取值：2-10 (Optional) */
    HealthyThreshold *int `json:"healthyThreshold"`

    /* 否  健康检查连续失败多少次后，将后端服务器的健康检查状态由success判定为fail。取值：2-10 (Optional) */
    UnhealthyThreshold *int `json:"unhealthyThreshold"`

    /* 否  接收来自运行状况检查的响应需要等待的时间。如果后端ECS在指定的时间内没有正确响应，则判定为健康检查失败。取值：1-300（秒） (Optional) */
    HealthCheckTimeout *int `json:"healthCheckTimeout"`

    /* 否  健康检查的时间间隔。取值：1-50（秒） (Optional) */
    HealthCheckInterval *int `json:"healthCheckInterval"`

    /* 否  健康检查正常的HTTP状态码，多个状态码用逗号分隔。取值：http_2xx（默认值） | http_3xx | http_4xx | http_5xx (Optional) */
    HealthCheckHttpCode *string `json:"healthCheckHttpCode"`
}
