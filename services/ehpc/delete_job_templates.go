package ehpc

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

// DeleteJobTemplates invokes the ehpc.DeleteJobTemplates API synchronously
// api document: https://help.aliyun.com/api/ehpc/deletejobtemplates.html
func (client *Client) DeleteJobTemplates(request *DeleteJobTemplatesRequest) (response *DeleteJobTemplatesResponse, err error) {
	response = CreateDeleteJobTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteJobTemplatesWithChan invokes the ehpc.DeleteJobTemplates API asynchronously
// api document: https://help.aliyun.com/api/ehpc/deletejobtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteJobTemplatesWithChan(request *DeleteJobTemplatesRequest) (<-chan *DeleteJobTemplatesResponse, <-chan error) {
	responseChan := make(chan *DeleteJobTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteJobTemplates(request)
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

// DeleteJobTemplatesWithCallback invokes the ehpc.DeleteJobTemplates API asynchronously
// api document: https://help.aliyun.com/api/ehpc/deletejobtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteJobTemplatesWithCallback(request *DeleteJobTemplatesRequest, callback func(response *DeleteJobTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteJobTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DeleteJobTemplates(request)
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

// DeleteJobTemplatesRequest is the request struct for api DeleteJobTemplates
type DeleteJobTemplatesRequest struct {
	*requests.RpcRequest
	Templates string `position:"Query" name:"Templates"`
}

// DeleteJobTemplatesResponse is the response struct for api DeleteJobTemplates
type DeleteJobTemplatesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteJobTemplatesRequest creates a request to invoke DeleteJobTemplates API
func CreateDeleteJobTemplatesRequest(request *DeleteJobTemplatesRequest) {
	request = &DeleteJobTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "DeleteJobTemplates", "ehs", "openAPI")
	return
}

// CreateDeleteJobTemplatesResponse creates a response to parse from DeleteJobTemplates response
func CreateDeleteJobTemplatesResponse() (response *DeleteJobTemplatesResponse) {
	response = &DeleteJobTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
