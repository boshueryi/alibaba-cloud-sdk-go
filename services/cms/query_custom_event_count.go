package cms

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

// QueryCustomEventCount invokes the cms.QueryCustomEventCount API synchronously
// api document: https://help.aliyun.com/api/cms/querycustomeventcount.html
func (client *Client) QueryCustomEventCount(request *QueryCustomEventCountRequest) (response *QueryCustomEventCountResponse, err error) {
	response = CreateQueryCustomEventCountResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCustomEventCountWithChan invokes the cms.QueryCustomEventCount API asynchronously
// api document: https://help.aliyun.com/api/cms/querycustomeventcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCustomEventCountWithChan(request *QueryCustomEventCountRequest) (<-chan *QueryCustomEventCountResponse, <-chan error) {
	responseChan := make(chan *QueryCustomEventCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCustomEventCount(request)
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

// QueryCustomEventCountWithCallback invokes the cms.QueryCustomEventCount API asynchronously
// api document: https://help.aliyun.com/api/cms/querycustomeventcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCustomEventCountWithCallback(request *QueryCustomEventCountRequest, callback func(response *QueryCustomEventCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCustomEventCountResponse
		var err error
		defer close(result)
		response, err = client.QueryCustomEventCount(request)
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

// QueryCustomEventCountRequest is the request struct for api QueryCustomEventCount
type QueryCustomEventCountRequest struct {
	*requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

// QueryCustomEventCountResponse is the response struct for api QueryCustomEventCount
type QueryCustomEventCountResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateQueryCustomEventCountRequest creates a request to invoke QueryCustomEventCount API
func CreateQueryCustomEventCountRequest(request *QueryCustomEventCountRequest) {
	request = &QueryCustomEventCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "QueryCustomEventCount", "cms", "openAPI")
	return
}

// CreateQueryCustomEventCountResponse creates a response to parse from QueryCustomEventCount response
func CreateQueryCustomEventCountResponse() (response *QueryCustomEventCountResponse) {
	response = &QueryCustomEventCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
