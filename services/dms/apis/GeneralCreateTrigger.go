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
    dms "github.com/jdcloud-api/jdcloud-sdk-go/services/dms/models"
)

type GeneralCreateTriggerRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id。 (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 数据库名称。 (Optional) */
    DbName *string `json:"dbName"`

    /* 触发器名称。 (Optional) */
    TriggerName *string `json:"triggerName"`

    /* 触发时机，BEFORE("BEFORE", 1),AFTER("AFTER", 2)。 (Optional) */
    TriggerTiming *string `json:"triggerTiming"`

    /* 激活触发器的事件，INSERT("INSERT", 1),UPDATE("UPDATE", 2), DELETE("DELETE", 3)。 (Optional) */
    TriggerEvent *string `json:"triggerEvent"`

    /* 触发表。 (Optional) */
    TriggerTable *string `json:"triggerTable"`

    /* 触发器定义。 (Optional) */
    TriggerStatement *string `json:"triggerStatement"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGeneralCreateTriggerRequest(
    regionId string,
) *GeneralCreateTriggerRequest {

	return &GeneralCreateTriggerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/trigger:generalCreate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param dataSourceId: 数据源id。 (Optional)
 * param dbName: 数据库名称。 (Optional)
 * param triggerName: 触发器名称。 (Optional)
 * param triggerTiming: 触发时机，BEFORE("BEFORE", 1),AFTER("AFTER", 2)。 (Optional)
 * param triggerEvent: 激活触发器的事件，INSERT("INSERT", 1),UPDATE("UPDATE", 2), DELETE("DELETE", 3)。 (Optional)
 * param triggerTable: 触发表。 (Optional)
 * param triggerStatement: 触发器定义。 (Optional)
 */
func NewGeneralCreateTriggerRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    dbName *string,
    triggerName *string,
    triggerTiming *string,
    triggerEvent *string,
    triggerTable *string,
    triggerStatement *string,
) *GeneralCreateTriggerRequest {

    return &GeneralCreateTriggerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/trigger:generalCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        DbName: dbName,
        TriggerName: triggerName,
        TriggerTiming: triggerTiming,
        TriggerEvent: triggerEvent,
        TriggerTable: triggerTable,
        TriggerStatement: triggerStatement,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGeneralCreateTriggerRequestWithoutParam() *GeneralCreateTriggerRequest {

    return &GeneralCreateTriggerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/trigger:generalCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *GeneralCreateTriggerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param dataSourceId: 数据源id。(Optional) */
func (r *GeneralCreateTriggerRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}
/* param dbName: 数据库名称。(Optional) */
func (r *GeneralCreateTriggerRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}
/* param triggerName: 触发器名称。(Optional) */
func (r *GeneralCreateTriggerRequest) SetTriggerName(triggerName string) {
    r.TriggerName = &triggerName
}
/* param triggerTiming: 触发时机，BEFORE("BEFORE", 1),AFTER("AFTER", 2)。(Optional) */
func (r *GeneralCreateTriggerRequest) SetTriggerTiming(triggerTiming string) {
    r.TriggerTiming = &triggerTiming
}
/* param triggerEvent: 激活触发器的事件，INSERT("INSERT", 1),UPDATE("UPDATE", 2), DELETE("DELETE", 3)。(Optional) */
func (r *GeneralCreateTriggerRequest) SetTriggerEvent(triggerEvent string) {
    r.TriggerEvent = &triggerEvent
}
/* param triggerTable: 触发表。(Optional) */
func (r *GeneralCreateTriggerRequest) SetTriggerTable(triggerTable string) {
    r.TriggerTable = &triggerTable
}
/* param triggerStatement: 触发器定义。(Optional) */
func (r *GeneralCreateTriggerRequest) SetTriggerStatement(triggerStatement string) {
    r.TriggerStatement = &triggerStatement
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GeneralCreateTriggerRequest) GetRegionId() string {
    return r.RegionId
}

type GeneralCreateTriggerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GeneralCreateTriggerResult `json:"result"`
}

type GeneralCreateTriggerResult struct {
    DmsSqls []dms.DmsSql `json:"dmsSqls"`
}