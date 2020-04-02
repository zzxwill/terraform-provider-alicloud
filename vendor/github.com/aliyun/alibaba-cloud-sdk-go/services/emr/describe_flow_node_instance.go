package emr

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

// DescribeFlowNodeInstance invokes the emr.DescribeFlowNodeInstance API synchronously
// api document: https://help.aliyun.com/api/emr/describeflownodeinstance.html
func (client *Client) DescribeFlowNodeInstance(request *DescribeFlowNodeInstanceRequest) (response *DescribeFlowNodeInstanceResponse, err error) {
	response = CreateDescribeFlowNodeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowNodeInstanceWithChan invokes the emr.DescribeFlowNodeInstance API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflownodeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowNodeInstanceWithChan(request *DescribeFlowNodeInstanceRequest) (<-chan *DescribeFlowNodeInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowNodeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowNodeInstance(request)
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

// DescribeFlowNodeInstanceWithCallback invokes the emr.DescribeFlowNodeInstance API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflownodeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowNodeInstanceWithCallback(request *DescribeFlowNodeInstanceRequest, callback func(response *DescribeFlowNodeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowNodeInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowNodeInstance(request)
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

// DescribeFlowNodeInstanceRequest is the request struct for api DescribeFlowNodeInstance
type DescribeFlowNodeInstanceRequest struct {
	*requests.RpcRequest
	Id        string `position:"Query" name:"Id"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DescribeFlowNodeInstanceResponse is the response struct for api DescribeFlowNodeInstance
type DescribeFlowNodeInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Id               string `json:"Id" xml:"Id"`
	GmtCreate        int64  `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified      int64  `json:"GmtModified" xml:"GmtModified"`
	Type             string `json:"Type" xml:"Type"`
	Status           string `json:"Status" xml:"Status"`
	JobId            string `json:"JobId" xml:"JobId"`
	JobName          string `json:"JobName" xml:"JobName"`
	JobType          string `json:"JobType" xml:"JobType"`
	JobParams        string `json:"JobParams" xml:"JobParams"`
	FailAct          string `json:"FailAct" xml:"FailAct"`
	MaxRetry         string `json:"MaxRetry" xml:"MaxRetry"`
	RetryInterval    string `json:"RetryInterval" xml:"RetryInterval"`
	NodeName         string `json:"NodeName" xml:"NodeName"`
	FlowId           string `json:"FlowId" xml:"FlowId"`
	FlowInstanceId   string `json:"FlowInstanceId" xml:"FlowInstanceId"`
	ClusterId        string `json:"ClusterId" xml:"ClusterId"`
	HostName         string `json:"HostName" xml:"HostName"`
	ProjectId        string `json:"ProjectId" xml:"ProjectId"`
	Pending          bool   `json:"Pending" xml:"Pending"`
	StartTime        int64  `json:"StartTime" xml:"StartTime"`
	EndTime          int64  `json:"EndTime" xml:"EndTime"`
	Duration         int64  `json:"Duration" xml:"Duration"`
	Retries          int    `json:"Retries" xml:"Retries"`
	ExternalId       string `json:"ExternalId" xml:"ExternalId"`
	ExternalSubId    string `json:"ExternalSubId" xml:"ExternalSubId"`
	ExternalChildIds string `json:"ExternalChildIds" xml:"ExternalChildIds"`
	ExternalStatus   string `json:"ExternalStatus" xml:"ExternalStatus"`
	ExternalInfo     string `json:"ExternalInfo" xml:"ExternalInfo"`
	ParamConf        string `json:"ParamConf" xml:"ParamConf"`
	EnvConf          string `json:"EnvConf" xml:"EnvConf"`
	RunConf          string `json:"RunConf" xml:"RunConf"`
	Adhoc            bool   `json:"Adhoc" xml:"Adhoc"`
	MonitorConf      string `json:"MonitorConf" xml:"MonitorConf"`
	Mode             string `json:"Mode" xml:"Mode"`
	ClusterName      string `json:"ClusterName" xml:"ClusterName"`
}

// CreateDescribeFlowNodeInstanceRequest creates a request to invoke DescribeFlowNodeInstance API
func CreateDescribeFlowNodeInstanceRequest() (request *DescribeFlowNodeInstanceRequest) {
	request = &DescribeFlowNodeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowNodeInstance", "emr", "openAPI")
	return
}

// CreateDescribeFlowNodeInstanceResponse creates a response to parse from DescribeFlowNodeInstance response
func CreateDescribeFlowNodeInstanceResponse() (response *DescribeFlowNodeInstanceResponse) {
	response = &DescribeFlowNodeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
