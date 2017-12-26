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

func (client *Client) ListTimeLinePhotos(request *ListTimeLinePhotosRequest) (response *ListTimeLinePhotosResponse, err error) {
	response = CreateListTimeLinePhotosResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListTimeLinePhotosWithChan(request *ListTimeLinePhotosRequest) (<-chan *ListTimeLinePhotosResponse, <-chan error) {
	responseChan := make(chan *ListTimeLinePhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTimeLinePhotos(request)
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

func (client *Client) ListTimeLinePhotosWithCallback(request *ListTimeLinePhotosRequest, callback func(response *ListTimeLinePhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTimeLinePhotosResponse
		var err error
		defer close(result)
		response, err = client.ListTimeLinePhotos(request)
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

type ListTimeLinePhotosRequest struct {
	*requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	FilterBy  string `position:"Query" name:"FilterBy"`
	StartTime string `position:"Query" name:"StartTime"`
	Order     string `position:"Query" name:"Order"`
	Page      string `position:"Query" name:"Page"`
	Direction string `position:"Query" name:"Direction"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Size      string `position:"Query" name:"Size"`
}

type ListTimeLinePhotosResponse struct {
	*responses.BaseResponse
	Code       string           `json:"Code" xml:"Code"`
	Message    string           `json:"Message" xml:"Message"`
	TotalCount requests.Integer `json:"TotalCount" xml:"TotalCount"`
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	Action     string           `json:"Action" xml:"Action"`
	Photos     []struct {
		Id              requests.Integer `json:"Id" xml:"Id"`
		Title           string           `json:"Title" xml:"Title"`
		FileId          string           `json:"FileId" xml:"FileId"`
		State           string           `json:"State" xml:"State"`
		Md5             string           `json:"Md5" xml:"Md5"`
		IsVideo         requests.Boolean `json:"IsVideo" xml:"IsVideo"`
		Remark          string           `json:"Remark" xml:"Remark"`
		Width           requests.Integer `json:"Width" xml:"Width"`
		Height          requests.Integer `json:"Height" xml:"Height"`
		Ctime           requests.Integer `json:"Ctime" xml:"Ctime"`
		Mtime           requests.Integer `json:"Mtime" xml:"Mtime"`
		TakenAt         requests.Integer `json:"TakenAt" xml:"TakenAt"`
		ShareExpireTime requests.Integer `json:"ShareExpireTime" xml:"ShareExpireTime"`
		Like            requests.Integer `json:"Like" xml:"Like"`
	} `json:"Photos" xml:"Photos"`
}

func CreateListTimeLinePhotosRequest() (request *ListTimeLinePhotosRequest) {
	request = &ListTimeLinePhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTimeLinePhotos", "", "")
	return
}

func CreateListTimeLinePhotosResponse() (response *ListTimeLinePhotosResponse) {
	response = &ListTimeLinePhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
