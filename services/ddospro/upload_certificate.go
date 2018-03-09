package ddospro

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

// UploadCertificate invokes the ddospro.UploadCertificate API synchronously
// api document: https://help.aliyun.com/api/ddospro/uploadcertificate.html
func (client *Client) UploadCertificate(request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
	response = CreateUploadCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// UploadCertificateWithChan invokes the ddospro.UploadCertificate API asynchronously
// api document: https://help.aliyun.com/api/ddospro/uploadcertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadCertificateWithChan(request *UploadCertificateRequest) (<-chan *UploadCertificateResponse, <-chan error) {
	responseChan := make(chan *UploadCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadCertificate(request)
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

// UploadCertificateWithCallback invokes the ddospro.UploadCertificate API asynchronously
// api document: https://help.aliyun.com/api/ddospro/uploadcertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadCertificateWithCallback(request *UploadCertificateRequest, callback func(response *UploadCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadCertificateResponse
		var err error
		defer close(result)
		response, err = client.UploadCertificate(request)
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

// UploadCertificateRequest is the request struct for api UploadCertificate
type UploadCertificateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Domain          string           `position:"Query" name:"Domain"`
	Cert            string           `position:"Query" name:"Cert"`
	Key             string           `position:"Query" name:"Key"`
}

// UploadCertificateResponse is the response struct for api UploadCertificate
type UploadCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUploadCertificateRequest creates a request to invoke UploadCertificate API
func CreateUploadCertificateRequest(request *UploadCertificateRequest) {
	request = &UploadCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "UploadCertificate", "", "")
	return
}

// CreateUploadCertificateResponse creates a response to parse from UploadCertificate response
func CreateUploadCertificateResponse() (response *UploadCertificateResponse) {
	response = &UploadCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
