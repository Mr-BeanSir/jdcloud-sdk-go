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


type Details struct {

    /* client名称  */
    Name string `json:"name"`

    /* client的端口  */
    Port string `json:"port"`

    /* client的已连接时长，单位：秒  */
    Age string `json:"age"`

    /* client的空闲时长，单位：秒  */
    Idle string `json:"idle"`

    /* client连接的db  */
    Db string `json:"db"`

    /* 最近执行的命令名称  */
    LastCmd string `json:"lastCmd"`
}
