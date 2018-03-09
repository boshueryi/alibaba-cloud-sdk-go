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

// ModifyInstanceAutoReleaseTime invokes the ecs.ModifyInstanceAutoReleaseTime API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstanceautoreleasetime.html
func (client *Client) ModifyInstanceAutoReleaseTime(request *ModifyInstanceAutoReleaseTimeRequest) (response *ModifyInstanceAutoReleaseTimeResponse, err error) {
	response = CreateModifyInstanceAutoReleaseTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceAutoReleaseTimeWithChan invokes the ecs.ModifyInstanceAutoReleaseTime API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstanceautoreleasetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceAutoReleaseTimeWithChan(request *ModifyInstanceAutoReleaseTimeRequest) (<-chan *ModifyInstanceAutoReleaseTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceAutoReleaseTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceAutoReleaseTime(request)
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

// ModifyInstanceAutoReleaseTimeWithCallback invokes the ecs.ModifyInstanceAutoReleaseTime API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyinstanceautoreleasetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceAutoReleaseTimeWithCallback(request *ModifyInstanceAutoReleaseTimeRequest, callback func(response *ModifyInstanceAutoReleaseTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceAutoReleaseTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceAutoReleaseTime(request)
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

// ModifyInstanceAutoReleaseTimeRequest is the request struct for api ModifyInstanceAutoReleaseTime
type ModifyInstanceAutoReleaseTimeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	AutoReleaseTime      string           `position:"Query" name:"AutoReleaseTime"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// ModifyInstanceAutoReleaseTimeResponse is the response struct for api ModifyInstanceAutoReleaseTime
type ModifyInstanceAutoReleaseTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceAutoReleaseTimeRequest creates a request to invoke ModifyInstanceAutoReleaseTime API
func CreateModifyInstanceAutoReleaseTimeRequest(request *ModifyInstanceAutoReleaseTimeRequest) {
	request = &ModifyInstanceAutoReleaseTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAutoReleaseTime", "ecs", "openAPI")
	return
}

// CreateModifyInstanceAutoReleaseTimeResponse creates a response to parse from ModifyInstanceAutoReleaseTime response
func CreateModifyInstanceAutoReleaseTimeResponse() (response *ModifyInstanceAutoReleaseTimeResponse) {
	response = &ModifyInstanceAutoReleaseTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
