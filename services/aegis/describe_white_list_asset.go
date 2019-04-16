package aegis

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

// DescribeWhiteListAsset invokes the aegis.DescribeWhiteListAsset API synchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelistasset.html
func (client *Client) DescribeWhiteListAsset(request *DescribeWhiteListAssetRequest) (response *DescribeWhiteListAssetResponse, err error) {
	response = CreateDescribeWhiteListAssetResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWhiteListAssetWithChan invokes the aegis.DescribeWhiteListAsset API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelistasset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListAssetWithChan(request *DescribeWhiteListAssetRequest) (<-chan *DescribeWhiteListAssetResponse, <-chan error) {
	responseChan := make(chan *DescribeWhiteListAssetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWhiteListAsset(request)
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

// DescribeWhiteListAssetWithCallback invokes the aegis.DescribeWhiteListAsset API asynchronously
// api document: https://help.aliyun.com/api/aegis/describewhitelistasset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWhiteListAssetWithCallback(request *DescribeWhiteListAssetRequest, callback func(response *DescribeWhiteListAssetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWhiteListAssetResponse
		var err error
		defer close(result)
		response, err = client.DescribeWhiteListAsset(request)
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

// DescribeWhiteListAssetRequest is the request struct for api DescribeWhiteListAsset
type DescribeWhiteListAssetRequest struct {
	*requests.RpcRequest
	SourceIp   string           `position:"Query" name:"SourceIp"`
	LastMaxId  requests.Integer `position:"Query" name:"LastMaxId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	StrategyId requests.Integer `position:"Query" name:"StrategyId"`
	Lang       string           `position:"Query" name:"Lang"`
	Type       requests.Integer `position:"Query" name:"Type"`
}

// DescribeWhiteListAssetResponse is the response struct for api DescribeWhiteListAsset
type DescribeWhiteListAssetResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Assets    []Asset `json:"Assets" xml:"Assets"`
}

// CreateDescribeWhiteListAssetRequest creates a request to invoke DescribeWhiteListAsset API
func CreateDescribeWhiteListAssetRequest() (request *DescribeWhiteListAssetRequest) {
	request = &DescribeWhiteListAssetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeWhiteListAsset", "vipaegis", "openAPI")
	return
}

// CreateDescribeWhiteListAssetResponse creates a response to parse from DescribeWhiteListAsset response
func CreateDescribeWhiteListAssetResponse() (response *DescribeWhiteListAssetResponse) {
	response = &DescribeWhiteListAssetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}