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

func (client *Client) ListCoverPipeline(request *ListCoverPipelineRequest) (response *ListCoverPipelineResponse, err error) {
	response = CreateListCoverPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListCoverPipelineWithChan(request *ListCoverPipelineRequest) (<-chan *ListCoverPipelineResponse, <-chan error) {
	responseChan := make(chan *ListCoverPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCoverPipeline(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) ListCoverPipelineWithCallback(request *ListCoverPipelineRequest, callback func(response *ListCoverPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCoverPipelineResponse
		var err error
		defer close(result)
		response, err = client.ListCoverPipeline(request)
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

type ListCoverPipelineRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	State                string `position:"Query" name:"State"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type ListCoverPipelineResponse struct {
	*responses.BaseResponse
	RequestId    string           `json:"RequestId" xml:"RequestId"`
	TotalCount   requests.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber   requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
	PipelineList struct {
		Pipeline []struct {
			UserId       requests.Integer `json:"UserId" xml:"UserId"`
			PipelineId   string           `json:"PipelineId" xml:"PipelineId"`
			Name         string           `json:"Name" xml:"Name"`
			State        string           `json:"State" xml:"State"`
			Priority     string           `json:"Priority" xml:"Priority"`
			QuotaNum     requests.Integer `json:"quotaNum" xml:"quotaNum"`
			QuotaUsed    requests.Integer `json:"quotaUsed" xml:"quotaUsed"`
			NotifyConfig string           `json:"NotifyConfig" xml:"NotifyConfig"`
			Role         string           `json:"Role" xml:"Role"`
			ExtendConfig string           `json:"ExtendConfig" xml:"ExtendConfig"`
		} `json:"Pipeline" xml:"Pipeline"`
	} `json:"PipelineList" xml:"PipelineList"`
}

func CreateListCoverPipelineRequest() (request *ListCoverPipelineRequest) {
	request = &ListCoverPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListCoverPipeline", "", "")
	return
}

func CreateListCoverPipelineResponse() (response *ListCoverPipelineResponse) {
	response = &ListCoverPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
