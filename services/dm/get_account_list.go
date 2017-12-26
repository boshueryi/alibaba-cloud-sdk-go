package dm

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

func (client *Client) GetAccountList(request *GetAccountListRequest) (response *GetAccountListResponse, err error) {
	response = CreateGetAccountListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetAccountListWithChan(request *GetAccountListRequest) (<-chan *GetAccountListResponse, <-chan error) {
	responseChan := make(chan *GetAccountListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountList(request)
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

func (client *Client) GetAccountListWithCallback(request *GetAccountListRequest, callback func(response *GetAccountListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccountListResponse
		var err error
		defer close(result)
		response, err = client.GetAccountList(request)
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

type GetAccountListRequest struct {
	*requests.RpcRequest
	Total                string `position:"Query" name:"Total"`
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	OffsetCreateTimeDesc string `position:"Query" name:"OffsetCreateTimeDesc"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Offset               string `position:"Query" name:"Offset"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	OffsetCreateTime     string `position:"Query" name:"OffsetCreateTime"`
}

type GetAccountListResponse struct {
	*responses.BaseResponse
	RequestId string           `json:"RequestId" xml:"RequestId"`
	Total     requests.Integer `json:"Total" xml:"Total"`
	PageNo    requests.Integer `json:"PageNo" xml:"PageNo"`
	PageSize  requests.Integer `json:"PageSize" xml:"PageSize"`
	Data      struct {
		AccountNotificationInfo []struct {
			Region     string `json:"Region" xml:"Region"`
			UpdateTime string `json:"UpdateTime" xml:"UpdateTime"`
			Status     string `json:"Status" xml:"Status"`
		} `json:"accountNotificationInfo" xml:"accountNotificationInfo"`
	} `json:"data" xml:"data"`
}

func CreateGetAccountListRequest() (request *GetAccountListRequest) {
	request = &GetAccountListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "GetAccountList", "", "")
	return
}

func CreateGetAccountListResponse() (response *GetAccountListResponse) {
	response = &GetAccountListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
