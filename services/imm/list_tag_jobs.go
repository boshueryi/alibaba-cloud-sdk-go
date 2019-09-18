package imm

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

// ListTagJobs invokes the imm.ListTagJobs API synchronously
// api document: https://help.aliyun.com/api/imm/listtagjobs.html
func (client *Client) ListTagJobs(request *ListTagJobsRequest) (response *ListTagJobsResponse, err error) {
	response = CreateListTagJobsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagJobsWithChan invokes the imm.ListTagJobs API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagJobsWithChan(request *ListTagJobsRequest) (<-chan *ListTagJobsResponse, <-chan error) {
	responseChan := make(chan *ListTagJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagJobs(request)
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

// ListTagJobsWithCallback invokes the imm.ListTagJobs API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagJobsWithCallback(request *ListTagJobsRequest, callback func(response *ListTagJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagJobsResponse
		var err error
		defer close(result)
		response, err = client.ListTagJobs(request)
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

// ListTagJobsRequest is the request struct for api ListTagJobs
type ListTagJobsRequest struct {
	*requests.RpcRequest
	MaxKeys   requests.Integer `position:"Query" name:"MaxKeys"`
	Project   string           `position:"Query" name:"Project"`
	Condition string           `position:"Query" name:"Condition"`
	Marker    string           `position:"Query" name:"Marker"`
}

// ListTagJobsResponse is the response struct for api ListTagJobs
type ListTagJobsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	NextMarker string     `json:"NextMarker" xml:"NextMarker"`
	Jobs       []JobsItem `json:"Jobs" xml:"Jobs"`
}

// CreateListTagJobsRequest creates a request to invoke ListTagJobs API
func CreateListTagJobsRequest() (request *ListTagJobsRequest) {
	request = &ListTagJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListTagJobs", "imm", "openAPI")
	return
}

// CreateListTagJobsResponse creates a response to parse from ListTagJobs response
func CreateListTagJobsResponse() (response *ListTagJobsResponse) {
	response = &ListTagJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
