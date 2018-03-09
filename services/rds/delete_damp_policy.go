package rds

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

// DeleteDampPolicy invokes the rds.DeleteDampPolicy API synchronously
// api document: https://help.aliyun.com/api/rds/deletedamppolicy.html
func (client *Client) DeleteDampPolicy(request *DeleteDampPolicyRequest) (response *DeleteDampPolicyResponse, err error) {
	response = CreateDeleteDampPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDampPolicyWithChan invokes the rds.DeleteDampPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/deletedamppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDampPolicyWithChan(request *DeleteDampPolicyRequest) (<-chan *DeleteDampPolicyResponse, <-chan error) {
	responseChan := make(chan *DeleteDampPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDampPolicy(request)
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

// DeleteDampPolicyWithCallback invokes the rds.DeleteDampPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/deletedamppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDampPolicyWithCallback(request *DeleteDampPolicyRequest, callback func(response *DeleteDampPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDampPolicyResponse
		var err error
		defer close(result)
		response, err = client.DeleteDampPolicy(request)
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

// DeleteDampPolicyRequest is the request struct for api DeleteDampPolicy
type DeleteDampPolicyRequest struct {
	*requests.RpcRequest
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	PolicyName           string           `position:"Query" name:"PolicyName"`
}

// DeleteDampPolicyResponse is the response struct for api DeleteDampPolicy
type DeleteDampPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDampPolicyRequest creates a request to invoke DeleteDampPolicy API
func CreateDeleteDampPolicyRequest(request *DeleteDampPolicyRequest) {
	request = &DeleteDampPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteDampPolicy", "rds", "openAPI")
	return
}

// CreateDeleteDampPolicyResponse creates a response to parse from DeleteDampPolicy response
func CreateDeleteDampPolicyResponse() (response *DeleteDampPolicyResponse) {
	response = &DeleteDampPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
