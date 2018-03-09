package vpc

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

// DeleteSnatEntry invokes the vpc.DeleteSnatEntry API synchronously
// api document: https://help.aliyun.com/api/vpc/deletesnatentry.html
func (client *Client) DeleteSnatEntry(request *DeleteSnatEntryRequest) (response *DeleteSnatEntryResponse, err error) {
	response = CreateDeleteSnatEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSnatEntryWithChan invokes the vpc.DeleteSnatEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/deletesnatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSnatEntryWithChan(request *DeleteSnatEntryRequest) (<-chan *DeleteSnatEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteSnatEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSnatEntry(request)
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

// DeleteSnatEntryWithCallback invokes the vpc.DeleteSnatEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/deletesnatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSnatEntryWithCallback(request *DeleteSnatEntryRequest, callback func(response *DeleteSnatEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSnatEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteSnatEntry(request)
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

// DeleteSnatEntryRequest is the request struct for api DeleteSnatEntry
type DeleteSnatEntryRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SnatTableId          string           `position:"Query" name:"SnatTableId"`
	SnatEntryId          string           `position:"Query" name:"SnatEntryId"`
}

// DeleteSnatEntryResponse is the response struct for api DeleteSnatEntry
type DeleteSnatEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSnatEntryRequest creates a request to invoke DeleteSnatEntry API
func CreateDeleteSnatEntryRequest(request *DeleteSnatEntryRequest) {
	request = &DeleteSnatEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteSnatEntry", "vpc", "openAPI")
	return
}

// CreateDeleteSnatEntryResponse creates a response to parse from DeleteSnatEntry response
func CreateDeleteSnatEntryResponse() (response *DeleteSnatEntryResponse) {
	response = &DeleteSnatEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
