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

// StopCdnDomain invokes the cdn.StopCdnDomain API synchronously
// api document: https://help.aliyun.com/api/cdn/stopcdndomain.html
func (client *Client) StopCdnDomain(request *StopCdnDomainRequest) (response *StopCdnDomainResponse, err error) {
	response = CreateStopCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// StopCdnDomainWithChan invokes the cdn.StopCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/stopcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopCdnDomainWithChan(request *StopCdnDomainRequest) (<-chan *StopCdnDomainResponse, <-chan error) {
	responseChan := make(chan *StopCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopCdnDomain(request)
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

// StopCdnDomainWithCallback invokes the cdn.StopCdnDomain API asynchronously
// api document: https://help.aliyun.com/api/cdn/stopcdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopCdnDomainWithCallback(request *StopCdnDomainRequest, callback func(response *StopCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.StopCdnDomain(request)
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

// StopCdnDomainRequest is the request struct for api StopCdnDomain
type StopCdnDomainRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
}

// StopCdnDomainResponse is the response struct for api StopCdnDomain
type StopCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopCdnDomainRequest creates a request to invoke StopCdnDomain API
func CreateStopCdnDomainRequest(request *StopCdnDomainRequest) {
	request = &StopCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "StopCdnDomain", "", "")
	return
}

// CreateStopCdnDomainResponse creates a response to parse from StopCdnDomain response
func CreateStopCdnDomainResponse() (response *StopCdnDomainResponse) {
	response = &StopCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
