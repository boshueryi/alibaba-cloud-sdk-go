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

// GetVideoCover invokes the cloudphoto.GetVideoCover API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/getvideocover.html
func (client *Client) GetVideoCover(request *GetVideoCoverRequest) (response *GetVideoCoverResponse, err error) {
	response = CreateGetVideoCoverResponse()
	err = client.DoAction(request, response)
	return
}

// GetVideoCoverWithChan invokes the cloudphoto.GetVideoCover API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getvideocover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoCoverWithChan(request *GetVideoCoverRequest) (<-chan *GetVideoCoverResponse, <-chan error) {
	responseChan := make(chan *GetVideoCoverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVideoCover(request)
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

// GetVideoCoverWithCallback invokes the cloudphoto.GetVideoCover API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getvideocover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoCoverWithCallback(request *GetVideoCoverRequest, callback func(response *GetVideoCoverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVideoCoverResponse
		var err error
		defer close(result)
		response, err = client.GetVideoCover(request)
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

// GetVideoCoverRequest is the request struct for api GetVideoCover
type GetVideoCoverRequest struct {
	*requests.RpcRequest
	PhotoId   requests.Integer `position:"Query" name:"PhotoId"`
	ZoomType  string           `position:"Query" name:"ZoomType"`
	StoreName string           `position:"Query" name:"StoreName"`
	LibraryId string           `position:"Query" name:"LibraryId"`
}

// GetVideoCoverResponse is the response struct for api GetVideoCover
type GetVideoCoverResponse struct {
	*responses.BaseResponse
	Code          string `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	VideoCoverUrl string `json:"VideoCoverUrl" xml:"VideoCoverUrl"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Action        string `json:"Action" xml:"Action"`
}

// CreateGetVideoCoverRequest creates a request to invoke GetVideoCover API
func CreateGetVideoCoverRequest(request *GetVideoCoverRequest) {
	request = &GetVideoCoverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetVideoCover", "cloudphoto", "openAPI")
	return
}

// CreateGetVideoCoverResponse creates a response to parse from GetVideoCover response
func CreateGetVideoCoverResponse() (response *GetVideoCoverResponse) {
	response = &GetVideoCoverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
