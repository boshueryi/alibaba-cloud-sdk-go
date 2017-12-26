package rds

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

func (client *Client) CreateDiagnosticReport(request *CreateDiagnosticReportRequest) (response *CreateDiagnosticReportResponse, err error) {
	response = CreateCreateDiagnosticReportResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateDiagnosticReportWithChan(request *CreateDiagnosticReportRequest) (<-chan *CreateDiagnosticReportResponse, <-chan error) {
	responseChan := make(chan *CreateDiagnosticReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDiagnosticReport(request)
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

func (client *Client) CreateDiagnosticReportWithCallback(request *CreateDiagnosticReportRequest, callback func(response *CreateDiagnosticReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDiagnosticReportResponse
		var err error
		defer close(result)
		response, err = client.CreateDiagnosticReport(request)
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

type CreateDiagnosticReportRequest struct {
	*requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

type CreateDiagnosticReportResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ReportId  string `json:"ReportId" xml:"ReportId"`
}

func CreateCreateDiagnosticReportRequest() (request *CreateDiagnosticReportRequest) {
	request = &CreateDiagnosticReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDiagnosticReport", "", "")
	return
}

func CreateCreateDiagnosticReportResponse() (response *CreateDiagnosticReportResponse) {
	response = &CreateDiagnosticReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
