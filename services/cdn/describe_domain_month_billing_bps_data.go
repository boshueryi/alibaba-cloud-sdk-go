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

// DescribeDomainMonthBillingBpsData invokes the cdn.DescribeDomainMonthBillingBpsData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainmonthbillingbpsdata.html
func (client *Client) DescribeDomainMonthBillingBpsData(request *DescribeDomainMonthBillingBpsDataRequest) (response *DescribeDomainMonthBillingBpsDataResponse, err error) {
	response = CreateDescribeDomainMonthBillingBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainMonthBillingBpsDataWithChan invokes the cdn.DescribeDomainMonthBillingBpsData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainmonthbillingbpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainMonthBillingBpsDataWithChan(request *DescribeDomainMonthBillingBpsDataRequest) (<-chan *DescribeDomainMonthBillingBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainMonthBillingBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainMonthBillingBpsData(request)
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

// DescribeDomainMonthBillingBpsDataWithCallback invokes the cdn.DescribeDomainMonthBillingBpsData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainmonthbillingbpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainMonthBillingBpsDataWithCallback(request *DescribeDomainMonthBillingBpsDataRequest, callback func(response *DescribeDomainMonthBillingBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainMonthBillingBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainMonthBillingBpsData(request)
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

// DescribeDomainMonthBillingBpsDataRequest is the request struct for api DescribeDomainMonthBillingBpsData
type DescribeDomainMonthBillingBpsDataRequest struct {
	*requests.RpcRequest
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken      string           `position:"Query" name:"SecurityToken"`
	DomainName         string           `position:"Query" name:"DomainName"`
	StartTime          string           `position:"Query" name:"StartTime"`
	EndTime            string           `position:"Query" name:"EndTime"`
	InternetChargeType string           `position:"Query" name:"InternetChargeType"`
}

// DescribeDomainMonthBillingBpsDataResponse is the response struct for api DescribeDomainMonthBillingBpsData
type DescribeDomainMonthBillingBpsDataResponse struct {
	*responses.BaseResponse
	RequestId              string  `json:"RequestId" xml:"RequestId"`
	DomainName             string  `json:"DomainName" xml:"DomainName"`
	StartTime              string  `json:"StartTime" xml:"StartTime"`
	EndTime                string  `json:"EndTime" xml:"EndTime"`
	Bandwidth95Bps         float64 `json:"Bandwidth95Bps" xml:"Bandwidth95Bps"`
	DomesticBandwidth95Bps float64 `json:"DomesticBandwidth95Bps" xml:"DomesticBandwidth95Bps"`
	OverseasBandwidth95Bps float64 `json:"OverseasBandwidth95Bps" xml:"OverseasBandwidth95Bps"`
	MonthavgBps            float64 `json:"MonthavgBps" xml:"MonthavgBps"`
	DomesticMonthavgBps    float64 `json:"DomesticMonthavgBps" xml:"DomesticMonthavgBps"`
	OverseasMonthavgBps    float64 `json:"OverseasMonthavgBps" xml:"OverseasMonthavgBps"`
	Month4thBps            float64 `json:"Month4thBps" xml:"Month4thBps"`
	DomesticMonth4thBps    float64 `json:"DomesticMonth4thBps" xml:"DomesticMonth4thBps"`
	OverseasMonth4thBps    float64 `json:"OverseasMonth4thBps" xml:"OverseasMonth4thBps"`
}

// CreateDescribeDomainMonthBillingBpsDataRequest creates a request to invoke DescribeDomainMonthBillingBpsData API
func CreateDescribeDomainMonthBillingBpsDataRequest(request *DescribeDomainMonthBillingBpsDataRequest) {
	request = &DescribeDomainMonthBillingBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainMonthBillingBpsData", "", "")
	return
}

// CreateDescribeDomainMonthBillingBpsDataResponse creates a response to parse from DescribeDomainMonthBillingBpsData response
func CreateDescribeDomainMonthBillingBpsDataResponse() (response *DescribeDomainMonthBillingBpsDataResponse) {
	response = &DescribeDomainMonthBillingBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
