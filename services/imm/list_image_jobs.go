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

// ListImageJobs invokes the imm.ListImageJobs API synchronously
// api document: https://help.aliyun.com/api/imm/listimagejobs.html
func (client *Client) ListImageJobs(request *ListImageJobsRequest) (response *ListImageJobsResponse, err error) {
	response = CreateListImageJobsResponse()
	err = client.DoAction(request, response)
	return
}

// ListImageJobsWithChan invokes the imm.ListImageJobs API asynchronously
// api document: https://help.aliyun.com/api/imm/listimagejobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListImageJobsWithChan(request *ListImageJobsRequest) (<-chan *ListImageJobsResponse, <-chan error) {
	responseChan := make(chan *ListImageJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListImageJobs(request)
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

// ListImageJobsWithCallback invokes the imm.ListImageJobs API asynchronously
// api document: https://help.aliyun.com/api/imm/listimagejobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListImageJobsWithCallback(request *ListImageJobsRequest, callback func(response *ListImageJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListImageJobsResponse
		var err error
		defer close(result)
		response, err = client.ListImageJobs(request)
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

// ListImageJobsRequest is the request struct for api ListImageJobs
type ListImageJobsRequest struct {
	*requests.RpcRequest
	MaxKeys requests.Integer `position:"Query" name:"MaxKeys"`
	Project string           `position:"Query" name:"Project"`
	JobType string           `position:"Query" name:"JobType"`
	Marker  string           `position:"Query" name:"Marker"`
}

// ListImageJobsResponse is the response struct for api ListImageJobs
type ListImageJobsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	NextMarker string     `json:"NextMarker" xml:"NextMarker"`
	Jobs       []JobsItem `json:"Jobs" xml:"Jobs"`
}

// CreateListImageJobsRequest creates a request to invoke ListImageJobs API
func CreateListImageJobsRequest() (request *ListImageJobsRequest) {
	request = &ListImageJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListImageJobs", "imm", "openAPI")
	return
}

// CreateListImageJobsResponse creates a response to parse from ListImageJobs response
func CreateListImageJobsResponse() (response *ListImageJobsResponse) {
	response = &ListImageJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
