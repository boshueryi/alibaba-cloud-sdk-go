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

// ModifyDampPolicy invokes the rds.ModifyDampPolicy API synchronously
// api document: https://help.aliyun.com/api/rds/modifydamppolicy.html
func (client *Client) ModifyDampPolicy(request *ModifyDampPolicyRequest) (response *ModifyDampPolicyResponse, err error) {
	response = CreateModifyDampPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDampPolicyWithChan invokes the rds.ModifyDampPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydamppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDampPolicyWithChan(request *ModifyDampPolicyRequest) (<-chan *ModifyDampPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyDampPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDampPolicy(request)
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

// ModifyDampPolicyWithCallback invokes the rds.ModifyDampPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydamppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDampPolicyWithCallback(request *ModifyDampPolicyRequest, callback func(response *ModifyDampPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDampPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyDampPolicy(request)
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

// ModifyDampPolicyRequest is the request struct for api ModifyDampPolicy
type ModifyDampPolicyRequest struct {
	*requests.RpcRequest
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	PolicyName           string           `position:"Query" name:"PolicyName"`
	Handlers             string           `position:"Query" name:"Handlers"`
	SourceRules          string           `position:"Query" name:"SourceRules"`
	TimeRules            string           `position:"Query" name:"TimeRules"`
	ActionRules          string           `position:"Query" name:"ActionRules"`
}

// ModifyDampPolicyResponse is the response struct for api ModifyDampPolicy
type ModifyDampPolicyResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PolicyId   string `json:"PolicyId" xml:"PolicyId"`
	PolicyName string `json:"PolicyName" xml:"PolicyName"`
}

// CreateModifyDampPolicyRequest creates a request to invoke ModifyDampPolicy API
func CreateModifyDampPolicyRequest(request *ModifyDampPolicyRequest) {
	request = &ModifyDampPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDampPolicy", "rds", "openAPI")
	return
}

// CreateModifyDampPolicyResponse creates a response to parse from ModifyDampPolicy response
func CreateModifyDampPolicyResponse() (response *ModifyDampPolicyResponse) {
	response = &ModifyDampPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
