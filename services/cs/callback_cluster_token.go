package cs

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

// CallbackClusterToken invokes the cs.CallbackClusterToken API synchronously
// api document: https://help.aliyun.com/api/cs/callbackclustertoken.html
func (client *Client) CallbackClusterToken(request *CallbackClusterTokenRequest) (response *CallbackClusterTokenResponse, err error) {
	response = CreateCallbackClusterTokenResponse()
	err = client.DoAction(request, response)
	return
}

// CallbackClusterTokenWithChan invokes the cs.CallbackClusterToken API asynchronously
// api document: https://help.aliyun.com/api/cs/callbackclustertoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CallbackClusterTokenWithChan(request *CallbackClusterTokenRequest) (<-chan *CallbackClusterTokenResponse, <-chan error) {
	responseChan := make(chan *CallbackClusterTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CallbackClusterToken(request)
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

// CallbackClusterTokenWithCallback invokes the cs.CallbackClusterToken API asynchronously
// api document: https://help.aliyun.com/api/cs/callbackclustertoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CallbackClusterTokenWithCallback(request *CallbackClusterTokenRequest, callback func(response *CallbackClusterTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CallbackClusterTokenResponse
		var err error
		defer close(result)
		response, err = client.CallbackClusterToken(request)
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

// CallbackClusterTokenRequest is the request struct for api CallbackClusterToken
type CallbackClusterTokenRequest struct {
	*requests.RoaRequest
	Token   string `position:"Path" name:"Token"`
	ReqOnce string `position:"Path" name:"ReqOnce"`
}

// CallbackClusterTokenResponse is the response struct for api CallbackClusterToken
type CallbackClusterTokenResponse struct {
	*responses.BaseResponse
}

// CreateCallbackClusterTokenRequest creates a request to invoke CallbackClusterToken API
func CreateCallbackClusterTokenRequest(request *CallbackClusterTokenRequest) {
	request = &CallbackClusterTokenRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "CallbackClusterToken", "/token/[Token]/req_once/[ReqOnce]/callback", "", "")
	request.Method = requests.POST
	return
}

// CreateCallbackClusterTokenResponse creates a response to parse from CallbackClusterToken response
func CreateCallbackClusterTokenResponse() (response *CallbackClusterTokenResponse) {
	response = &CallbackClusterTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
