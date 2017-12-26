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

func (client *Client) DescribeInstanceHistoryEvents(request *DescribeInstanceHistoryEventsRequest) (response *DescribeInstanceHistoryEventsResponse, err error) {
	response = CreateDescribeInstanceHistoryEventsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstanceHistoryEventsWithChan(request *DescribeInstanceHistoryEventsRequest) (<-chan *DescribeInstanceHistoryEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceHistoryEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceHistoryEvents(request)
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

func (client *Client) DescribeInstanceHistoryEventsWithCallback(request *DescribeInstanceHistoryEventsRequest, callback func(response *DescribeInstanceHistoryEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceHistoryEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceHistoryEvents(request)
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

type DescribeInstanceHistoryEventsRequest struct {
	*requests.RpcRequest
	PageSize              string    `position:"Query" name:"PageSize"`
	EventId               *[]string `position:"Query" name:"EventId"  type:"Repeated"`
	NotBeforeEnd          string    `position:"Query" name:"NotBefore.End"`
	ResourceOwnerAccount  string    `position:"Query" name:"ResourceOwnerAccount"`
	NotBeforeStart        string    `position:"Query" name:"NotBefore.Start"`
	ResourceOwnerId       string    `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount          string    `position:"Query" name:"OwnerAccount"`
	EventPublishTimeStart string    `position:"Query" name:"EventPublishTime.Start"`
	PageNumber            string    `position:"Query" name:"PageNumber"`
	EventCycleStatus      string    `position:"Query" name:"EventCycleStatus"`
	OwnerId               string    `position:"Query" name:"OwnerId"`
	EventPublishTimeEnd   string    `position:"Query" name:"EventPublishTime.End"`
	EventType             string    `position:"Query" name:"EventType"`
	InstanceId            string    `position:"Query" name:"InstanceId"`
}

type DescribeInstanceHistoryEventsResponse struct {
	*responses.BaseResponse
	RequestId              string           `json:"RequestId" xml:"RequestId"`
	TotalCount             requests.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber             requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize               requests.Integer `json:"PageSize" xml:"PageSize"`
	InstanceSystemEventSet struct {
		InstanceSystemEventType []struct {
			InstanceId       string `json:"InstanceId" xml:"InstanceId"`
			EventId          string `json:"EventId" xml:"EventId"`
			EventPublishTime string `json:"EventPublishTime" xml:"EventPublishTime"`
			NotBefore        string `json:"NotBefore" xml:"NotBefore"`
			EventType        struct {
				Code requests.Integer `json:"Code" xml:"Code"`
				Name string           `json:"Name" xml:"Name"`
			} `json:"EventType" xml:"EventType"`
		} `json:"InstanceSystemEventType" xml:"InstanceSystemEventType"`
	} `json:"InstanceSystemEventSet" xml:"InstanceSystemEventSet"`
}

func CreateDescribeInstanceHistoryEventsRequest() (request *DescribeInstanceHistoryEventsRequest) {
	request = &DescribeInstanceHistoryEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceHistoryEvents", "", "")
	return
}

func CreateDescribeInstanceHistoryEventsResponse() (response *DescribeInstanceHistoryEventsResponse) {
	response = &DescribeInstanceHistoryEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
