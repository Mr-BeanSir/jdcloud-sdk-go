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
    detection "github.com/jdcloud-api/jdcloud-sdk-go/services/detection/models"
)

type UpdateSiteMonitorRequest struct {

    core.JDCloudRequest

    /* 地址  */
    Address string `json:"address"`

    /*  (Optional) */
    AdvanceChecked *string `json:"advanceChecked"`

    /*  (Optional) */
    CreatedTime *int64 `json:"createdTime"`

    /* 探测频率  */
    Cycle int64 `json:"cycle"`

    /*  (Optional) */
    DefaultSource *string `json:"defaultSource"`

    /*  (Optional) */
    DnsOption *detection.SiteMonitorDnsOption `json:"dnsOption"`

    /*  (Optional) */
    Enabled *string `json:"enabled"`

    /*  (Optional) */
    FtpOption *detection.SiteMonitorFtpOption `json:"ftpOption"`

    /*  (Optional) */
    HawkeyeId *int64 `json:"hawkeyeId"`

    /*  (Optional) */
    HttpOption *detection.SiteMonitorHttpOption `json:"httpOption"`

    /*  (Optional) */
    Id *string `json:"id"`

    /*  (Optional) */
    IsDeleted *string `json:"isDeleted"`

    /* 任务名称  */
    Name string `json:"name"`

    /*  (Optional) */
    Pin *string `json:"pin"`

    /*  (Optional) */
    PingOption *detection.SiteMonitorPingOption `json:"pingOption"`

    /*  (Optional) */
    Pop3Option *detection.SiteMonitorPop3Option `json:"pop3Option"`

    /* 端口 (Optional) */
    Port *string `json:"port"`

    /*  (Optional) */
    SmtpOption *detection.SiteMonitorSmtpOption `json:"smtpOption"`

    /* 探测源  */
    Source []detection.SiteMonitorSource `json:"source"`

    /*  (Optional) */
    Stats *interface{} `json:"stats"`

    /* 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP  */
    TaskType string `json:"taskType"`

    /*  (Optional) */
    TcpOption *detection.SiteMonitorTcpOption `json:"tcpOption"`

    /*  (Optional) */
    UdpOption *detection.SiteMonitorUdpOption `json:"udpOption"`

    /*  (Optional) */
    UpdatedTime *int64 `json:"updatedTime"`
}

/*
 * param address: 地址 (Required)
 * param cycle: 探测频率 (Required)
 * param name: 任务名称 (Required)
 * param source: 探测源 (Required)
 * param taskType: 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateSiteMonitorRequest(
    address string,
    cycle int64,
    name string,
    source []detection.SiteMonitorSource,
    taskType string,
) *UpdateSiteMonitorRequest {

	return &UpdateSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/siteMonitor",
			Method:  "PUT",
			Header:  nil,
			Version: "v2",
		},
        Address: address,
        Cycle: cycle,
        Name: name,
        Source: source,
        TaskType: taskType,
	}
}

/*
 * param address: 地址 (Required)
 * param advanceChecked:  (Optional)
 * param createdTime:  (Optional)
 * param cycle: 探测频率 (Required)
 * param defaultSource:  (Optional)
 * param dnsOption:  (Optional)
 * param enabled:  (Optional)
 * param ftpOption:  (Optional)
 * param hawkeyeId:  (Optional)
 * param httpOption:  (Optional)
 * param id:  (Optional)
 * param isDeleted:  (Optional)
 * param name: 任务名称 (Required)
 * param pin:  (Optional)
 * param pingOption:  (Optional)
 * param pop3Option:  (Optional)
 * param port: 端口 (Optional)
 * param smtpOption:  (Optional)
 * param source: 探测源 (Required)
 * param stats:  (Optional)
 * param taskType: 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP (Required)
 * param tcpOption:  (Optional)
 * param udpOption:  (Optional)
 * param updatedTime:  (Optional)
 */
func NewUpdateSiteMonitorRequestWithAllParams(
    address string,
    advanceChecked *string,
    createdTime *int64,
    cycle int64,
    defaultSource *string,
    dnsOption *detection.SiteMonitorDnsOption,
    enabled *string,
    ftpOption *detection.SiteMonitorFtpOption,
    hawkeyeId *int64,
    httpOption *detection.SiteMonitorHttpOption,
    id *string,
    isDeleted *string,
    name string,
    pin *string,
    pingOption *detection.SiteMonitorPingOption,
    pop3Option *detection.SiteMonitorPop3Option,
    port *string,
    smtpOption *detection.SiteMonitorSmtpOption,
    source []detection.SiteMonitorSource,
    stats *interface{},
    taskType string,
    tcpOption *detection.SiteMonitorTcpOption,
    udpOption *detection.SiteMonitorUdpOption,
    updatedTime *int64,
) *UpdateSiteMonitorRequest {

    return &UpdateSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
        Address: address,
        AdvanceChecked: advanceChecked,
        CreatedTime: createdTime,
        Cycle: cycle,
        DefaultSource: defaultSource,
        DnsOption: dnsOption,
        Enabled: enabled,
        FtpOption: ftpOption,
        HawkeyeId: hawkeyeId,
        HttpOption: httpOption,
        Id: id,
        IsDeleted: isDeleted,
        Name: name,
        Pin: pin,
        PingOption: pingOption,
        Pop3Option: pop3Option,
        Port: port,
        SmtpOption: smtpOption,
        Source: source,
        Stats: stats,
        TaskType: taskType,
        TcpOption: tcpOption,
        UdpOption: udpOption,
        UpdatedTime: updatedTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateSiteMonitorRequestWithoutParam() *UpdateSiteMonitorRequest {

    return &UpdateSiteMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param address: 地址(Required) */
func (r *UpdateSiteMonitorRequest) SetAddress(address string) {
    r.Address = address
}

/* param advanceChecked: (Optional) */
func (r *UpdateSiteMonitorRequest) SetAdvanceChecked(advanceChecked string) {
    r.AdvanceChecked = &advanceChecked
}

/* param createdTime: (Optional) */
func (r *UpdateSiteMonitorRequest) SetCreatedTime(createdTime int64) {
    r.CreatedTime = &createdTime
}

/* param cycle: 探测频率(Required) */
func (r *UpdateSiteMonitorRequest) SetCycle(cycle int64) {
    r.Cycle = cycle
}

/* param defaultSource: (Optional) */
func (r *UpdateSiteMonitorRequest) SetDefaultSource(defaultSource string) {
    r.DefaultSource = &defaultSource
}

/* param dnsOption: (Optional) */
func (r *UpdateSiteMonitorRequest) SetDnsOption(dnsOption *detection.SiteMonitorDnsOption) {
    r.DnsOption = dnsOption
}

/* param enabled: (Optional) */
func (r *UpdateSiteMonitorRequest) SetEnabled(enabled string) {
    r.Enabled = &enabled
}

/* param ftpOption: (Optional) */
func (r *UpdateSiteMonitorRequest) SetFtpOption(ftpOption *detection.SiteMonitorFtpOption) {
    r.FtpOption = ftpOption
}

/* param hawkeyeId: (Optional) */
func (r *UpdateSiteMonitorRequest) SetHawkeyeId(hawkeyeId int64) {
    r.HawkeyeId = &hawkeyeId
}

/* param httpOption: (Optional) */
func (r *UpdateSiteMonitorRequest) SetHttpOption(httpOption *detection.SiteMonitorHttpOption) {
    r.HttpOption = httpOption
}

/* param id: (Optional) */
func (r *UpdateSiteMonitorRequest) SetId(id string) {
    r.Id = &id
}

/* param isDeleted: (Optional) */
func (r *UpdateSiteMonitorRequest) SetIsDeleted(isDeleted string) {
    r.IsDeleted = &isDeleted
}

/* param name: 任务名称(Required) */
func (r *UpdateSiteMonitorRequest) SetName(name string) {
    r.Name = name
}

/* param pin: (Optional) */
func (r *UpdateSiteMonitorRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param pingOption: (Optional) */
func (r *UpdateSiteMonitorRequest) SetPingOption(pingOption *detection.SiteMonitorPingOption) {
    r.PingOption = pingOption
}

/* param pop3Option: (Optional) */
func (r *UpdateSiteMonitorRequest) SetPop3Option(pop3Option *detection.SiteMonitorPop3Option) {
    r.Pop3Option = pop3Option
}

/* param port: 端口(Optional) */
func (r *UpdateSiteMonitorRequest) SetPort(port string) {
    r.Port = &port
}

/* param smtpOption: (Optional) */
func (r *UpdateSiteMonitorRequest) SetSmtpOption(smtpOption *detection.SiteMonitorSmtpOption) {
    r.SmtpOption = smtpOption
}

/* param source: 探测源(Required) */
func (r *UpdateSiteMonitorRequest) SetSource(source []detection.SiteMonitorSource) {
    r.Source = source
}

/* param stats: (Optional) */
func (r *UpdateSiteMonitorRequest) SetStats(stats interface{}) {
    r.Stats = &stats
}

/* param taskType: 任务类型，可选值：HTTP、PING 、TCP 、UDP、DNS、SMTP、POP3和FTP(Required) */
func (r *UpdateSiteMonitorRequest) SetTaskType(taskType string) {
    r.TaskType = taskType
}

/* param tcpOption: (Optional) */
func (r *UpdateSiteMonitorRequest) SetTcpOption(tcpOption *detection.SiteMonitorTcpOption) {
    r.TcpOption = tcpOption
}

/* param udpOption: (Optional) */
func (r *UpdateSiteMonitorRequest) SetUdpOption(udpOption *detection.SiteMonitorUdpOption) {
    r.UdpOption = udpOption
}

/* param updatedTime: (Optional) */
func (r *UpdateSiteMonitorRequest) SetUpdatedTime(updatedTime int64) {
    r.UpdatedTime = &updatedTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateSiteMonitorRequest) GetRegionId() string {
    return ""
}

type UpdateSiteMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateSiteMonitorResult `json:"result"`
}

type UpdateSiteMonitorResult struct {
    Suc bool `json:"suc"`
}