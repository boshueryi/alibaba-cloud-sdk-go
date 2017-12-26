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

func (client *Client) CreateRouterInterface(request *CreateRouterInterfaceRequest) (response *CreateRouterInterfaceResponse, err error) {
	response = CreateCreateRouterInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateRouterInterfaceWithChan(request *CreateRouterInterfaceRequest) (<-chan *CreateRouterInterfaceResponse, <-chan error) {
	responseChan := make(chan *CreateRouterInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRouterInterface(request)
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

func (client *Client) CreateRouterInterfaceWithCallback(request *CreateRouterInterfaceRequest, callback func(response *CreateRouterInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRouterInterfaceResponse
		var err error
		defer close(result)
		response, err = client.CreateRouterInterface(request)
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

type CreateRouterInterfaceRequest struct {
	*requests.RpcRequest
	Spec                     string `position:"Query" name:"Spec"`
	OppositeRegionId         string `position:"Query" name:"OppositeRegionId"`
	HealthCheckTargetIp      string `position:"Query" name:"HealthCheckTargetIp"`
	ClientToken              string `position:"Query" name:"ClientToken"`
	OppositeInterfaceOwnerId string `position:"Query" name:"OppositeInterfaceOwnerId"`
	OppositeRouterId         string `position:"Query" name:"OppositeRouterId"`
	UserCidr                 string `position:"Query" name:"UserCidr"`
	AccessPointId            string `position:"Query" name:"AccessPointId"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
	HealthCheckSourceIp      string `position:"Query" name:"HealthCheckSourceIp"`
	OppositeRouterType       string `position:"Query" name:"OppositeRouterType"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	Description              string `position:"Query" name:"Description"`
	Name                     string `position:"Query" name:"Name"`
	RouterId                 string `position:"Query" name:"RouterId"`
	OppositeInterfaceId      string `position:"Query" name:"OppositeInterfaceId"`
	Role                     string `position:"Query" name:"Role"`
	RouterType               string `position:"Query" name:"RouterType"`
	ResourceOwnerId          string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OppositeAccessPointId    string `position:"Query" name:"OppositeAccessPointId"`
}

type CreateRouterInterfaceResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	RouterInterfaceId string `json:"RouterInterfaceId" xml:"RouterInterfaceId"`
}

func CreateCreateRouterInterfaceRequest() (request *CreateRouterInterfaceRequest) {
	request = &CreateRouterInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateRouterInterface", "", "")
	return
}

func CreateCreateRouterInterfaceResponse() (response *CreateRouterInterfaceResponse) {
	response = &CreateRouterInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
