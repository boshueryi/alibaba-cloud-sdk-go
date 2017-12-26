package slb

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

func (client *Client) DescribeCACertificates(request *DescribeCACertificatesRequest) (response *DescribeCACertificatesResponse, err error) {
	response = CreateDescribeCACertificatesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCACertificatesWithChan(request *DescribeCACertificatesRequest) (<-chan *DescribeCACertificatesResponse, <-chan error) {
	responseChan := make(chan *DescribeCACertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCACertificates(request)
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

func (client *Client) DescribeCACertificatesWithCallback(request *DescribeCACertificatesRequest, callback func(response *DescribeCACertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCACertificatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCACertificates(request)
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

type DescribeCACertificatesRequest struct {
	*requests.RpcRequest
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	CACertificateId      string `position:"Query" name:"CACertificateId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeCACertificatesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	CACertificates struct {
		CACertificate []struct {
			RegionId          string           `json:"RegionId" xml:"RegionId"`
			CACertificateId   string           `json:"CACertificateId" xml:"CACertificateId"`
			CACertificateName string           `json:"CACertificateName" xml:"CACertificateName"`
			Fingerprint       string           `json:"Fingerprint" xml:"Fingerprint"`
			ResourceGroupId   string           `json:"ResourceGroupId" xml:"ResourceGroupId"`
			CreateTime        string           `json:"CreateTime" xml:"CreateTime"`
			CreateTimeStamp   requests.Integer `json:"CreateTimeStamp" xml:"CreateTimeStamp"`
		} `json:"CACertificate" xml:"CACertificate"`
	} `json:"CACertificates" xml:"CACertificates"`
}

func CreateDescribeCACertificatesRequest() (request *DescribeCACertificatesRequest) {
	request = &DescribeCACertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeCACertificates", "", "")
	return
}

func CreateDescribeCACertificatesResponse() (response *DescribeCACertificatesResponse) {
	response = &DescribeCACertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
