package cdn

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

func (client *Client) DeleteHttpHeaderConfig(request *DeleteHttpHeaderConfigRequest) (response *DeleteHttpHeaderConfigResponse, err error) {
	response = CreateDeleteHttpHeaderConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteHttpHeaderConfigWithChan(request *DeleteHttpHeaderConfigRequest) (<-chan *DeleteHttpHeaderConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteHttpHeaderConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHttpHeaderConfig(request)
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

func (client *Client) DeleteHttpHeaderConfigWithCallback(request *DeleteHttpHeaderConfigRequest, callback func(response *DeleteHttpHeaderConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHttpHeaderConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteHttpHeaderConfig(request)
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

type DeleteHttpHeaderConfigRequest struct {
	*requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DeleteHttpHeaderConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteHttpHeaderConfigRequest() (request *DeleteHttpHeaderConfigRequest) {
	request = &DeleteHttpHeaderConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DeleteHttpHeaderConfig", "", "")
	return
}

func CreateDeleteHttpHeaderConfigResponse() (response *DeleteHttpHeaderConfigResponse) {
	response = &DeleteHttpHeaderConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
