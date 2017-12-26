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

func (client *Client) UpdateDomainTrackName(request *UpdateDomainTrackNameRequest) (response *UpdateDomainTrackNameResponse, err error) {
	response = CreateUpdateDomainTrackNameResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateDomainTrackNameWithChan(request *UpdateDomainTrackNameRequest) (<-chan *UpdateDomainTrackNameResponse, <-chan error) {
	responseChan := make(chan *UpdateDomainTrackNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDomainTrackName(request)
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

func (client *Client) UpdateDomainTrackNameWithCallback(request *UpdateDomainTrackNameRequest, callback func(response *UpdateDomainTrackNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDomainTrackNameResponse
		var err error
		defer close(result)
		response, err = client.UpdateDomainTrackName(request)
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

type UpdateDomainTrackNameRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CnameTrackRecord     string `position:"Query" name:"CnameTrackRecord"`
	DomainId             string `position:"Query" name:"DomainId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type UpdateDomainTrackNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateUpdateDomainTrackNameRequest() (request *UpdateDomainTrackNameRequest) {
	request = &UpdateDomainTrackNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "UpdateDomainTrackName", "", "")
	return
}

func CreateUpdateDomainTrackNameResponse() (response *UpdateDomainTrackNameResponse) {
	response = &UpdateDomainTrackNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
