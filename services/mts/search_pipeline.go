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

// SearchPipeline invokes the mts.SearchPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/searchpipeline.html
func (client *Client) SearchPipeline(request *SearchPipelineRequest) (response *SearchPipelineResponse, err error) {
	response = CreateSearchPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// SearchPipelineWithChan invokes the mts.SearchPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/searchpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchPipelineWithChan(request *SearchPipelineRequest) (<-chan *SearchPipelineResponse, <-chan error) {
	responseChan := make(chan *SearchPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchPipeline(request)
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

// SearchPipelineWithCallback invokes the mts.SearchPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/searchpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchPipelineWithCallback(request *SearchPipelineRequest, callback func(response *SearchPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchPipelineResponse
		var err error
		defer close(result)
		response, err = client.SearchPipeline(request)
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

// SearchPipelineRequest is the request struct for api SearchPipeline
type SearchPipelineRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	State                string           `position:"Query" name:"State"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// SearchPipelineResponse is the response struct for api SearchPipeline
type SearchPipelineResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	TotalCount   int                          `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int                          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                          `json:"PageSize" xml:"PageSize"`
	PipelineList PipelineListInSearchPipeline `json:"PipelineList" xml:"PipelineList"`
}

// CreateSearchPipelineRequest creates a request to invoke SearchPipeline API
func CreateSearchPipelineRequest(request *SearchPipelineRequest) {
	request = &SearchPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SearchPipeline", "mts", "openAPI")
	return
}

// CreateSearchPipelineResponse creates a response to parse from SearchPipeline response
func CreateSearchPipelineResponse() (response *SearchPipelineResponse) {
	response = &SearchPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
