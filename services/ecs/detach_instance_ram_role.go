package ecs

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

func (client *Client) DetachInstanceRamRole(request *DetachInstanceRamRoleRequest) (response *DetachInstanceRamRoleResponse, err error) {
	response = CreateDetachInstanceRamRoleResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DetachInstanceRamRoleWithChan(request *DetachInstanceRamRoleRequest) (<-chan *DetachInstanceRamRoleResponse, <-chan error) {
	responseChan := make(chan *DetachInstanceRamRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachInstanceRamRole(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DetachInstanceRamRoleWithCallback(request *DetachInstanceRamRoleRequest, callback func(response *DetachInstanceRamRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachInstanceRamRoleResponse
		var err error
		defer close(result)
		response, err = client.DetachInstanceRamRole(request)
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

type DetachInstanceRamRoleRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DetachInstanceRamRoleResponse struct {
	*responses.BaseResponse
	RequestId                    string           `json:"RequestId" xml:"RequestId"`
	TotalCount                   requests.Integer `json:"TotalCount" xml:"TotalCount"`
	FailCount                    requests.Integer `json:"FailCount" xml:"FailCount"`
	RamRoleName                  string           `json:"RamRoleName" xml:"RamRoleName"`
	DetachInstanceRamRoleResults struct {
		DetachInstanceRamRoleResult []struct {
			InstanceId          string           `json:"InstanceId" xml:"InstanceId"`
			Success             requests.Boolean `json:"Success" xml:"Success"`
			Code                string           `json:"Code" xml:"Code"`
			Message             string           `json:"Message" xml:"Message"`
			InstanceRamRoleSets struct {
				InstanceRamRoleSet []struct {
					InstanceId  string `json:"InstanceId" xml:"InstanceId"`
					RamRoleName string `json:"RamRoleName" xml:"RamRoleName"`
				} `json:"InstanceRamRoleSet" xml:"InstanceRamRoleSet"`
			} `json:"InstanceRamRoleSets" xml:"InstanceRamRoleSets"`
		} `json:"DetachInstanceRamRoleResult" xml:"DetachInstanceRamRoleResult"`
	} `json:"DetachInstanceRamRoleResults" xml:"DetachInstanceRamRoleResults"`
}

func CreateDetachInstanceRamRoleRequest() (request *DetachInstanceRamRoleRequest) {
	request = &DetachInstanceRamRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DetachInstanceRamRole", "", "")
	return
}

func CreateDetachInstanceRamRoleResponse() (response *DetachInstanceRamRoleResponse) {
	response = &DetachInstanceRamRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
