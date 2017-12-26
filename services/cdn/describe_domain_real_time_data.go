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

func (client *Client) DescribeDomainRealTimeData(request *DescribeDomainRealTimeDataRequest) (response *DescribeDomainRealTimeDataResponse, err error) {
	response = CreateDescribeDomainRealTimeDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainRealTimeDataWithChan(request *DescribeDomainRealTimeDataRequest) (<-chan *DescribeDomainRealTimeDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeData(request)
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

func (client *Client) DescribeDomainRealTimeDataWithCallback(request *DescribeDomainRealTimeDataRequest, callback func(response *DescribeDomainRealTimeDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeData(request)
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

type DescribeDomainRealTimeDataRequest struct {
	*requests.RpcRequest
	Field         string `position:"Query" name:"Field"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeDomainRealTimeDataResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	DomainName      string `json:"DomainName" xml:"DomainName"`
	Field           string `json:"Field" xml:"Field"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	EndTime         string `json:"EndTime" xml:"EndTime"`
	DataPerInterval struct {
		DataModule []struct {
			TimeStamp string `json:"TimeStamp" xml:"TimeStamp"`
			Value     string `json:"Value" xml:"Value"`
		} `json:"DataModule" xml:"DataModule"`
	} `json:"DataPerInterval" xml:"DataPerInterval"`
}

func CreateDescribeDomainRealTimeDataRequest() (request *DescribeDomainRealTimeDataRequest) {
	request = &DescribeDomainRealTimeDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeData", "", "")
	return
}

func CreateDescribeDomainRealTimeDataResponse() (response *DescribeDomainRealTimeDataResponse) {
	response = &DescribeDomainRealTimeDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
