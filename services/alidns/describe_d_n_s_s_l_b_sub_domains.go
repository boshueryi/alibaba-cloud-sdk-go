package alidns

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

func (client *Client) DescribeDNSSLBSubDomains(request *DescribeDNSSLBSubDomainsRequest) (response *DescribeDNSSLBSubDomainsResponse, err error) {
	response = CreateDescribeDNSSLBSubDomainsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDNSSLBSubDomainsWithChan(request *DescribeDNSSLBSubDomainsRequest) (<-chan *DescribeDNSSLBSubDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeDNSSLBSubDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDNSSLBSubDomains(request)
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

func (client *Client) DescribeDNSSLBSubDomainsWithCallback(request *DescribeDNSSLBSubDomainsRequest, callback func(response *DescribeDNSSLBSubDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDNSSLBSubDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDNSSLBSubDomains(request)
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

type DescribeDNSSLBSubDomainsRequest struct {
	*requests.RpcRequest
	PageSize     string `position:"Query" name:"PageSize"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   string `position:"Query" name:"PageNumber"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

type DescribeDNSSLBSubDomainsResponse struct {
	*responses.BaseResponse
	RequestId     string           `json:"RequestId" xml:"RequestId"`
	TotalCount    requests.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber    requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize      requests.Integer `json:"PageSize" xml:"PageSize"`
	SlbSubDomains struct {
		SlbSubDomain []struct {
			SubDomain   string           `json:"SubDomain" xml:"SubDomain"`
			RecordCount requests.Integer `json:"RecordCount" xml:"RecordCount"`
			Open        requests.Boolean `json:"Open" xml:"Open"`
			Type        string           `json:"Type" xml:"Type"`
		} `json:"SlbSubDomain" xml:"SlbSubDomain"`
	} `json:"SlbSubDomains" xml:"SlbSubDomains"`
}

func CreateDescribeDNSSLBSubDomainsRequest() (request *DescribeDNSSLBSubDomainsRequest) {
	request = &DescribeDNSSLBSubDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDNSSLBSubDomains", "", "")
	return
}

func CreateDescribeDNSSLBSubDomainsResponse() (response *DescribeDNSSLBSubDomainsResponse) {
	response = &DescribeDNSSLBSubDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
