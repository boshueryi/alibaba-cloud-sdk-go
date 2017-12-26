package slb

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

func (client *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
	response = CreateDeleteRulesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteRulesWithChan(request *DeleteRulesRequest) (<-chan *DeleteRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRules(request)
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

func (client *Client) DeleteRulesWithCallback(request *DeleteRulesRequest, callback func(response *DeleteRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteRules(request)
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

type DeleteRulesRequest struct {
	*requests.RpcRequest
	Tags                 string `position:"Query" name:"Tags"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	RuleIds              string `position:"Query" name:"RuleIds"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DeleteRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteRulesRequest() (request *DeleteRulesRequest) {
	request = &DeleteRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteRules", "", "")
	return
}

func CreateDeleteRulesResponse() (response *DeleteRulesResponse) {
	response = &DeleteRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
