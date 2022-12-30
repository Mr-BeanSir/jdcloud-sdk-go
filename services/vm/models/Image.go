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


type Image struct {

    /* 镜像ID。 (Optional) */
    ImageId string `json:"imageId"`

    /* 镜像名称。 (Optional) */
    Name string `json:"name"`

    /* 镜像的操作系统平台名称。
取值范围：`Ubuntu、CentOS、Windows Server、Other Linux、Other Windows`。
 (Optional) */
    Platform string `json:"platform"`

    /* 镜像的操作系统版本。 (Optional) */
    OsVersion string `json:"osVersion"`

    /* 镜像架构。取值范围：`x86_64、arm64`。 (Optional) */
    Architecture string `json:"architecture"`

    /* 镜像系统盘大小。 (Optional) */
    SystemDiskSizeGB int `json:"systemDiskSizeGB"`

    /* 镜像来源，取值范围：
`public`：官方镜像。
`thirdparty`：镜像市场镜像。
`private`：用户自己的私有镜像。
`shared`：其他用户分享的镜像。
`community`：社区镜像。
 (Optional) */
    ImageSource string `json:"imageSource"`

    /* 镜像的操作系统类型。取值范围：`windows、linux`。 (Optional) */
    OsType string `json:"osType"`

    /* 镜像状态。参考 [镜像状态](https://docs.jdcloud.com/virtual-machines/api/image_status)。 (Optional) */
    Status string `json:"status"`

    /* 镜像的创建时间。 (Optional) */
    CreateTime string `json:"createTime"`

    /* 镜像文件的实际大小。 (Optional) */
    SizeMB int `json:"sizeMB"`

    /* 镜像描述。 (Optional) */
    Desc string `json:"desc"`

    /* 该镜像拥有者的用户PIN。 (Optional) */
    OwnerPin string `json:"ownerPin"`

    /* 镜像的使用权限。取值范围：
`all`：没有限制，所有人均可以使用。
`specifiedUsers`：只有共享用户可以使用。
`ownerOnly`：镜像拥有者自己可以使用。
 (Optional) */
    LaunchPermission string `json:"launchPermission"`

    /* 镜像系统盘配置。 (Optional) */
    SystemDisk InstanceDiskAttachment `json:"systemDisk"`

    /* 镜像数据盘配置列表。 (Optional) */
    DataDisks []InstanceDiskAttachment `json:"dataDisks"`

    /* 创建云盘系统盘所使用的快照ID。系统盘类型为本地盘的镜像，此参数为空。 (Optional) */
    SnapshotId string `json:"snapshotId"`

    /* 镜像支持的系统盘类型。取值范围：
`localDisk`：本地盘系统盘。
`cloudDisk`：云盘系统盘。
 (Optional) */
    RootDeviceType string `json:"rootDeviceType"`

    /* 镜像复制和转换时的进度，仅显示数值，单位为百分比。 (Optional) */
    Progress string `json:"progress"`

    /* 镜像的上下线状态。`offline=true` 的镜像不再允许创建云主机。 (Optional) */
    Offline bool `json:"offline"`

    /* 已废弃。 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 是否来自导入镜像。 (Optional) */
    Imported bool `json:"imported"`
}
