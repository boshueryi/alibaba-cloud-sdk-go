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

// DescribeDomainSlowRatio invokes the cdn.DescribeDomainSlowRatio API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainslowratio.html
func (client *Client) DescribeDomainSlowRatio(request *DescribeDomainSlowRatioRequest) (response *DescribeDomainSlowRatioResponse, err error) {
	response = CreateDescribeDomainSlowRatioResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainSlowRatioWithChan invokes the cdn.DescribeDomainSlowRatio API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainslowratio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainSlowRatioWithChan(request *DescribeDomainSlowRatioRequest) (<-chan *DescribeDomainSlowRatioResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainSlowRatioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainSlowRatio(request)
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

// DescribeDomainSlowRatioWithCallback invokes the cdn.DescribeDomainSlowRatio API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainslowratio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainSlowRatioWithCallback(request *DescribeDomainSlowRatioRequest, callback func(response *DescribeDomainSlowRatioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainSlowRatioResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainSlowRatio(request)
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

// DescribeDomainSlowRatioRequest is the request struct for api DescribeDomainSlowRatio
type DescribeDomainSlowRatioRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainSlowRatioResponse is the response struct for api DescribeDomainSlowRatio
type DescribeDomainSlowRatioResponse struct {
	*responses.BaseResponse
	EndTime                  string                   `json:"EndTime" xml:"EndTime"`
	DataInterval             int                      `json:"DataInterval" xml:"DataInterval"`
	PageNumber               int                      `json:"PageNumber" xml:"PageNumber"`
	PageSize                 int                      `json:"PageSize" xml:"PageSize"`
	TotalCount               int                      `json:"TotalCount" xml:"TotalCount"`
	StartTime                string                   `json:"StartTime" xml:"StartTime"`
	SlowRatioDataPerInterval SlowRatioDataPerInterval `json:"SlowRatioDataPerInterval" xml:"SlowRatioDataPerInterval"`
}

// CreateDescribeDomainSlowRatioRequest creates a request to invoke DescribeDomainSlowRatio API
func CreateDescribeDomainSlowRatioRequest() (request *DescribeDomainSlowRatioRequest) {
	request = &DescribeDomainSlowRatioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainSlowRatio", "", "")
	return
}

// CreateDescribeDomainSlowRatioResponse creates a response to parse from DescribeDomainSlowRatio response
func CreateDescribeDomainSlowRatioResponse() (response *DescribeDomainSlowRatioResponse) {
	response = &DescribeDomainSlowRatioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
