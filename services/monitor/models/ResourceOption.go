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


type ResourceOption struct {

    /* 指定具体资源ID设置报警规则，每次最多100个。优先resourceItems生效 (Optional) */
    ResourceItems []ResourceItem `json:"resourceItems"`

    /*  (Optional) */
    TagsOption *TagsOption `json:"tagsOption"`

    /* 指定资源组设置报警规则 (Optional) */
    ResourceGroups []string `json:"resourceGroups"`

    /* 资源筛选的类型,1:指定具体资源 2:标签筛选 3:资源组筛选 (Optional) */
    ResourceFilterType *int64 `json:"resourceFilterType"`
}
