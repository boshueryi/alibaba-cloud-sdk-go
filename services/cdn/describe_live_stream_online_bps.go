package cdn

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

// DescribeLiveStreamOnlineBps invokes the cdn.DescribeLiveStreamOnlineBps API synchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamonlinebps.html
func (client *Client) DescribeLiveStreamOnlineBps(request *DescribeLiveStreamOnlineBpsRequest) (response *DescribeLiveStreamOnlineBpsResponse, err error) {
	response = CreateDescribeLiveStreamOnlineBpsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamOnlineBpsWithChan invokes the cdn.DescribeLiveStreamOnlineBps API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamonlinebps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamOnlineBpsWithChan(request *DescribeLiveStreamOnlineBpsRequest) (<-chan *DescribeLiveStreamOnlineBpsResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamOnlineBpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamOnlineBps(request)
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

// DescribeLiveStreamOnlineBpsWithCallback invokes the cdn.DescribeLiveStreamOnlineBps API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamonlinebps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamOnlineBpsWithCallback(request *DescribeLiveStreamOnlineBpsRequest, callback func(response *DescribeLiveStreamOnlineBpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamOnlineBpsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamOnlineBps(request)
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

// DescribeLiveStreamOnlineBpsRequest is the request struct for api DescribeLiveStreamOnlineBps
type DescribeLiveStreamOnlineBpsRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
}

// DescribeLiveStreamOnlineBpsResponse is the response struct for api DescribeLiveStreamOnlineBps
type DescribeLiveStreamOnlineBpsResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	TotalUserNumber          int                      `json:"TotalUserNumber" xml:"TotalUserNumber"`
	FlvBps                   float64                  `json:"FlvBps" xml:"FlvBps"`
	HlsBps                   float64                  `json:"HlsBps" xml:"HlsBps"`
	LiveStreamOnlineBpsInfos LiveStreamOnlineBpsInfos `json:"LiveStreamOnlineBpsInfos" xml:"LiveStreamOnlineBpsInfos"`
}

// CreateDescribeLiveStreamOnlineBpsRequest creates a request to invoke DescribeLiveStreamOnlineBps API
func CreateDescribeLiveStreamOnlineBpsRequest(request *DescribeLiveStreamOnlineBpsRequest) {
	request = &DescribeLiveStreamOnlineBpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamOnlineBps", "", "")
	return
}

// CreateDescribeLiveStreamOnlineBpsResponse creates a response to parse from DescribeLiveStreamOnlineBps response
func CreateDescribeLiveStreamOnlineBpsResponse() (response *DescribeLiveStreamOnlineBpsResponse) {
	response = &DescribeLiveStreamOnlineBpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
