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

func (client *Client) CreateMigrateTask(request *CreateMigrateTaskRequest) (response *CreateMigrateTaskResponse, err error) {
	response = CreateCreateMigrateTaskResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateMigrateTaskWithChan(request *CreateMigrateTaskRequest) (<-chan *CreateMigrateTaskResponse, <-chan error) {
	responseChan := make(chan *CreateMigrateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMigrateTask(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) CreateMigrateTaskWithCallback(request *CreateMigrateTaskRequest, callback func(response *CreateMigrateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMigrateTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateMigrateTask(request)
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

type CreateMigrateTaskRequest struct {
	*requests.RpcRequest
	TaskType             string `position:"Query" name:"TaskType"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBName               string `position:"Query" name:"DBName"`
	OSSUrls              string `position:"Query" name:"OSSUrls"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	IsOnlineDB           string `position:"Query" name:"IsOnlineDB"`
}

type CreateMigrateTaskResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceId   string `json:"DBInstanceId" xml:"DBInstanceId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
	DBName         string `json:"DBName" xml:"DBName"`
	MigrateIaskId  string `json:"MigrateIaskId" xml:"MigrateIaskId"`
	TaskType       string `json:"TaskType" xml:"TaskType"`
}

func CreateCreateMigrateTaskRequest() (request *CreateMigrateTaskRequest) {
	request = &CreateMigrateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateMigrateTask", "", "")
	return
}

func CreateCreateMigrateTaskResponse() (response *CreateMigrateTaskResponse) {
	response = &CreateMigrateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
