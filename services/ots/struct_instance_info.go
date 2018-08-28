package ots

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

// InstanceInfo is a nested struct in ots response
type InstanceInfo struct {
	Status        int                   `json:"Status" xml:"Status"`
	WriteCapacity int                   `json:"WriteCapacity" xml:"WriteCapacity"`
	ReadCapacity  int                   `json:"ReadCapacity" xml:"ReadCapacity"`
	ClusterType   string                `json:"ClusterType" xml:"ClusterType"`
	Timestamp     string                `json:"Timestamp" xml:"Timestamp"`
	UserId        string                `json:"UserId" xml:"UserId"`
	InstanceName  string                `json:"InstanceName" xml:"InstanceName"`
	CreateTime    string                `json:"CreateTime" xml:"CreateTime"`
	Network       string                `json:"Network" xml:"Network"`
	Description   string                `json:"Description" xml:"Description"`
	Quota         Quota                 `json:"Quota" xml:"Quota"`
	TagInfos      TagInfosInGetInstance `json:"TagInfos" xml:"TagInfos"`
}