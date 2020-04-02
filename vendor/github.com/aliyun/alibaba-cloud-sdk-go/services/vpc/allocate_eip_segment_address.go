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

// AllocateEipSegmentAddress invokes the vpc.AllocateEipSegmentAddress API synchronously
// api document: https://help.aliyun.com/api/vpc/allocateeipsegmentaddress.html
func (client *Client) AllocateEipSegmentAddress(request *AllocateEipSegmentAddressRequest) (response *AllocateEipSegmentAddressResponse, err error) {
	response = CreateAllocateEipSegmentAddressResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateEipSegmentAddressWithChan invokes the vpc.AllocateEipSegmentAddress API asynchronously
// api document: https://help.aliyun.com/api/vpc/allocateeipsegmentaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateEipSegmentAddressWithChan(request *AllocateEipSegmentAddressRequest) (<-chan *AllocateEipSegmentAddressResponse, <-chan error) {
	responseChan := make(chan *AllocateEipSegmentAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateEipSegmentAddress(request)
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

// AllocateEipSegmentAddressWithCallback invokes the vpc.AllocateEipSegmentAddress API asynchronously
// api document: https://help.aliyun.com/api/vpc/allocateeipsegmentaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateEipSegmentAddressWithCallback(request *AllocateEipSegmentAddressRequest, callback func(response *AllocateEipSegmentAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateEipSegmentAddressResponse
		var err error
		defer close(result)
		response, err = client.AllocateEipSegmentAddress(request)
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

// AllocateEipSegmentAddressRequest is the request struct for api AllocateEipSegmentAddress
type AllocateEipSegmentAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	Netmode              string           `position:"Query" name:"Netmode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EipMask              string           `position:"Query" name:"EipMask"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
}

// AllocateEipSegmentAddressResponse is the response struct for api AllocateEipSegmentAddress
type AllocateEipSegmentAddressResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	EipSegmentInstanceId string `json:"EipSegmentInstanceId" xml:"EipSegmentInstanceId"`
	OrderId              int64  `json:"OrderId" xml:"OrderId"`
}

// CreateAllocateEipSegmentAddressRequest creates a request to invoke AllocateEipSegmentAddress API
func CreateAllocateEipSegmentAddressRequest() (request *AllocateEipSegmentAddressRequest) {
	request = &AllocateEipSegmentAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AllocateEipSegmentAddress", "Vpc", "openAPI")
	return
}

// CreateAllocateEipSegmentAddressResponse creates a response to parse from AllocateEipSegmentAddress response
func CreateAllocateEipSegmentAddressResponse() (response *AllocateEipSegmentAddressResponse) {
	response = &AllocateEipSegmentAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
