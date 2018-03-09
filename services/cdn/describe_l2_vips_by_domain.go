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

// DescribeL2VipsByDomain invokes the cdn.DescribeL2VipsByDomain API synchronously
// api document: https://help.aliyun.com/api/cdn/describel2vipsbydomain.html
func (client *Client) DescribeL2VipsByDomain(request *DescribeL2VipsByDomainRequest) (response *DescribeL2VipsByDomainResponse, err error) {
	response = CreateDescribeL2VipsByDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeL2VipsByDomainWithChan invokes the cdn.DescribeL2VipsByDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/describel2vipsbydomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeL2VipsByDomainWithChan(request *DescribeL2VipsByDomainRequest) (<-chan *DescribeL2VipsByDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeL2VipsByDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeL2VipsByDomain(request)
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

// DescribeL2VipsByDomainWithCallback invokes the cdn.DescribeL2VipsByDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/describel2vipsbydomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeL2VipsByDomainWithCallback(request *DescribeL2VipsByDomainRequest, callback func(response *DescribeL2VipsByDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeL2VipsByDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeL2VipsByDomain(request)
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

// DescribeL2VipsByDomainRequest is the request struct for api DescribeL2VipsByDomain
type DescribeL2VipsByDomainRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
}

// DescribeL2VipsByDomainResponse is the response struct for api DescribeL2VipsByDomain
type DescribeL2VipsByDomainResponse struct {
	*responses.BaseResponse
	RequestId  string                       `json:"RequestId" xml:"RequestId"`
	DomainName string                       `json:"DomainName" xml:"DomainName"`
	Vips       VipsInDescribeL2VipsByDomain `json:"Vips" xml:"Vips"`
}

// CreateDescribeL2VipsByDomainRequest creates a request to invoke DescribeL2VipsByDomain API
func CreateDescribeL2VipsByDomainRequest(request *DescribeL2VipsByDomainRequest) {
	request = &DescribeL2VipsByDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeL2VipsByDomain", "", "")
	return
}

// CreateDescribeL2VipsByDomainResponse creates a response to parse from DescribeL2VipsByDomain response
func CreateDescribeL2VipsByDomainResponse() (response *DescribeL2VipsByDomainResponse) {
	response = &DescribeL2VipsByDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
