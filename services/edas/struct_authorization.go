package edas

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

// Authorization is a nested struct in edas response
type Authorization struct {
	AdminUserId    string `json:"AdminUserId" xml:"AdminUserId"`
	AdminEdasId    string `json:"AdminEdasId" xml:"AdminEdasId"`
	UserId         string `json:"UserId" xml:"UserId"`
	EdasId         string `json:"EdasId" xml:"EdasId"`
	AppIdData      string `json:"AppIdData" xml:"AppIdData"`
	RoleIdData     string `json:"RoleIdData" xml:"RoleIdData"`
	CreateTime     int64  `json:"CreateTime" xml:"CreateTime"`
	UpdateTime     int64  `json:"UpdateTime" xml:"UpdateTime"`
	ResGroupId     int64  `json:"ResGroupId" xml:"ResGroupId"`
	ResGroupIdData string `json:"ResGroupIdData" xml:"ResGroupIdData"`
	IsRamSlave     bool   `json:"IsRamSlave" xml:"IsRamSlave"`
	IsRamDel       bool   `json:"IsRamDel" xml:"IsRamDel"`
	SubUserKp      string `json:"SubUserKp" xml:"SubUserKp"`
	RamOperation   bool   `json:"RamOperation" xml:"RamOperation"`
	DelegateAdmin  bool   `json:"DelegateAdmin" xml:"DelegateAdmin"`
	Sts            bool   `json:"Sts" xml:"Sts"`
}
