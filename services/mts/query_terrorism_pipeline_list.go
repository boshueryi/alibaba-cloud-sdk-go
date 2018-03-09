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

// QueryTerrorismPipelineList invokes the mts.QueryTerrorismPipelineList API synchronously
// api document: https://help.aliyun.com/api/mts/queryterrorismpipelinelist.html
func (client *Client) QueryTerrorismPipelineList(request *QueryTerrorismPipelineListRequest) (response *QueryTerrorismPipelineListResponse, err error) {
	response = CreateQueryTerrorismPipelineListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTerrorismPipelineListWithChan invokes the mts.QueryTerrorismPipelineList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryterrorismpipelinelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTerrorismPipelineListWithChan(request *QueryTerrorismPipelineListRequest) (<-chan *QueryTerrorismPipelineListResponse, <-chan error) {
	responseChan := make(chan *QueryTerrorismPipelineListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTerrorismPipelineList(request)
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

// QueryTerrorismPipelineListWithCallback invokes the mts.QueryTerrorismPipelineList API asynchronously
// api document: https://help.aliyun.com/api/mts/queryterrorismpipelinelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTerrorismPipelineListWithCallback(request *QueryTerrorismPipelineListRequest, callback func(response *QueryTerrorismPipelineListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTerrorismPipelineListResponse
		var err error
		defer close(result)
		response, err = client.QueryTerrorismPipelineList(request)
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

// QueryTerrorismPipelineListRequest is the request struct for api QueryTerrorismPipelineList
type QueryTerrorismPipelineListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PipelineIds          string           `position:"Query" name:"PipelineIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// QueryTerrorismPipelineListResponse is the response struct for api QueryTerrorismPipelineList
type QueryTerrorismPipelineListResponse struct {
	*responses.BaseResponse
	RequestId    string                                   `json:"RequestId" xml:"RequestId"`
	NonExistIds  NonExistIdsInQueryTerrorismPipelineList  `json:"NonExistIds" xml:"NonExistIds"`
	PipelineList PipelineListInQueryTerrorismPipelineList `json:"PipelineList" xml:"PipelineList"`
}

// CreateQueryTerrorismPipelineListRequest creates a request to invoke QueryTerrorismPipelineList API
func CreateQueryTerrorismPipelineListRequest(request *QueryTerrorismPipelineListRequest) {
	request = &QueryTerrorismPipelineListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryTerrorismPipelineList", "mts", "openAPI")
	return
}

// CreateQueryTerrorismPipelineListResponse creates a response to parse from QueryTerrorismPipelineList response
func CreateQueryTerrorismPipelineListResponse() (response *QueryTerrorismPipelineListResponse) {
	response = &QueryTerrorismPipelineListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
