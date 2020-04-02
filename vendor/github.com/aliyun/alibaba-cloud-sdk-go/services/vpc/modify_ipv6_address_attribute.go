package vpc

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

// ModifyIpv6AddressAttribute invokes the vpc.ModifyIpv6AddressAttribute API synchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6addressattribute.html
func (client *Client) ModifyIpv6AddressAttribute(request *ModifyIpv6AddressAttributeRequest) (response *ModifyIpv6AddressAttributeResponse, err error) {
	response = CreateModifyIpv6AddressAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIpv6AddressAttributeWithChan invokes the vpc.ModifyIpv6AddressAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6addressattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIpv6AddressAttributeWithChan(request *ModifyIpv6AddressAttributeRequest) (<-chan *ModifyIpv6AddressAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyIpv6AddressAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIpv6AddressAttribute(request)
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

// ModifyIpv6AddressAttributeWithCallback invokes the vpc.ModifyIpv6AddressAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6addressattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIpv6AddressAttributeWithCallback(request *ModifyIpv6AddressAttributeRequest, callback func(response *ModifyIpv6AddressAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIpv6AddressAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyIpv6AddressAttribute(request)
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

// ModifyIpv6AddressAttributeRequest is the request struct for api ModifyIpv6AddressAttribute
type ModifyIpv6AddressAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ipv6AddressId        string           `position:"Query" name:"Ipv6AddressId"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyIpv6AddressAttributeResponse is the response struct for api ModifyIpv6AddressAttribute
type ModifyIpv6AddressAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyIpv6AddressAttributeRequest creates a request to invoke ModifyIpv6AddressAttribute API
func CreateModifyIpv6AddressAttributeRequest() (request *ModifyIpv6AddressAttributeRequest) {
	request = &ModifyIpv6AddressAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyIpv6AddressAttribute", "Vpc", "openAPI")
	return
}

// CreateModifyIpv6AddressAttributeResponse creates a response to parse from ModifyIpv6AddressAttribute response
func CreateModifyIpv6AddressAttributeResponse() (response *ModifyIpv6AddressAttributeResponse) {
	response = &ModifyIpv6AddressAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
