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

// DescribeDomainConfigs invokes the cdn.DescribeDomainConfigs API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainconfigs.html
func (client *Client) DescribeDomainConfigs(request *DescribeDomainConfigsRequest) (response *DescribeDomainConfigsResponse, err error) {
	response = CreateDescribeDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainConfigsWithChan invokes the cdn.DescribeDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainConfigsWithChan(request *DescribeDomainConfigsRequest) (<-chan *DescribeDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainConfigs(request)
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

// DescribeDomainConfigsWithCallback invokes the cdn.DescribeDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainConfigsWithCallback(request *DescribeDomainConfigsRequest, callback func(response *DescribeDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainConfigs(request)
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

// DescribeDomainConfigsRequest is the request struct for api DescribeDomainConfigs
type DescribeDomainConfigsRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	ConfigList    string           `position:"Query" name:"ConfigList"`
}

// DescribeDomainConfigsResponse is the response struct for api DescribeDomainConfigs
type DescribeDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	DomainConfigs DomainConfigs `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeDomainConfigsRequest creates a request to invoke DescribeDomainConfigs API
func CreateDescribeDomainConfigsRequest(request *DescribeDomainConfigsRequest) {
	request = &DescribeDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainConfigs", "", "")
	return
}

// CreateDescribeDomainConfigsResponse creates a response to parse from DescribeDomainConfigs response
func CreateDescribeDomainConfigsResponse() (response *DescribeDomainConfigsResponse) {
	response = &DescribeDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
