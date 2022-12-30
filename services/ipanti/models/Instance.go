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


type Instance struct {

    /* 实例 ID (Optional) */
    Id string `json:"id"`

    /* 实例名称 (Optional) */
    Name string `json:"name"`

    /* 链路类型. <br>- 1: 电信<br>- 3: 电信、联通和移动<br>- 4: BGP 线路 (Optional) */
    Carrier int `json:"carrier"`

    /* 可防护 IP 类型, 目前仅电信线路支持 IPV6 线路. <br>- 0: IPV4. <br>- 1: IPV4/IPV6 (Optional) */
    IpType int `json:"ipType"`

    /* IP 数量 (Optional) */
    IpCount int `json:"ipCount"`

    /* 可配的转发端口数量 (Optional) */
    PortCount int `json:"portCount"`

    /* 可配的网站规则域名数量 (Optional) */
    DomainCount int `json:"domainCount"`

    /* 触发弹性带宽的次数 (Optional) */
    ElasticTriggerCount int `json:"elasticTriggerCount"`

    /* 超峰次数 (Optional) */
    AbovePeakCount int `json:"abovePeakCount"`

    /* 保底带宽 (Optional) */
    InBitslimit int `json:"inBitslimit"`

    /* 弹性带宽 (Optional) */
    ResilientBitslimit int `json:"resilientBitslimit"`

    /* 业务带宽大小 (Optional) */
    BusinessBitslimit int `json:"businessBitslimit"`

    /* CC 阈值大小 (Optional) */
    CcThreshold int `json:"ccThreshold"`

    /* CC 防护峰值, 单位: QPS (Optional) */
    CcPeakQPS int `json:"ccPeakQPS"`

    /* 非网站类规则数 (Optional) */
    RuleCount int `json:"ruleCount"`

    /* 网站类规则数 (Optional) */
    WebRuleCount int `json:"webRuleCount"`

    /* 防护调度规则数 (Optional) */
    DispatchRuleCount int `json:"dispatchRuleCount"`

    /* 计费状态. <br>- PAID: 已支付<br>- ARREARS: 欠费<br>- EXPIRED: 过期 (Optional) */
    ChargeStatus string `json:"chargeStatus"`

    /* 安全状态. <br>- SAFE: 安全<br>- CLEANING: 清洗中<br>- BLOCKING: 封禁中 (Optional) */
    SecurityStatus string `json:"securityStatus"`

    /* 实例的创建的时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 实例的过期时间 (Optional) */
    ExpireTime string `json:"expireTime"`

    /* 资源 ID, 升级和续费时使用 (Optional) */
    ResourceId string `json:"resourceId"`

    /* CC 防护观察者模式. <br>- 0: 关闭 <br>- 1: 开启 (Optional) */
    CcObserveMode int `json:"ccObserveMode"`

    /* CC 防护模式. <br>- 0: 正常 <br>- 1: 紧急 <br>- 2: 宽松 <br>- 3: 自定义 (Optional) */
    CcProtectMode int `json:"ccProtectMode"`

    /* CC 开关状态. <br>- 0: 关闭 <br>- 1: 开启 (Optional) */
    CcProtectStatus int `json:"ccProtectStatus"`

    /* CC 防护模式为自定义时的限速大小 (Optional) */
    CcSpeedLimit int `json:"ccSpeedLimit"`

    /* CC 防护模式为自定义时的限速周期 (Optional) */
    CcSpeedPeriod int `json:"ccSpeedPeriod"`

    /* IP 黑名单列表 (Optional) */
    IpBlackList []string `json:"ipBlackList"`

    /* IP 黑名单状态. <br>- 0: 关闭 <br>- 1: 开启 (Optional) */
    IpBlackStatus int `json:"ipBlackStatus"`

    /* IP 白名单列表 (Optional) */
    IpWhiteList []string `json:"ipWhiteList"`

    /* IP 白名单状态. <br>- 0: 关闭<br>- 1: 开启 (Optional) */
    IpWhiteStatus int `json:"ipWhiteStatus"`

    /* url白名单列表 (Optional) */
    UrlWhitelist []string `json:"urlWhitelist"`

    /* url白名单状态. <br>- 0: 关闭<br>- 1: 开启 (Optional) */
    UrlWhitelistStatus int `json:"urlWhitelistStatus"`

    /* ccProtectMode为自定义模式时, 每个Host的防护阈值 (Optional) */
    HostQps int `json:"hostQps"`

    /* ccProtectMode为自定义模式时, 每个Host+URI的防护阈值 (Optional) */
    HostUrlQps int `json:"hostUrlQps"`

    /* ccProtectMode为自定义模式时, 每个源IP对Host的防护阈值 (Optional) */
    IpHostQps int `json:"ipHostQps"`

    /* ccProtectMode为自定义模式时, 每个源IP对Host+URI的防护阈值 (Optional) */
    IpHostUrlQps int `json:"ipHostUrlQps"`

    /* 关联的自定义页面id (Optional) */
    PageId string `json:"pageId"`

    /* 关联的自定义页面名称 (Optional) */
    PageName string `json:"pageName"`

    /* 是否开启自定义页面, 关闭时透传状态码.  <br>- 0: 关闭<br>- 1: 开启 (Optional) */
    PageStatus int `json:"pageStatus"`

    /* 每条网站规则可配的http/https端口数 (Optional) */
    WebRulePortLimit int `json:"webRulePortLimit"`

    /* Tag信息 (Optional) */
    Tags []Tag `json:"tags"`
}
