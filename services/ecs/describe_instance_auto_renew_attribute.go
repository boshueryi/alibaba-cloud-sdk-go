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

func (client *Client) DescribeInstanceAutoRenewAttribute(request *DescribeInstanceAutoRenewAttributeRequest) (response *DescribeInstanceAutoRenewAttributeResponse, err error) {
	response = CreateDescribeInstanceAutoRenewAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstanceAutoRenewAttributeWithChan(request *DescribeInstanceAutoRenewAttributeRequest) (<-chan *DescribeInstanceAutoRenewAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAutoRenewAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAutoRenewAttribute(request)
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

func (client *Client) DescribeInstanceAutoRenewAttributeWithCallback(request *DescribeInstanceAutoRenewAttributeRequest, callback func(response *DescribeInstanceAutoRenewAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAutoRenewAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAutoRenewAttribute(request)
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

type DescribeInstanceAutoRenewAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	InstanceRenewAttributes struct {
		InstanceRenewAttribute []struct {
			InstanceId       string           `json:"InstanceId" xml:"InstanceId"`
			AutoRenewEnabled requests.Boolean `json:"AutoRenewEnabled" xml:"AutoRenewEnabled"`
			Duration         requests.Integer `json:"Duration" xml:"Duration"`
			PeriodUnit       string           `json:"PeriodUnit" xml:"PeriodUnit"`
			RenewalStatus    string           `json:"RenewalStatus" xml:"RenewalStatus"`
		} `json:"InstanceRenewAttribute" xml:"InstanceRenewAttribute"`
	} `json:"InstanceRenewAttributes" xml:"InstanceRenewAttributes"`
}

func CreateDescribeInstanceAutoRenewAttributeRequest() (request *DescribeInstanceAutoRenewAttributeRequest) {
	request = &DescribeInstanceAutoRenewAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceAutoRenewAttribute", "", "")
	return
}

func CreateDescribeInstanceAutoRenewAttributeResponse() (response *DescribeInstanceAutoRenewAttributeResponse) {
	response = &DescribeInstanceAutoRenewAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
