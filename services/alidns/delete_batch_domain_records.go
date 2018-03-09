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

// DeleteBatchDomainRecords invokes the alidns.DeleteBatchDomainRecords API synchronously
// api document: https://help.aliyun.com/api/alidns/deletebatchdomainrecords.html
func (client *Client) DeleteBatchDomainRecords(request *DeleteBatchDomainRecordsRequest) (response *DeleteBatchDomainRecordsResponse, err error) {
	response = CreateDeleteBatchDomainRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBatchDomainRecordsWithChan invokes the alidns.DeleteBatchDomainRecords API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletebatchdomainrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBatchDomainRecordsWithChan(request *DeleteBatchDomainRecordsRequest) (<-chan *DeleteBatchDomainRecordsResponse, <-chan error) {
	responseChan := make(chan *DeleteBatchDomainRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBatchDomainRecords(request)
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

// DeleteBatchDomainRecordsWithCallback invokes the alidns.DeleteBatchDomainRecords API asynchronously
// api document: https://help.aliyun.com/api/alidns/deletebatchdomainrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBatchDomainRecordsWithCallback(request *DeleteBatchDomainRecordsRequest, callback func(response *DeleteBatchDomainRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBatchDomainRecordsResponse
		var err error
		defer close(result)
		response, err = client.DeleteBatchDomainRecords(request)
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

// DeleteBatchDomainRecordsRequest is the request struct for api DeleteBatchDomainRecords
type DeleteBatchDomainRecordsRequest struct {
	*requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Records      string `position:"Query" name:"Records"`
}

// DeleteBatchDomainRecordsResponse is the response struct for api DeleteBatchDomainRecords
type DeleteBatchDomainRecordsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
}

// CreateDeleteBatchDomainRecordsRequest creates a request to invoke DeleteBatchDomainRecords API
func CreateDeleteBatchDomainRecordsRequest(request *DeleteBatchDomainRecordsRequest) {
	request = &DeleteBatchDomainRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DeleteBatchDomainRecords", "", "")
	return
}

// CreateDeleteBatchDomainRecordsResponse creates a response to parse from DeleteBatchDomainRecords response
func CreateDeleteBatchDomainRecordsResponse() (response *DeleteBatchDomainRecordsResponse) {
	response = &DeleteBatchDomainRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
