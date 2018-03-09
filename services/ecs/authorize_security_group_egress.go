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

// AuthorizeSecurityGroupEgress invokes the ecs.AuthorizeSecurityGroupEgress API synchronously
// api document: https://help.aliyun.com/api/ecs/authorizesecuritygroupegress.html
func (client *Client) AuthorizeSecurityGroupEgress(request *AuthorizeSecurityGroupEgressRequest) (response *AuthorizeSecurityGroupEgressResponse, err error) {
	response = CreateAuthorizeSecurityGroupEgressResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizeSecurityGroupEgressWithChan invokes the ecs.AuthorizeSecurityGroupEgress API asynchronously
// api document: https://help.aliyun.com/api/ecs/authorizesecuritygroupegress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizeSecurityGroupEgressWithChan(request *AuthorizeSecurityGroupEgressRequest) (<-chan *AuthorizeSecurityGroupEgressResponse, <-chan error) {
	responseChan := make(chan *AuthorizeSecurityGroupEgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizeSecurityGroupEgress(request)
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

// AuthorizeSecurityGroupEgressWithCallback invokes the ecs.AuthorizeSecurityGroupEgress API asynchronously
// api document: https://help.aliyun.com/api/ecs/authorizesecuritygroupegress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AuthorizeSecurityGroupEgressWithCallback(request *AuthorizeSecurityGroupEgressRequest, callback func(response *AuthorizeSecurityGroupEgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizeSecurityGroupEgressResponse
		var err error
		defer close(result)
		response, err = client.AuthorizeSecurityGroupEgress(request)
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

// AuthorizeSecurityGroupEgressRequest is the request struct for api AuthorizeSecurityGroupEgress
type AuthorizeSecurityGroupEgressRequest struct {
	*requests.RpcRequest
}

// AuthorizeSecurityGroupEgressResponse is the response struct for api AuthorizeSecurityGroupEgress
type AuthorizeSecurityGroupEgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAuthorizeSecurityGroupEgressRequest creates a request to invoke AuthorizeSecurityGroupEgress API
func CreateAuthorizeSecurityGroupEgressRequest(request *AuthorizeSecurityGroupEgressRequest) {
	request = &AuthorizeSecurityGroupEgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AuthorizeSecurityGroupEgress", "ecs", "openAPI")
	return
}

// CreateAuthorizeSecurityGroupEgressResponse creates a response to parse from AuthorizeSecurityGroupEgress response
func CreateAuthorizeSecurityGroupEgressResponse() (response *AuthorizeSecurityGroupEgressResponse) {
	response = &AuthorizeSecurityGroupEgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
