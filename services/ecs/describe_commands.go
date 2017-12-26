package ecs

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

func (client *Client) DescribeCommands(request *DescribeCommandsRequest) (response *DescribeCommandsResponse, err error) {
	response = CreateDescribeCommandsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCommandsWithChan(request *DescribeCommandsRequest) (<-chan *DescribeCommandsResponse, <-chan error) {
	responseChan := make(chan *DescribeCommandsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCommands(request)
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

func (client *Client) DescribeCommandsWithCallback(request *DescribeCommandsRequest, callback func(response *DescribeCommandsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCommandsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCommands(request)
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

type DescribeCommandsRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	Type                 string `position:"Query" name:"Type"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	Name                 string `position:"Query" name:"Name"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	CommandId            string `position:"Query" name:"CommandId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeCommandsResponse struct {
	*responses.BaseResponse
	RequestId   string           `json:"RequestId" xml:"RequestId"`
	TotalCount  requests.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber  requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize    requests.Integer `json:"PageSize" xml:"PageSize"`
	CommandList struct {
		Command []struct {
			CommandId      string           `json:"CommandId" xml:"CommandId"`
			Name           string           `json:"Name" xml:"Name"`
			Type           string           `json:"Type" xml:"Type"`
			Description    string           `json:"Description" xml:"Description"`
			CommandContent string           `json:"CommandContent" xml:"CommandContent"`
			WorkingDir     string           `json:"WorkingDir" xml:"WorkingDir"`
			Timeout        requests.Integer `json:"Timeout" xml:"Timeout"`
		} `json:"Command" xml:"Command"`
	} `json:"CommandList" xml:"CommandList"`
}

func CreateDescribeCommandsRequest() (request *DescribeCommandsRequest) {
	request = &DescribeCommandsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeCommands", "", "")
	return
}

func CreateDescribeCommandsResponse() (response *DescribeCommandsResponse) {
	response = &DescribeCommandsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
