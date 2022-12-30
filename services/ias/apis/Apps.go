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
    ias "github.com/jdcloud-api/jdcloud-sdk-go/services/ias/models"
)

type AppsRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* pin  */
    Pin string `json:"pin"`

    /* appName  */
    AppName string `json:"appName"`

    /* clientId  */
    ClientId string `json:"clientId"`

    /* multiTenant  */
    MultiTenant bool `json:"multiTenant"`

    /* state  */
    State string `json:"state"`

    /* scope  */
    Scope string `json:"scope"`

    /* startTime  */
    StartTime int `json:"startTime"`

    /* endTime  */
    EndTime int `json:"endTime"`

    /* accountType  */
    AccountType string `json:"accountType"`

    /* pageIndex  */
    PageIndex int `json:"pageIndex"`

    /* pageSize  */
    PageSize int `json:"pageSize"`

    /* offset  */
    Offset int `json:"offset"`
}

/*
 * param regionId:  (Required)
 * param pin: pin (Required)
 * param appName: appName (Required)
 * param clientId: clientId (Required)
 * param multiTenant: multiTenant (Required)
 * param state: state (Required)
 * param scope: scope (Required)
 * param startTime: startTime (Required)
 * param endTime: endTime (Required)
 * param accountType: accountType (Required)
 * param pageIndex: pageIndex (Required)
 * param pageSize: pageSize (Required)
 * param offset: offset (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAppsRequest(
    regionId string,
    pin string,
    appName string,
    clientId string,
    multiTenant bool,
    state string,
    scope string,
    startTime int,
    endTime int,
    accountType string,
    pageIndex int,
    pageSize int,
    offset int,
) *AppsRequest {

	return &AppsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/operate_backend/apps",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
        AppName: appName,
        ClientId: clientId,
        MultiTenant: multiTenant,
        State: state,
        Scope: scope,
        StartTime: startTime,
        EndTime: endTime,
        AccountType: accountType,
        PageIndex: pageIndex,
        PageSize: pageSize,
        Offset: offset,
	}
}

/*
 * param regionId:  (Required)
 * param pin: pin (Required)
 * param appName: appName (Required)
 * param clientId: clientId (Required)
 * param multiTenant: multiTenant (Required)
 * param state: state (Required)
 * param scope: scope (Required)
 * param startTime: startTime (Required)
 * param endTime: endTime (Required)
 * param accountType: accountType (Required)
 * param pageIndex: pageIndex (Required)
 * param pageSize: pageSize (Required)
 * param offset: offset (Required)
 */
func NewAppsRequestWithAllParams(
    regionId string,
    pin string,
    appName string,
    clientId string,
    multiTenant bool,
    state string,
    scope string,
    startTime int,
    endTime int,
    accountType string,
    pageIndex int,
    pageSize int,
    offset int,
) *AppsRequest {

    return &AppsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/operate_backend/apps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        AppName: appName,
        ClientId: clientId,
        MultiTenant: multiTenant,
        State: state,
        Scope: scope,
        StartTime: startTime,
        EndTime: endTime,
        AccountType: accountType,
        PageIndex: pageIndex,
        PageSize: pageSize,
        Offset: offset,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAppsRequestWithoutParam() *AppsRequest {

    return &AppsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/operate_backend/apps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AppsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: pin(Required) */
func (r *AppsRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param appName: appName(Required) */
func (r *AppsRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param clientId: clientId(Required) */
func (r *AppsRequest) SetClientId(clientId string) {
    r.ClientId = clientId
}

/* param multiTenant: multiTenant(Required) */
func (r *AppsRequest) SetMultiTenant(multiTenant bool) {
    r.MultiTenant = multiTenant
}

/* param state: state(Required) */
func (r *AppsRequest) SetState(state string) {
    r.State = state
}

/* param scope: scope(Required) */
func (r *AppsRequest) SetScope(scope string) {
    r.Scope = scope
}

/* param startTime: startTime(Required) */
func (r *AppsRequest) SetStartTime(startTime int) {
    r.StartTime = startTime
}

/* param endTime: endTime(Required) */
func (r *AppsRequest) SetEndTime(endTime int) {
    r.EndTime = endTime
}

/* param accountType: accountType(Required) */
func (r *AppsRequest) SetAccountType(accountType string) {
    r.AccountType = accountType
}

/* param pageIndex: pageIndex(Required) */
func (r *AppsRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = pageIndex
}

/* param pageSize: pageSize(Required) */
func (r *AppsRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param offset: offset(Required) */
func (r *AppsRequest) SetOffset(offset int) {
    r.Offset = offset
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AppsRequest) GetRegionId() string {
    return r.RegionId
}

type AppsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AppsResult `json:"result"`
}

type AppsResult struct {
    Pagination ias.Pagination `json:"pagination"`
    Result []ias.AppQueryResultItem `json:"result"`
}