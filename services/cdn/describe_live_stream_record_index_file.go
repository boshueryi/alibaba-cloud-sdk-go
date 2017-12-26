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

func (client *Client) DescribeLiveStreamRecordIndexFile(request *DescribeLiveStreamRecordIndexFileRequest) (response *DescribeLiveStreamRecordIndexFileResponse, err error) {
	response = CreateDescribeLiveStreamRecordIndexFileResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamRecordIndexFileWithChan(request *DescribeLiveStreamRecordIndexFileRequest) (<-chan *DescribeLiveStreamRecordIndexFileResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRecordIndexFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRecordIndexFile(request)
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

func (client *Client) DescribeLiveStreamRecordIndexFileWithCallback(request *DescribeLiveStreamRecordIndexFileRequest, callback func(response *DescribeLiveStreamRecordIndexFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRecordIndexFileResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRecordIndexFile(request)
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

type DescribeLiveStreamRecordIndexFileRequest struct {
	*requests.RpcRequest
	StreamName    string `position:"Query" name:"StreamName"`
	DomainName    string `position:"Query" name:"DomainName"`
	AppName       string `position:"Query" name:"AppName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	RecordId      string `position:"Query" name:"RecordId"`
}

type DescribeLiveStreamRecordIndexFileResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	RecordIndexInfo struct {
		RecordId   string           `json:"RecordId" xml:"RecordId"`
		RecordUrl  string           `json:"RecordUrl" xml:"RecordUrl"`
		DomainName string           `json:"DomainName" xml:"DomainName"`
		AppName    string           `json:"AppName" xml:"AppName"`
		StreamName string           `json:"StreamName" xml:"StreamName"`
		OssObject  string           `json:"OssObject" xml:"OssObject"`
		StartTime  string           `json:"StartTime" xml:"StartTime"`
		EndTime    string           `json:"EndTime" xml:"EndTime"`
		Duration   requests.Float   `json:"Duration" xml:"Duration"`
		Height     requests.Integer `json:"Height" xml:"Height"`
		Width      requests.Integer `json:"Width" xml:"Width"`
		CreateTime string           `json:"CreateTime" xml:"CreateTime"`
	} `json:"RecordIndexInfo" xml:"RecordIndexInfo"`
}

func CreateDescribeLiveStreamRecordIndexFileRequest() (request *DescribeLiveStreamRecordIndexFileRequest) {
	request = &DescribeLiveStreamRecordIndexFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRecordIndexFile", "", "")
	return
}

func CreateDescribeLiveStreamRecordIndexFileResponse() (response *DescribeLiveStreamRecordIndexFileResponse) {
	response = &DescribeLiveStreamRecordIndexFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
