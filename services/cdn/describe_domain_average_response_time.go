package cdn

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

// DescribeDomainAverageResponseTime invokes the cdn.DescribeDomainAverageResponseTime API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainaverageresponsetime.html
func (client *Client) DescribeDomainAverageResponseTime(request *DescribeDomainAverageResponseTimeRequest) (response *DescribeDomainAverageResponseTimeResponse, err error) {
	response = CreateDescribeDomainAverageResponseTimeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainAverageResponseTimeWithChan invokes the cdn.DescribeDomainAverageResponseTime API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainaverageresponsetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainAverageResponseTimeWithChan(request *DescribeDomainAverageResponseTimeRequest) (<-chan *DescribeDomainAverageResponseTimeResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainAverageResponseTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainAverageResponseTime(request)
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

// DescribeDomainAverageResponseTimeWithCallback invokes the cdn.DescribeDomainAverageResponseTime API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainaverageresponsetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainAverageResponseTimeWithCallback(request *DescribeDomainAverageResponseTimeRequest, callback func(response *DescribeDomainAverageResponseTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainAverageResponseTimeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainAverageResponseTime(request)
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

// DescribeDomainAverageResponseTimeRequest is the request struct for api DescribeDomainAverageResponseTime
type DescribeDomainAverageResponseTimeRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
	TimeMerge     string           `position:"Query" name:"TimeMerge"`
	Interval      string           `position:"Query" name:"Interval"`
}

// DescribeDomainAverageResponseTimeResponse is the response struct for api DescribeDomainAverageResponseTime
type DescribeDomainAverageResponseTimeResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	DomainName       string           `json:"DomainName" xml:"DomainName"`
	DataInterval     string           `json:"DataInterval" xml:"DataInterval"`
	StartTime        string           `json:"StartTime" xml:"StartTime"`
	EndTime          string           `json:"EndTime" xml:"EndTime"`
	AvgRTPerInterval AvgRTPerInterval `json:"AvgRTPerInterval" xml:"AvgRTPerInterval"`
}

// CreateDescribeDomainAverageResponseTimeRequest creates a request to invoke DescribeDomainAverageResponseTime API
func CreateDescribeDomainAverageResponseTimeRequest(request *DescribeDomainAverageResponseTimeRequest) {
	request = &DescribeDomainAverageResponseTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainAverageResponseTime", "", "")
	return
}

// CreateDescribeDomainAverageResponseTimeResponse creates a response to parse from DescribeDomainAverageResponseTime response
func CreateDescribeDomainAverageResponseTimeResponse() (response *DescribeDomainAverageResponseTimeResponse) {
	response = &DescribeDomainAverageResponseTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
