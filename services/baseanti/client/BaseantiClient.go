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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    baseanti "github.com/jdcloud-api/jdcloud-sdk-go/services/baseanti/apis"
    "encoding/json"
    "errors"
)

type BaseantiClient struct {
    core.JDCloudClient
}

func NewBaseantiClient(credential *core.Credential) *BaseantiClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("baseanti.jdcloud-api.com")

    return &BaseantiClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "baseanti",
            Revision:    "1.3.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *BaseantiClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *BaseantiClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *BaseantiClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询基础防护已防护的Web应用防火墙 IP 的安全信息 */
func (c *BaseantiClient) DescribeWafIpResources(request *baseanti.DescribeWafIpResourcesRequest) (*baseanti.DescribeWafIpResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeWafIpResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询基础防护已防护公网 IP 安全信息, 支持 ipv4 和 ipv6 */
func (c *BaseantiClient) DescribeIpSafetyInfo(request *baseanti.DescribeIpSafetyInfoRequest) (*baseanti.DescribeIpSafetyInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeIpSafetyInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 设置基础防护已防护公网 IP 的清洗阈值, 支持 ipv4 和 ipv6 */
func (c *BaseantiClient) SetIpCleanThreshold(request *baseanti.SetIpCleanThresholdRequest) (*baseanti.SetIpCleanThresholdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.SetIpCleanThresholdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询多个公网 IP 的监控流量, 支持 ipv4 和 ipv6 */
func (c *BaseantiClient) DescribeIpMonitorFlow(request *baseanti.DescribeIpMonitorFlowRequest) (*baseanti.DescribeIpMonitorFlowResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeIpMonitorFlowResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询基础防护已防护的公网 IP 的安全信息列表. 包括私有网络的弹性公网 IP(运营商级 NAT 保留地址除外), 云物理服务器的公网 IP 和弹性公网 IP. (已废弃, 建议使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeelasticipresources'>describeElasticIpResources</a>, <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describecpsipresources'>describeCpsIpResources</a> 接口) */
func (c *BaseantiClient) DescribeIpResources(request *baseanti.DescribeIpResourcesRequest) (*baseanti.DescribeIpResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeIpResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 攻击情况统计 */
func (c *BaseantiClient) DescribeAttackStatistics(request *baseanti.DescribeAttackStatisticsRequest) (*baseanti.DescribeAttackStatisticsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeAttackStatisticsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询基础防护已防护的私有网络的弹性公网 IP 的安全信息. 包括私有网络的弹性公网 IP(运营商级 NAT 保留地址除外)
 */
func (c *BaseantiClient) DescribeElasticIpResources(request *baseanti.DescribeElasticIpResourcesRequest) (*baseanti.DescribeElasticIpResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeElasticIpResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询基础防护已防护的托管区 IP 的安全信息 */
func (c *BaseantiClient) DescribeCcsIpResources(request *baseanti.DescribeCcsIpResourcesRequest) (*baseanti.DescribeCcsIpResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeCcsIpResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询公网 IP 可设置清洗阈值范围, 支持 ipv4 和 ipv6 */
func (c *BaseantiClient) DescribeIpCleanThresholdRange(request *baseanti.DescribeIpCleanThresholdRangeRequest) (*baseanti.DescribeIpCleanThresholdRangeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeIpCleanThresholdRangeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询各类型攻击次数 */
func (c *BaseantiClient) DescribeAttackTypeCount(request *baseanti.DescribeAttackTypeCountRequest) (*baseanti.DescribeAttackTypeCountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeAttackTypeCountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询公网 IP 的攻击记录, 仅支持 ipv4. (已废弃, 建议使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeattacklogs'>describeAttackLogs</a> 接口)
 */
func (c *BaseantiClient) DescribeIpResourceProtectInfo(request *baseanti.DescribeIpResourceProtectInfoRequest) (*baseanti.DescribeIpResourceProtectInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeIpResourceProtectInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询基础防护已防护的云物理服务器公网 IP 的安全信息. 包括云物理服务器的公网 IP 和弹性公网 IP.
 */
func (c *BaseantiClient) DescribeCpsIpResources(request *baseanti.DescribeCpsIpResourcesRequest) (*baseanti.DescribeCpsIpResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeCpsIpResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 设置基础防护已防护公网 IP 的清洗阈值, 仅支持 ipv4. (已废弃, 建议使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/setipcleanthreshold'>setIpCleanThreshold</a> 接口)
 */
func (c *BaseantiClient) SetCleanThreshold(request *baseanti.SetCleanThresholdRequest) (*baseanti.SetCleanThresholdResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.SetCleanThresholdResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询公网 IP 的 endTime 之前 15 分钟内监控流量, 仅支持 ipv4. (已废弃, 建议使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeipmonitorflow'>describeIpMonitorFlow</a> 接口)
 */
func (c *BaseantiClient) DescribeIpResourceFlow(request *baseanti.DescribeIpResourceFlowRequest) (*baseanti.DescribeIpResourceFlowResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeIpResourceFlowResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询公网 IP 安全信息, 仅支持 ipv4. (已废弃, 建议使用 <a href='http://docs.jdcloud.com/anti-ddos-basic/api/describeipsafetyinfo'>describeIpSafetyInfo</a> 接口)
 */
func (c *BaseantiClient) DescribeIpResourceInfo(request *baseanti.DescribeIpResourceInfoRequest) (*baseanti.DescribeIpResourceInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeIpResourceInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询攻击记录 */
func (c *BaseantiClient) DescribeAttackLogs(request *baseanti.DescribeAttackLogsRequest) (*baseanti.DescribeAttackLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &baseanti.DescribeAttackLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

