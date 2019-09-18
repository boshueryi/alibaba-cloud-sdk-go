package imm

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

// CreateVideoCompressTask invokes the imm.CreateVideoCompressTask API synchronously
// api document: https://help.aliyun.com/api/imm/createvideocompresstask.html
func (client *Client) CreateVideoCompressTask(request *CreateVideoCompressTaskRequest) (response *CreateVideoCompressTaskResponse, err error) {
	response = CreateCreateVideoCompressTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVideoCompressTaskWithChan invokes the imm.CreateVideoCompressTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createvideocompresstask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVideoCompressTaskWithChan(request *CreateVideoCompressTaskRequest) (<-chan *CreateVideoCompressTaskResponse, <-chan error) {
	responseChan := make(chan *CreateVideoCompressTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVideoCompressTask(request)
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

// CreateVideoCompressTaskWithCallback invokes the imm.CreateVideoCompressTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createvideocompresstask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVideoCompressTaskWithCallback(request *CreateVideoCompressTaskRequest, callback func(response *CreateVideoCompressTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVideoCompressTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateVideoCompressTask(request)
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

// CreateVideoCompressTaskRequest is the request struct for api CreateVideoCompressTask
type CreateVideoCompressTaskRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	TargetList      string `position:"Query" name:"TargetList"`
	VideoUri        string `position:"Query" name:"VideoUri"`
}

// CreateVideoCompressTaskResponse is the response struct for api CreateVideoCompressTask
type CreateVideoCompressTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateVideoCompressTaskRequest creates a request to invoke CreateVideoCompressTask API
func CreateCreateVideoCompressTaskRequest() (request *CreateVideoCompressTaskRequest) {
	request = &CreateVideoCompressTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateVideoCompressTask", "imm", "openAPI")
	return
}

// CreateCreateVideoCompressTaskResponse creates a response to parse from CreateVideoCompressTask response
func CreateCreateVideoCompressTaskResponse() (response *CreateVideoCompressTaskResponse) {
	response = &CreateVideoCompressTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
