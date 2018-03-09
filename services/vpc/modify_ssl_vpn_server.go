package vpc

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

// ModifySslVpnServer invokes the vpc.ModifySslVpnServer API synchronously
// api document: https://help.aliyun.com/api/vpc/modifysslvpnserver.html
func (client *Client) ModifySslVpnServer(request *ModifySslVpnServerRequest) (response *ModifySslVpnServerResponse, err error) {
	response = CreateModifySslVpnServerResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySslVpnServerWithChan invokes the vpc.ModifySslVpnServer API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifysslvpnserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySslVpnServerWithChan(request *ModifySslVpnServerRequest) (<-chan *ModifySslVpnServerResponse, <-chan error) {
	responseChan := make(chan *ModifySslVpnServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySslVpnServer(request)
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

// ModifySslVpnServerWithCallback invokes the vpc.ModifySslVpnServer API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifysslvpnserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySslVpnServerWithCallback(request *ModifySslVpnServerRequest, callback func(response *ModifySslVpnServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySslVpnServerResponse
		var err error
		defer close(result)
		response, err = client.ModifySslVpnServer(request)
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

// ModifySslVpnServerRequest is the request struct for api ModifySslVpnServer
type ModifySslVpnServerRequest struct {
	*requests.RpcRequest
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	SslVpnServerId       string           `position:"Query" name:"SslVpnServerId"`
	Name                 string           `position:"Query" name:"Name"`
	ClientIpPool         string           `position:"Query" name:"ClientIpPool"`
	LocalSubnet          string           `position:"Query" name:"LocalSubnet"`
	Proto                string           `position:"Query" name:"Proto"`
	Cipher               string           `position:"Query" name:"Cipher"`
	Port                 requests.Integer `position:"Query" name:"Port"`
	Compress             requests.Boolean `position:"Query" name:"Compress"`
}

// ModifySslVpnServerResponse is the response struct for api ModifySslVpnServer
type ModifySslVpnServerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	RegionId       string `json:"RegionId" xml:"RegionId"`
	SslVpnServerId string `json:"SslVpnServerId" xml:"SslVpnServerId"`
	VpnGatewayId   string `json:"VpnGatewayId" xml:"VpnGatewayId"`
	Name           string `json:"Name" xml:"Name"`
	LocalSubnet    string `json:"LocalSubnet" xml:"LocalSubnet"`
	ClientIpPool   string `json:"ClientIpPool" xml:"ClientIpPool"`
	CreateTime     int    `json:"CreateTime" xml:"CreateTime"`
	Cipher         string `json:"Cipher" xml:"Cipher"`
	Proto          string `json:"Proto" xml:"Proto"`
	Port           int    `json:"Port" xml:"Port"`
	Compress       bool   `json:"Compress" xml:"Compress"`
	Connections    int    `json:"Connections" xml:"Connections"`
	MaxConnections int    `json:"MaxConnections" xml:"MaxConnections"`
	InternetIp     string `json:"InternetIp" xml:"InternetIp"`
}

// CreateModifySslVpnServerRequest creates a request to invoke ModifySslVpnServer API
func CreateModifySslVpnServerRequest(request *ModifySslVpnServerRequest) {
	request = &ModifySslVpnServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifySslVpnServer", "vpc", "openAPI")
	return
}

// CreateModifySslVpnServerResponse creates a response to parse from ModifySslVpnServer response
func CreateModifySslVpnServerResponse() (response *ModifySslVpnServerResponse) {
	response = &ModifySslVpnServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
