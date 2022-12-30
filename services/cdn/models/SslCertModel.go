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


type SslCertModel struct {

    /* 证书Id (Optional) */
    SslCertId string `json:"sslCertId"`

    /* 证书名称 (Optional) */
    CertName string `json:"certName"`

    /* 主域名 (Optional) */
    CommonName string `json:"commonName"`

    /* 证书类型 (Optional) */
    CertType string `json:"certType"`

    /* 开始时间 (Optional) */
    SslCertStartTime string `json:"sslCertStartTime"`

    /* 结束时间 (Optional) */
    SslCertEndTime string `json:"sslCertEndTime"`

    /* 是否允许被删除,1允许,0不允许 (Optional) */
    Deletable int `json:"deletable"`

    /* 对私钥文件使用sha256算法计算的摘要信息 (Optional) */
    Digest string `json:"digest"`

    /* 证书别名 (Optional) */
    AliasName string `json:"aliasName"`

    /* 备用域名 (Optional) */
    RelatedDomains []string `json:"relatedDomains"`
}
