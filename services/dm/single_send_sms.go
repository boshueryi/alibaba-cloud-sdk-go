package dm

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

func (client *Client) SingleSendSms(request *SingleSendSmsRequest) (response *SingleSendSmsResponse, err error) {
	response = CreateSingleSendSmsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SingleSendSmsWithChan(request *SingleSendSmsRequest) (<-chan *SingleSendSmsResponse, <-chan error) {
	responseChan := make(chan *SingleSendSmsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SingleSendSms(request)
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

func (client *Client) SingleSendSmsWithCallback(request *SingleSendSmsRequest, callback func(response *SingleSendSmsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SingleSendSmsResponse
		var err error
		defer close(result)
		response, err = client.SingleSendSms(request)
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

type SingleSendSmsRequest struct {
	*requests.RpcRequest
	TemplateCode         string `position:"Query" name:"TemplateCode"`
	SignName             string `position:"Query" name:"SignName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecNum               string `position:"Query" name:"RecNum"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Version              string `position:"Query" name:"Version"`
	ParamString          string `position:"Query" name:"ParamString"`
}

type SingleSendSmsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSingleSendSmsRequest() (request *SingleSendSmsRequest) {
	request = &SingleSendSmsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "SingleSendSms", "", "")
	return
}

func CreateSingleSendSmsResponse() (response *SingleSendSmsResponse) {
	response = &SingleSendSmsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
