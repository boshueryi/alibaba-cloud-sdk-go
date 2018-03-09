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

// DeleteVSwitch invokes the ecs.DeleteVSwitch API synchronously
// api document: https://help.aliyun.com/api/ecs/deletevswitch.html
func (client *Client) DeleteVSwitch(request *DeleteVSwitchRequest) (response *DeleteVSwitchResponse, err error) {
	response = CreateDeleteVSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVSwitchWithChan invokes the ecs.DeleteVSwitch API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletevswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVSwitchWithChan(request *DeleteVSwitchRequest) (<-chan *DeleteVSwitchResponse, <-chan error) {
	responseChan := make(chan *DeleteVSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVSwitch(request)
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

// DeleteVSwitchWithCallback invokes the ecs.DeleteVSwitch API asynchronously
// api document: https://help.aliyun.com/api/ecs/deletevswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVSwitchWithCallback(request *DeleteVSwitchRequest, callback func(response *DeleteVSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVSwitchResponse
		var err error
		defer close(result)
		response, err = client.DeleteVSwitch(request)
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

// DeleteVSwitchRequest is the request struct for api DeleteVSwitch
type DeleteVSwitchRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// DeleteVSwitchResponse is the response struct for api DeleteVSwitch
type DeleteVSwitchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVSwitchRequest creates a request to invoke DeleteVSwitch API
func CreateDeleteVSwitchRequest(request *DeleteVSwitchRequest) {
	request = &DeleteVSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteVSwitch", "ecs", "openAPI")
	return
}

// CreateDeleteVSwitchResponse creates a response to parse from DeleteVSwitch response
func CreateDeleteVSwitchResponse() (response *DeleteVSwitchResponse) {
	response = &DeleteVSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
