package push

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

func (client *Client) QueryUniqueDeviceStat(request *QueryUniqueDeviceStatRequest) (response *QueryUniqueDeviceStatResponse, err error) {
	response = CreateQueryUniqueDeviceStatResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryUniqueDeviceStatWithChan(request *QueryUniqueDeviceStatRequest) (<-chan *QueryUniqueDeviceStatResponse, <-chan error) {
	responseChan := make(chan *QueryUniqueDeviceStatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUniqueDeviceStat(request)
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

func (client *Client) QueryUniqueDeviceStatWithCallback(request *QueryUniqueDeviceStatRequest, callback func(response *QueryUniqueDeviceStatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUniqueDeviceStatResponse
		var err error
		defer close(result)
		response, err = client.QueryUniqueDeviceStat(request)
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

type QueryUniqueDeviceStatRequest struct {
	*requests.RpcRequest
	EndTime     string `position:"Query" name:"EndTime"`
	StartTime   string `position:"Query" name:"StartTime"`
	Granularity string `position:"Query" name:"Granularity"`
	AppKey      string `position:"Query" name:"AppKey"`
}

type QueryUniqueDeviceStatResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	AppDeviceStats struct {
		AppDeviceStat []struct {
			Time  string           `json:"Time" xml:"Time"`
			Count requests.Integer `json:"Count" xml:"Count"`
		} `json:"AppDeviceStat" xml:"AppDeviceStat"`
	} `json:"AppDeviceStats" xml:"AppDeviceStats"`
}

func CreateQueryUniqueDeviceStatRequest() (request *QueryUniqueDeviceStatRequest) {
	request = &QueryUniqueDeviceStatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryUniqueDeviceStat", "", "")
	return
}

func CreateQueryUniqueDeviceStatResponse() (response *QueryUniqueDeviceStatResponse) {
	response = &QueryUniqueDeviceStatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
