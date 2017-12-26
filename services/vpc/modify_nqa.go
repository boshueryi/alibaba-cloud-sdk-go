package vpc

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

func (client *Client) ModifyNqa(request *ModifyNqaRequest) (response *ModifyNqaResponse, err error) {
	response = CreateModifyNqaResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyNqaWithChan(request *ModifyNqaRequest) (<-chan *ModifyNqaResponse, <-chan error) {
	responseChan := make(chan *ModifyNqaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNqa(request)
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

func (client *Client) ModifyNqaWithCallback(request *ModifyNqaRequest, callback func(response *ModifyNqaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNqaResponse
		var err error
		defer close(result)
		response, err = client.ModifyNqa(request)
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

type ModifyNqaRequest struct {
	*requests.RpcRequest
	ClientToken          string `position:"Query" name:"ClientToken"`
	DestinationIp        string `position:"Query" name:"DestinationIp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	NqaId                string `position:"Query" name:"NqaId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type ModifyNqaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyNqaRequest() (request *ModifyNqaRequest) {
	request = &ModifyNqaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNqa", "", "")
	return
}

func CreateModifyNqaResponse() (response *ModifyNqaResponse) {
	response = &ModifyNqaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
