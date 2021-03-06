package live

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

// DeleteHtmlResource invokes the live.DeleteHtmlResource API synchronously
// api document: https://help.aliyun.com/api/live/deletehtmlresource.html
func (client *Client) DeleteHtmlResource(request *DeleteHtmlResourceRequest) (response *DeleteHtmlResourceResponse, err error) {
	response = CreateDeleteHtmlResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHtmlResourceWithChan invokes the live.DeleteHtmlResource API asynchronously
// api document: https://help.aliyun.com/api/live/deletehtmlresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHtmlResourceWithChan(request *DeleteHtmlResourceRequest) (<-chan *DeleteHtmlResourceResponse, <-chan error) {
	responseChan := make(chan *DeleteHtmlResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHtmlResource(request)
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

// DeleteHtmlResourceWithCallback invokes the live.DeleteHtmlResource API asynchronously
// api document: https://help.aliyun.com/api/live/deletehtmlresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHtmlResourceWithCallback(request *DeleteHtmlResourceRequest, callback func(response *DeleteHtmlResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHtmlResourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteHtmlResource(request)
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

// DeleteHtmlResourceRequest is the request struct for api DeleteHtmlResource
type DeleteHtmlResourceRequest struct {
	*requests.RpcRequest
	HtmlUrl        string           `position:"Query" name:"htmlUrl"`
	CasterId       string           `position:"Query" name:"CasterId"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	HtmlResourceId string           `position:"Query" name:"HtmlResourceId"`
}

// DeleteHtmlResourceResponse is the response struct for api DeleteHtmlResource
type DeleteHtmlResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteHtmlResourceRequest creates a request to invoke DeleteHtmlResource API
func CreateDeleteHtmlResourceRequest() (request *DeleteHtmlResourceRequest) {
	request = &DeleteHtmlResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteHtmlResource", "live", "openAPI")
	return
}

// CreateDeleteHtmlResourceResponse creates a response to parse from DeleteHtmlResource response
func CreateDeleteHtmlResourceResponse() (response *DeleteHtmlResourceResponse) {
	response = &DeleteHtmlResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
