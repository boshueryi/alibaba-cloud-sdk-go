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

func (client *Client) GetQuota(request *GetQuotaRequest) (response *GetQuotaResponse, err error) {
	response = CreateGetQuotaResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetQuotaWithChan(request *GetQuotaRequest) (<-chan *GetQuotaResponse, <-chan error) {
	responseChan := make(chan *GetQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQuota(request)
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

func (client *Client) GetQuotaWithCallback(request *GetQuotaRequest, callback func(response *GetQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQuotaResponse
		var err error
		defer close(result)
		response, err = client.GetQuota(request)
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

type GetQuotaRequest struct {
	*requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
}

type GetQuotaResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Quota     struct {
		TotalQuota  requests.Integer `json:"TotalQuota" xml:"TotalQuota"`
		FacesCount  requests.Integer `json:"FacesCount" xml:"FacesCount"`
		PhotosCount requests.Integer `json:"PhotosCount" xml:"PhotosCount"`
		UsedQuota   requests.Integer `json:"UsedQuota" xml:"UsedQuota"`
		VideosCount requests.Integer `json:"VideosCount" xml:"VideosCount"`
	} `json:"Quota" xml:"Quota"`
}

func CreateGetQuotaRequest() (request *GetQuotaRequest) {
	request = &GetQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetQuota", "", "")
	return
}

func CreateGetQuotaResponse() (response *GetQuotaResponse) {
	response = &GetQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
