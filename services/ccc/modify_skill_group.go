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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifySkillGroup invokes the ccc.ModifySkillGroup API synchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroup.html
func (client *Client) ModifySkillGroup(request *ModifySkillGroupRequest) (response *ModifySkillGroupResponse, err error) {
	response = CreateModifySkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySkillGroupWithChan invokes the ccc.ModifySkillGroup API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySkillGroupWithChan(request *ModifySkillGroupRequest) (<-chan *ModifySkillGroupResponse, <-chan error) {
	responseChan := make(chan *ModifySkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySkillGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifySkillGroupWithCallback invokes the ccc.ModifySkillGroup API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySkillGroupWithCallback(request *ModifySkillGroupRequest, callback func(response *ModifySkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySkillGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifySkillGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifySkillGroupRequest is the request struct for api ModifySkillGroup
type ModifySkillGroupRequest struct {
	*requests.RpcRequest
	InstanceId            string    `position:"Query" name:"InstanceId"`
	SkillGroupId          string    `position:"Query" name:"SkillGroupId"`
	Name                  string    `position:"Query" name:"Name"`
	Description           string    `position:"Query" name:"Description"`
	OutboundPhoneNumberId *[]string `position:"Query" name:"OutboundPhoneNumberId"  type:"Repeated"`
	UserId                *[]string `position:"Query" name:"UserId"  type:"Repeated"`
	SkillLevel            *[]string `position:"Query" name:"SkillLevel"  type:"Repeated"`
}

// ModifySkillGroupResponse is the response struct for api ModifySkillGroup
type ModifySkillGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifySkillGroupRequest creates a request to invoke ModifySkillGroup API
func CreateModifySkillGroupRequest(request *ModifySkillGroupRequest) {
	request = &ModifySkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifySkillGroup", "", "")
	return
}

// CreateModifySkillGroupResponse creates a response to parse from ModifySkillGroup response
func CreateModifySkillGroupResponse() (response *ModifySkillGroupResponse) {
	response = &ModifySkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
