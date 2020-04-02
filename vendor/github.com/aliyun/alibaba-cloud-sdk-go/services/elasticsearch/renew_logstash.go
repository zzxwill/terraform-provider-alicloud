package elasticsearch

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

// RenewLogstash invokes the elasticsearch.RenewLogstash API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/renewlogstash.html
func (client *Client) RenewLogstash(request *RenewLogstashRequest) (response *RenewLogstashResponse, err error) {
	response = CreateRenewLogstashResponse()
	err = client.DoAction(request, response)
	return
}

// RenewLogstashWithChan invokes the elasticsearch.RenewLogstash API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/renewlogstash.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewLogstashWithChan(request *RenewLogstashRequest) (<-chan *RenewLogstashResponse, <-chan error) {
	responseChan := make(chan *RenewLogstashResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewLogstash(request)
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

// RenewLogstashWithCallback invokes the elasticsearch.RenewLogstash API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/renewlogstash.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewLogstashWithCallback(request *RenewLogstashRequest, callback func(response *RenewLogstashResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewLogstashResponse
		var err error
		defer close(result)
		response, err = client.RenewLogstash(request)
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

// RenewLogstashRequest is the request struct for api RenewLogstash
type RenewLogstashRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// RenewLogstashResponse is the response struct for api RenewLogstash
type RenewLogstashResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateRenewLogstashRequest creates a request to invoke RenewLogstash API
func CreateRenewLogstashRequest() (request *RenewLogstashRequest) {
	request = &RenewLogstashRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "RenewLogstash", "/openapi/logstashes/[InstanceId]/actions/renew", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenewLogstashResponse creates a response to parse from RenewLogstash response
func CreateRenewLogstashResponse() (response *RenewLogstashResponse) {
	response = &RenewLogstashResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
