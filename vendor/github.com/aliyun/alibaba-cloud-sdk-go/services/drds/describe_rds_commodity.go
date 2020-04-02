package drds

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

// DescribeRdsCommodity invokes the drds.DescribeRdsCommodity API synchronously
// api document: https://help.aliyun.com/api/drds/describerdscommodity.html
func (client *Client) DescribeRdsCommodity(request *DescribeRdsCommodityRequest) (response *DescribeRdsCommodityResponse, err error) {
	response = CreateDescribeRdsCommodityResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRdsCommodityWithChan invokes the drds.DescribeRdsCommodity API asynchronously
// api document: https://help.aliyun.com/api/drds/describerdscommodity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRdsCommodityWithChan(request *DescribeRdsCommodityRequest) (<-chan *DescribeRdsCommodityResponse, <-chan error) {
	responseChan := make(chan *DescribeRdsCommodityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRdsCommodity(request)
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

// DescribeRdsCommodityWithCallback invokes the drds.DescribeRdsCommodity API asynchronously
// api document: https://help.aliyun.com/api/drds/describerdscommodity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRdsCommodityWithCallback(request *DescribeRdsCommodityRequest, callback func(response *DescribeRdsCommodityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRdsCommodityResponse
		var err error
		defer close(result)
		response, err = client.DescribeRdsCommodity(request)
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

// DescribeRdsCommodityRequest is the request struct for api DescribeRdsCommodity
type DescribeRdsCommodityRequest struct {
	*requests.RpcRequest
	CommodityCode  string `position:"Query" name:"CommodityCode"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	OrderType      string `position:"Query" name:"OrderType"`
}

// DescribeRdsCommodityResponse is the response struct for api DescribeRdsCommodity
type DescribeRdsCommodityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateDescribeRdsCommodityRequest creates a request to invoke DescribeRdsCommodity API
func CreateDescribeRdsCommodityRequest() (request *DescribeRdsCommodityRequest) {
	request = &DescribeRdsCommodityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeRdsCommodity", "Drds", "openAPI")
	return
}

// CreateDescribeRdsCommodityResponse creates a response to parse from DescribeRdsCommodity response
func CreateDescribeRdsCommodityResponse() (response *DescribeRdsCommodityResponse) {
	response = &DescribeRdsCommodityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
