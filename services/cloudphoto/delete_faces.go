package cloudphoto

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

// DeleteFaces invokes the cloudphoto.DeleteFaces API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletefaces.html
func (client *Client) DeleteFaces(request *DeleteFacesRequest) (response *DeleteFacesResponse, err error) {
	response = CreateDeleteFacesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFacesWithChan invokes the cloudphoto.DeleteFaces API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFacesWithChan(request *DeleteFacesRequest) (<-chan *DeleteFacesResponse, <-chan error) {
	responseChan := make(chan *DeleteFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaces(request)
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

// DeleteFacesWithCallback invokes the cloudphoto.DeleteFaces API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/deletefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFacesWithCallback(request *DeleteFacesRequest, callback func(response *DeleteFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFacesResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaces(request)
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

// DeleteFacesRequest is the request struct for api DeleteFaces
type DeleteFacesRequest struct {
	*requests.RpcRequest
	FaceId    *[]string `position:"Query" name:"FaceId"  type:"Repeated"`
	StoreName string    `position:"Query" name:"StoreName"`
	LibraryId string    `position:"Query" name:"LibraryId"`
}

// DeleteFacesResponse is the response struct for api DeleteFaces
type DeleteFacesResponse struct {
	*responses.BaseResponse
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Action    string   `json:"Action" xml:"Action"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateDeleteFacesRequest creates a request to invoke DeleteFaces API
func CreateDeleteFacesRequest(request *DeleteFacesRequest) {
	request = &DeleteFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "DeleteFaces", "cloudphoto", "openAPI")
	return
}

// CreateDeleteFacesResponse creates a response to parse from DeleteFaces response
func CreateDeleteFacesResponse() (response *DeleteFacesResponse) {
	response = &DeleteFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
