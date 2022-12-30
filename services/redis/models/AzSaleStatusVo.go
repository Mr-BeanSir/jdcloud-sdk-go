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


type AzSaleStatusVo struct {

    /* 物理az (Optional) */
    Az string `json:"az"`

    /* 是否售罄(0未售罄，1售罄) (Optional) */
    CanSale int `json:"canSale"`

    /* 是否可见(1可见，0不可见) (Optional) */
    Visible int `json:"visible"`
}
