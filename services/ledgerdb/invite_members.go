package ledgerdb

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

// InviteMembers invokes the ledgerdb.InviteMembers API synchronously
// api document: https://help.aliyun.com/api/ledgerdb/invitemembers.html
func (client *Client) InviteMembers(request *InviteMembersRequest) (response *InviteMembersResponse, err error) {
	response = CreateInviteMembersResponse()
	err = client.DoAction(request, response)
	return
}

// InviteMembersWithChan invokes the ledgerdb.InviteMembers API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/invitemembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InviteMembersWithChan(request *InviteMembersRequest) (<-chan *InviteMembersResponse, <-chan error) {
	responseChan := make(chan *InviteMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InviteMembers(request)
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

// InviteMembersWithCallback invokes the ledgerdb.InviteMembers API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/invitemembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InviteMembersWithCallback(request *InviteMembersRequest, callback func(response *InviteMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InviteMembersResponse
		var err error
		defer close(result)
		response, err = client.InviteMembers(request)
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

// InviteMembersRequest is the request struct for api InviteMembers
type InviteMembersRequest struct {
	*requests.RpcRequest
	AliUids  string `position:"Body" name:"AliUids"`
	LedgerId string `position:"Body" name:"LedgerId"`
}

// InviteMembersResponse is the response struct for api InviteMembers
type InviteMembersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInviteMembersRequest creates a request to invoke InviteMembers API
func CreateInviteMembersRequest() (request *InviteMembersRequest) {
	request = &InviteMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ledgerdb", "2019-11-22", "InviteMembers", "ledgerdb", "openAPI")
	return
}

// CreateInviteMembersResponse creates a response to parse from InviteMembers response
func CreateInviteMembersResponse() (response *InviteMembersResponse) {
	response = &InviteMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}