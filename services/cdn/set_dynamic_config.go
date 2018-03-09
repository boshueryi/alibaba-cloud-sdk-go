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

// SetDynamicConfig invokes the cdn.SetDynamicConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setdynamicconfig.html
func (client *Client) SetDynamicConfig(request *SetDynamicConfigRequest) (response *SetDynamicConfigResponse, err error) {
	response = CreateSetDynamicConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetDynamicConfigWithChan invokes the cdn.SetDynamicConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setdynamicconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDynamicConfigWithChan(request *SetDynamicConfigRequest) (<-chan *SetDynamicConfigResponse, <-chan error) {
	responseChan := make(chan *SetDynamicConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDynamicConfig(request)
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

// SetDynamicConfigWithCallback invokes the cdn.SetDynamicConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setdynamicconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDynamicConfigWithCallback(request *SetDynamicConfigRequest, callback func(response *SetDynamicConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDynamicConfigResponse
		var err error
		defer close(result)
		response, err = client.SetDynamicConfig(request)
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

// SetDynamicConfigRequest is the request struct for api SetDynamicConfig
type SetDynamicConfigRequest struct {
	*requests.RpcRequest
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken       string           `position:"Query" name:"SecurityToken"`
	DomainName          string           `position:"Query" name:"DomainName"`
	DynamicOrigin       string           `position:"Query" name:"DynamicOrigin"`
	StaticType          string           `position:"Query" name:"StaticType"`
	StaticUri           string           `position:"Query" name:"StaticUri"`
	StaticPath          string           `position:"Query" name:"StaticPath"`
	DynamicCacheControl string           `position:"Query" name:"DynamicCacheControl"`
}

// SetDynamicConfigResponse is the response struct for api SetDynamicConfig
type SetDynamicConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDynamicConfigRequest creates a request to invoke SetDynamicConfig API
func CreateSetDynamicConfigRequest(request *SetDynamicConfigRequest) {
	request = &SetDynamicConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetDynamicConfig", "", "")
	return
}

// CreateSetDynamicConfigResponse creates a response to parse from SetDynamicConfig response
func CreateSetDynamicConfigResponse() (response *SetDynamicConfigResponse) {
	response = &SetDynamicConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
