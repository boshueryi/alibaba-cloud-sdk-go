package slb

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

// DescribeLoadBalancerAutoReleaseTime invokes the slb.DescribeLoadBalancerAutoReleaseTime API synchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerautoreleasetime.html
func (client *Client) DescribeLoadBalancerAutoReleaseTime(request *DescribeLoadBalancerAutoReleaseTimeRequest) (response *DescribeLoadBalancerAutoReleaseTimeResponse, err error) {
	response = CreateDescribeLoadBalancerAutoReleaseTimeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerAutoReleaseTimeWithChan invokes the slb.DescribeLoadBalancerAutoReleaseTime API asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerautoreleasetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerAutoReleaseTimeWithChan(request *DescribeLoadBalancerAutoReleaseTimeRequest) (<-chan *DescribeLoadBalancerAutoReleaseTimeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerAutoReleaseTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerAutoReleaseTime(request)
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

// DescribeLoadBalancerAutoReleaseTimeWithCallback invokes the slb.DescribeLoadBalancerAutoReleaseTime API asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancerautoreleasetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerAutoReleaseTimeWithCallback(request *DescribeLoadBalancerAutoReleaseTimeRequest, callback func(response *DescribeLoadBalancerAutoReleaseTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerAutoReleaseTimeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerAutoReleaseTime(request)
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

// DescribeLoadBalancerAutoReleaseTimeRequest is the request struct for api DescribeLoadBalancerAutoReleaseTime
type DescribeLoadBalancerAutoReleaseTimeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerAutoReleaseTimeResponse is the response struct for api DescribeLoadBalancerAutoReleaseTime
type DescribeLoadBalancerAutoReleaseTimeResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	AutoReleaseTime int    `json:"AutoReleaseTime" xml:"AutoReleaseTime"`
}

// CreateDescribeLoadBalancerAutoReleaseTimeRequest creates a request to invoke DescribeLoadBalancerAutoReleaseTime API
func CreateDescribeLoadBalancerAutoReleaseTimeRequest(request *DescribeLoadBalancerAutoReleaseTimeRequest) {
	request = &DescribeLoadBalancerAutoReleaseTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerAutoReleaseTime", "slb", "openAPI")
	return
}

// CreateDescribeLoadBalancerAutoReleaseTimeResponse creates a response to parse from DescribeLoadBalancerAutoReleaseTime response
func CreateDescribeLoadBalancerAutoReleaseTimeResponse() (response *DescribeLoadBalancerAutoReleaseTimeResponse) {
	response = &DescribeLoadBalancerAutoReleaseTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
