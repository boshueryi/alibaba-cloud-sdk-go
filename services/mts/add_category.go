package mts

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

func (client *Client) AddCategory(request *AddCategoryRequest) (response *AddCategoryResponse, err error) {
	response = CreateAddCategoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddCategoryWithChan(request *AddCategoryRequest) (<-chan *AddCategoryResponse, <-chan error) {
	responseChan := make(chan *AddCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCategory(request)
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

func (client *Client) AddCategoryWithCallback(request *AddCategoryRequest, callback func(response *AddCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCategoryResponse
		var err error
		defer close(result)
		response, err = client.AddCategory(request)
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

type AddCategoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ParentId             string `position:"Query" name:"ParentId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	CateName             string `position:"Query" name:"CateName"`
}

type AddCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Category  struct {
		CateId   string `json:"CateId" xml:"CateId"`
		CateName string `json:"CateName" xml:"CateName"`
		ParentId string `json:"ParentId" xml:"ParentId"`
		Level    string `json:"Level" xml:"Level"`
	} `json:"Category" xml:"Category"`
}

func CreateAddCategoryRequest() (request *AddCategoryRequest) {
	request = &AddCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddCategory", "", "")
	return
}

func CreateAddCategoryResponse() (response *AddCategoryResponse) {
	response = &AddCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
