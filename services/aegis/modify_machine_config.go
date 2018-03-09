package aegis

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

// ModifyMachineConfig invokes the aegis.ModifyMachineConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/modifymachineconfig.html
func (client *Client) ModifyMachineConfig(request *ModifyMachineConfigRequest) (response *ModifyMachineConfigResponse, err error) {
	response = CreateModifyMachineConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMachineConfigWithChan invokes the aegis.ModifyMachineConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifymachineconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMachineConfigWithChan(request *ModifyMachineConfigRequest) (<-chan *ModifyMachineConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyMachineConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMachineConfig(request)
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

// ModifyMachineConfigWithCallback invokes the aegis.ModifyMachineConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifymachineconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMachineConfigWithCallback(request *ModifyMachineConfigRequest, callback func(response *ModifyMachineConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMachineConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyMachineConfig(request)
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

// ModifyMachineConfigRequest is the request struct for api ModifyMachineConfig
type ModifyMachineConfigRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Config          string           `position:"Query" name:"Config"`
	Type            string           `position:"Query" name:"Type"`
	Target          string           `position:"Query" name:"Target"`
	Lang            string           `position:"Query" name:"Lang"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// ModifyMachineConfigResponse is the response struct for api ModifyMachineConfig
type ModifyMachineConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMachineConfigRequest creates a request to invoke ModifyMachineConfig API
func CreateModifyMachineConfigRequest(request *ModifyMachineConfigRequest) {
	request = &ModifyMachineConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyMachineConfig", "vipaegis", "openAPI")
	return
}

// CreateModifyMachineConfigResponse creates a response to parse from ModifyMachineConfig response
func CreateModifyMachineConfigResponse() (response *ModifyMachineConfigResponse) {
	response = &ModifyMachineConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
