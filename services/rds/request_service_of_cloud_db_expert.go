package rds

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

// RequestServiceOfCloudDBExpert invokes the rds.RequestServiceOfCloudDBExpert API synchronously
// api document: https://help.aliyun.com/api/rds/requestserviceofclouddbexpert.html
func (client *Client) RequestServiceOfCloudDBExpert(request *RequestServiceOfCloudDBExpertRequest) (response *RequestServiceOfCloudDBExpertResponse, err error) {
	response = CreateRequestServiceOfCloudDBExpertResponse()
	err = client.DoAction(request, response)
	return
}

// RequestServiceOfCloudDBExpertWithChan invokes the rds.RequestServiceOfCloudDBExpert API asynchronously
// api document: https://help.aliyun.com/api/rds/requestserviceofclouddbexpert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RequestServiceOfCloudDBExpertWithChan(request *RequestServiceOfCloudDBExpertRequest) (<-chan *RequestServiceOfCloudDBExpertResponse, <-chan error) {
	responseChan := make(chan *RequestServiceOfCloudDBExpertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestServiceOfCloudDBExpert(request)
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

// RequestServiceOfCloudDBExpertWithCallback invokes the rds.RequestServiceOfCloudDBExpert API asynchronously
// api document: https://help.aliyun.com/api/rds/requestserviceofclouddbexpert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RequestServiceOfCloudDBExpertWithCallback(request *RequestServiceOfCloudDBExpertRequest, callback func(response *RequestServiceOfCloudDBExpertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestServiceOfCloudDBExpertResponse
		var err error
		defer close(result)
		response, err = client.RequestServiceOfCloudDBExpert(request)
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

// RequestServiceOfCloudDBExpertRequest is the request struct for api RequestServiceOfCloudDBExpert
type RequestServiceOfCloudDBExpertRequest struct {
	*requests.RpcRequest
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
}

// RequestServiceOfCloudDBExpertResponse is the response struct for api RequestServiceOfCloudDBExpert
type RequestServiceOfCloudDBExpertResponse struct {
	*responses.BaseResponse
	Message string `json:"Message" xml:"Message"`
	Data    string `json:"Data" xml:"Data"`
	Code    string `json:"Code" xml:"Code"`
}

// CreateRequestServiceOfCloudDBExpertRequest creates a request to invoke RequestServiceOfCloudDBExpert API
func CreateRequestServiceOfCloudDBExpertRequest() (request *RequestServiceOfCloudDBExpertRequest) {
	request = &RequestServiceOfCloudDBExpertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RequestServiceOfCloudDBExpert", "Rds", "openAPI")
	return
}

// CreateRequestServiceOfCloudDBExpertResponse creates a response to parse from RequestServiceOfCloudDBExpert response
func CreateRequestServiceOfCloudDBExpertResponse() (response *RequestServiceOfCloudDBExpertResponse) {
	response = &RequestServiceOfCloudDBExpertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}