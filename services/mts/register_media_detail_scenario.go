package mts

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

func (client *Client) RegisterMediaDetailScenario(request *RegisterMediaDetailScenarioRequest) (response *RegisterMediaDetailScenarioResponse, err error) {
	response = CreateRegisterMediaDetailScenarioResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RegisterMediaDetailScenarioWithChan(request *RegisterMediaDetailScenarioRequest) (<-chan *RegisterMediaDetailScenarioResponse, <-chan error) {
	responseChan := make(chan *RegisterMediaDetailScenarioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterMediaDetailScenario(request)
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

func (client *Client) RegisterMediaDetailScenarioWithCallback(request *RegisterMediaDetailScenarioRequest, callback func(response *RegisterMediaDetailScenarioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterMediaDetailScenarioResponse
		var err error
		defer close(result)
		response, err = client.RegisterMediaDetailScenario(request)
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

type RegisterMediaDetailScenarioRequest struct {
	*requests.RpcRequest
	Scenario             string `position:"Query" name:"Scenario"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	JobId                string `position:"Query" name:"JobId"`
}

type RegisterMediaDetailScenarioResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ScenarioId string `json:"ScenarioId" xml:"ScenarioId"`
}

func CreateRegisterMediaDetailScenarioRequest() (request *RegisterMediaDetailScenarioRequest) {
	request = &RegisterMediaDetailScenarioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "RegisterMediaDetailScenario", "", "")
	return
}

func CreateRegisterMediaDetailScenarioResponse() (response *RegisterMediaDetailScenarioResponse) {
	response = &RegisterMediaDetailScenarioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
