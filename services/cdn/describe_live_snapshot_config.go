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

func (client *Client) DescribeLiveSnapshotConfig(request *DescribeLiveSnapshotConfigRequest) (response *DescribeLiveSnapshotConfigResponse, err error) {
	response = CreateDescribeLiveSnapshotConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveSnapshotConfigWithChan(request *DescribeLiveSnapshotConfigRequest) (<-chan *DescribeLiveSnapshotConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveSnapshotConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveSnapshotConfig(request)
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

func (client *Client) DescribeLiveSnapshotConfigWithCallback(request *DescribeLiveSnapshotConfigRequest, callback func(response *DescribeLiveSnapshotConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveSnapshotConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveSnapshotConfig(request)
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

type DescribeLiveSnapshotConfigRequest struct {
	*requests.RpcRequest
	PageSize      string `position:"Query" name:"PageSize"`
	StreamName    string `position:"Query" name:"StreamName"`
	DomainName    string `position:"Query" name:"DomainName"`
	Order         string `position:"Query" name:"Order"`
	AppName       string `position:"Query" name:"AppName"`
	PageNum       string `position:"Query" name:"PageNum"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeLiveSnapshotConfigResponse struct {
	*responses.BaseResponse
	RequestId                    string           `json:"RequestId" xml:"RequestId"`
	PageNum                      requests.Integer `json:"PageNum" xml:"PageNum"`
	PageSize                     requests.Integer `json:"PageSize" xml:"PageSize"`
	Order                        string           `json:"Order" xml:"Order"`
	TotalNum                     requests.Integer `json:"TotalNum" xml:"TotalNum"`
	TotalPage                    requests.Integer `json:"TotalPage" xml:"TotalPage"`
	LiveStreamSnapshotConfigList struct {
		LiveStreamSnapshotConfig []struct {
			DomainName         string           `json:"DomainName" xml:"DomainName"`
			AppName            string           `json:"AppName" xml:"AppName"`
			TimeInterval       requests.Integer `json:"TimeInterval" xml:"TimeInterval"`
			OssEndpoint        string           `json:"OssEndpoint" xml:"OssEndpoint"`
			OssBucket          string           `json:"OssBucket" xml:"OssBucket"`
			OverwriteOssObject string           `json:"OverwriteOssObject" xml:"OverwriteOssObject"`
			SequenceOssObject  string           `json:"SequenceOssObject" xml:"SequenceOssObject"`
			CreateTime         string           `json:"CreateTime" xml:"CreateTime"`
		} `json:"LiveStreamSnapshotConfig" xml:"LiveStreamSnapshotConfig"`
	} `json:"LiveStreamSnapshotConfigList" xml:"LiveStreamSnapshotConfigList"`
}

func CreateDescribeLiveSnapshotConfigRequest() (request *DescribeLiveSnapshotConfigRequest) {
	request = &DescribeLiveSnapshotConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveSnapshotConfig", "", "")
	return
}

func CreateDescribeLiveSnapshotConfigResponse() (response *DescribeLiveSnapshotConfigResponse) {
	response = &DescribeLiveSnapshotConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
