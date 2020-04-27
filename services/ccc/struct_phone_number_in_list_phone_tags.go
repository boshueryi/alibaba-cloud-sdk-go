package ccc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// PhoneNumberInListPhoneTags is a nested struct in ccc response
type PhoneNumberInListPhoneTags struct {
	PhoneNumberId          string   `json:"PhoneNumberId" xml:"PhoneNumberId"`
	InstanceId             string   `json:"InstanceId" xml:"InstanceId"`
	Number                 string   `json:"Number" xml:"Number"`
	PhoneNumberDescription string   `json:"PhoneNumberDescription" xml:"PhoneNumberDescription"`
	Usage                  string   `json:"Usage" xml:"Usage"`
	Province               string   `json:"Province" xml:"Province"`
	City                   string   `json:"City" xml:"City"`
	Type                   int      `json:"Type" xml:"Type"`
	Concurrency            int      `json:"Concurrency" xml:"Concurrency"`
	ServiceTag             string   `json:"ServiceTag" xml:"ServiceTag"`
	Provider               string   `json:"Provider" xml:"Provider"`
	CreateTime             int64    `json:"CreateTime" xml:"CreateTime"`
	ContactFlowId          string   `json:"ContactFlowId" xml:"ContactFlowId"`
	SkillGroupIdList       []string `json:"SkillGroupIdList" xml:"SkillGroupIdList"`
}