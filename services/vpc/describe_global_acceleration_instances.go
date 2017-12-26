package vpc

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

func (client *Client) DescribeGlobalAccelerationInstances(request *DescribeGlobalAccelerationInstancesRequest) (response *DescribeGlobalAccelerationInstancesResponse, err error) {
	response = CreateDescribeGlobalAccelerationInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeGlobalAccelerationInstancesWithChan(request *DescribeGlobalAccelerationInstancesRequest) (<-chan *DescribeGlobalAccelerationInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeGlobalAccelerationInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGlobalAccelerationInstances(request)
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

func (client *Client) DescribeGlobalAccelerationInstancesWithCallback(request *DescribeGlobalAccelerationInstancesRequest, callback func(response *DescribeGlobalAccelerationInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGlobalAccelerationInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGlobalAccelerationInstances(request)
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

type DescribeGlobalAccelerationInstancesRequest struct {
	*requests.RpcRequest
	PageSize                     string `position:"Query" name:"PageSize"`
	Status                       string `position:"Query" name:"Status"`
	PageNumber                   string `position:"Query" name:"PageNumber"`
	OwnerId                      string `position:"Query" name:"OwnerId"`
	IpAddress                    string `position:"Query" name:"IpAddress"`
	ServerId                     string `position:"Query" name:"ServerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	Name                         string `position:"Query" name:"Name"`
	ServiceLocation              string `position:"Query" name:"ServiceLocation"`
	ResourceOwnerId              string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
}

type DescribeGlobalAccelerationInstancesResponse struct {
	*responses.BaseResponse
	RequestId                   string           `json:"RequestId" xml:"RequestId"`
	TotalCount                  requests.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber                  requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize                    requests.Integer `json:"PageSize" xml:"PageSize"`
	GlobalAccelerationInstances struct {
		GlobalAccelerationInstance []struct {
			RegionId                     string `json:"RegionId" xml:"RegionId"`
			GlobalAccelerationInstanceId string `json:"GlobalAccelerationInstanceId" xml:"GlobalAccelerationInstanceId"`
			IpAddress                    string `json:"IpAddress" xml:"IpAddress"`
			Status                       string `json:"Status" xml:"Status"`
			Bandwidth                    string `json:"Bandwidth" xml:"Bandwidth"`
			InternetChargeType           string `json:"InternetChargeType" xml:"InternetChargeType"`
			ChargeType                   string `json:"ChargeType" xml:"ChargeType"`
			AccelerationLocation         string `json:"AccelerationLocation" xml:"AccelerationLocation"`
			ServiceLocation              string `json:"ServiceLocation" xml:"ServiceLocation"`
			Name                         string `json:"Name" xml:"Name"`
			Description                  string `json:"Description" xml:"Description"`
			ExpiredTime                  string `json:"ExpiredTime" xml:"ExpiredTime"`
			CreationTime                 string `json:"CreationTime" xml:"CreationTime"`
			OperationLocks               struct {
				LockReason []struct {
					LockReason string `json:"LockReason" xml:"LockReason"`
				} `json:"LockReason" xml:"LockReason"`
			} `json:"OperationLocks" xml:"OperationLocks"`
			BackendServers struct {
				BackendServer []struct {
					RegionId        string `json:"RegionId" xml:"RegionId"`
					ServerId        string `json:"ServerId" xml:"ServerId"`
					ServerIpAddress string `json:"ServerIpAddress" xml:"ServerIpAddress"`
					ServerType      string `json:"ServerType" xml:"ServerType"`
				} `json:"BackendServer" xml:"BackendServer"`
			} `json:"BackendServers" xml:"BackendServers"`
		} `json:"GlobalAccelerationInstance" xml:"GlobalAccelerationInstance"`
	} `json:"GlobalAccelerationInstances" xml:"GlobalAccelerationInstances"`
}

func CreateDescribeGlobalAccelerationInstancesRequest() (request *DescribeGlobalAccelerationInstancesRequest) {
	request = &DescribeGlobalAccelerationInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeGlobalAccelerationInstances", "", "")
	return
}

func CreateDescribeGlobalAccelerationInstancesResponse() (response *DescribeGlobalAccelerationInstancesResponse) {
	response = &DescribeGlobalAccelerationInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
