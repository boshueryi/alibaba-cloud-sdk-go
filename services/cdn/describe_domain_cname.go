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

// DescribeDomainCname invokes the cdn.DescribeDomainCname API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomaincname.html
func (client *Client) DescribeDomainCname(request *DescribeDomainCnameRequest) (response *DescribeDomainCnameResponse, err error) {
	response = CreateDescribeDomainCnameResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainCnameWithChan invokes the cdn.DescribeDomainCname API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaincname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCnameWithChan(request *DescribeDomainCnameRequest) (<-chan *DescribeDomainCnameResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainCnameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainCname(request)
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

// DescribeDomainCnameWithCallback invokes the cdn.DescribeDomainCname API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaincname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainCnameWithCallback(request *DescribeDomainCnameRequest, callback func(response *DescribeDomainCnameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainCnameResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainCname(request)
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

// DescribeDomainCnameRequest is the request struct for api DescribeDomainCname
type DescribeDomainCnameRequest struct {
	*requests.RpcRequest
}

// DescribeDomainCnameResponse is the response struct for api DescribeDomainCname
type DescribeDomainCnameResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	CnameDatas CnameDatas `json:"CnameDatas" xml:"CnameDatas"`
}

// CreateDescribeDomainCnameRequest creates a request to invoke DescribeDomainCname API
func CreateDescribeDomainCnameRequest(request *DescribeDomainCnameRequest) {
	request = &DescribeDomainCnameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCname", "", "")
	return
}

// CreateDescribeDomainCnameResponse creates a response to parse from DescribeDomainCname response
func CreateDescribeDomainCnameResponse() (response *DescribeDomainCnameResponse) {
	response = &DescribeDomainCnameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
