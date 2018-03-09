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

// SetAuthConfig invokes the mts.SetAuthConfig API synchronously
// api document: https://help.aliyun.com/api/mts/setauthconfig.html
func (client *Client) SetAuthConfig(request *SetAuthConfigRequest) (response *SetAuthConfigResponse, err error) {
	response = CreateSetAuthConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetAuthConfigWithChan invokes the mts.SetAuthConfig API asynchronously
// api document: https://help.aliyun.com/api/mts/setauthconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetAuthConfigWithChan(request *SetAuthConfigRequest) (<-chan *SetAuthConfigResponse, <-chan error) {
	responseChan := make(chan *SetAuthConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetAuthConfig(request)
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

// SetAuthConfigWithCallback invokes the mts.SetAuthConfig API asynchronously
// api document: https://help.aliyun.com/api/mts/setauthconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetAuthConfigWithCallback(request *SetAuthConfigRequest, callback func(response *SetAuthConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetAuthConfigResponse
		var err error
		defer close(result)
		response, err = client.SetAuthConfig(request)
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

// SetAuthConfigRequest is the request struct for api SetAuthConfig
type SetAuthConfigRequest struct {
	*requests.RpcRequest
	OwnerId              string `position:"Query" name:"OwnerId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Key1                 string `position:"Query" name:"Key1"`
	Key2                 string `position:"Query" name:"Key2"`
}

// SetAuthConfigResponse is the response struct for api SetAuthConfig
type SetAuthConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Key1      string `json:"Key1" xml:"Key1"`
	Key2      string `json:"Key2" xml:"Key2"`
}

// CreateSetAuthConfigRequest creates a request to invoke SetAuthConfig API
func CreateSetAuthConfigRequest(request *SetAuthConfigRequest) {
	request = &SetAuthConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SetAuthConfig", "mts", "openAPI")
	return
}

// CreateSetAuthConfigResponse creates a response to parse from SetAuthConfig response
func CreateSetAuthConfigResponse() (response *SetAuthConfigResponse) {
	response = &SetAuthConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
