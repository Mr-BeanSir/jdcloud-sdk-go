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


type SamlAssertionInfo struct {

    /* SAML断言中NameID的格式 (Optional) */
    SubjectType string `json:"subjectType"`

    /* SAML断言中Subject.NameID字段值 (Optional) */
    Subject string `json:"subject"`

    /* SAML断言中Subject.SubjectConfirmation.SubjectConfirmationData字段中Recipient字段值 (Optional) */
    Recipient string `json:"recipient"`

    /* SAML断言中Issuer字段值 (Optional) */
    Issuer string `json:"issuer"`
}
