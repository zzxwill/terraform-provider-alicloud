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

// CreateFlowForWeb invokes the emr.CreateFlowForWeb API synchronously
// api document: https://help.aliyun.com/api/emr/createflowforweb.html
func (client *Client) CreateFlowForWeb(request *CreateFlowForWebRequest) (response *CreateFlowForWebResponse, err error) {
	response = CreateCreateFlowForWebResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFlowForWebWithChan invokes the emr.CreateFlowForWeb API asynchronously
// api document: https://help.aliyun.com/api/emr/createflowforweb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFlowForWebWithChan(request *CreateFlowForWebRequest) (<-chan *CreateFlowForWebResponse, <-chan error) {
	responseChan := make(chan *CreateFlowForWebResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFlowForWeb(request)
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

// CreateFlowForWebWithCallback invokes the emr.CreateFlowForWeb API asynchronously
// api document: https://help.aliyun.com/api/emr/createflowforweb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFlowForWebWithCallback(request *CreateFlowForWebRequest, callback func(response *CreateFlowForWebResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFlowForWebResponse
		var err error
		defer close(result)
		response, err = client.CreateFlowForWeb(request)
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

// CreateFlowForWebRequest is the request struct for api CreateFlowForWeb
type CreateFlowForWebRequest struct {
	*requests.RpcRequest
	CronExpr                string           `position:"Query" name:"CronExpr"`
	Description             string           `position:"Query" name:"Description"`
	AlertUserGroupBizId     string           `position:"Query" name:"AlertUserGroupBizId"`
	HostName                string           `position:"Query" name:"HostName"`
	CreateCluster           requests.Boolean `position:"Query" name:"CreateCluster"`
	EndSchedule             requests.Integer `position:"Query" name:"EndSchedule"`
	AlertConf               string           `position:"Query" name:"AlertConf"`
	ProjectId               string           `position:"Query" name:"ProjectId"`
	ParentFlowList          string           `position:"Query" name:"ParentFlowList"`
	AlertDingDingGroupBizId string           `position:"Query" name:"AlertDingDingGroupBizId"`
	StartSchedule           requests.Integer `position:"Query" name:"StartSchedule"`
	ClusterId               string           `position:"Query" name:"ClusterId"`
	Graph                   string           `position:"Query" name:"Graph"`
	Name                    string           `position:"Query" name:"Name"`
	ParentCategory          string           `position:"Query" name:"ParentCategory"`
}

// CreateFlowForWebResponse is the response struct for api CreateFlowForWeb
type CreateFlowForWebResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateCreateFlowForWebRequest creates a request to invoke CreateFlowForWeb API
func CreateCreateFlowForWebRequest() (request *CreateFlowForWebRequest) {
	request = &CreateFlowForWebRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowForWeb", "emr", "openAPI")
	return
}

// CreateCreateFlowForWebResponse creates a response to parse from CreateFlowForWeb response
func CreateCreateFlowForWebResponse() (response *CreateFlowForWebResponse) {
	response = &CreateFlowForWebResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
