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

// CreateLoadBalancer invokes the slb.CreateLoadBalancer API synchronously
// api document: https://help.aliyun.com/api/slb/createloadbalancer.html
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = CreateCreateLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerWithChan invokes the slb.CreateLoadBalancer API asynchronously
// api document: https://help.aliyun.com/api/slb/createloadbalancer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLoadBalancerWithChan(request *CreateLoadBalancerRequest) (<-chan *CreateLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancer(request)
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

// CreateLoadBalancerWithCallback invokes the slb.CreateLoadBalancer API asynchronously
// api document: https://help.aliyun.com/api/slb/createloadbalancer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLoadBalancerWithCallback(request *CreateLoadBalancerRequest, callback func(response *CreateLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancer(request)
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

// CreateLoadBalancerRequest is the request struct for api CreateLoadBalancer
type CreateLoadBalancerRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SupportPrivateLink   requests.Boolean `position:"Query" name:"SupportPrivateLink"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	CloudType            string           `position:"Query" name:"CloudType"`
	AddressIPVersion     string           `position:"Query" name:"AddressIPVersion"`
	MasterZoneId         string           `position:"Query" name:"MasterZoneId"`
	Duration             requests.Integer `position:"Query" name:"Duration"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	LoadBalancerName     string           `position:"Query" name:"LoadBalancerName"`
	AddressType          string           `position:"Query" name:"AddressType"`
	SlaveZoneId          string           `position:"Query" name:"SlaveZoneId"`
	DeleteProtection     string           `position:"Query" name:"DeleteProtection"`
	LoadBalancerSpec     string           `position:"Query" name:"LoadBalancerSpec"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	Address              string           `position:"Query" name:"Address"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	EnableVpcVipFlow     string           `position:"Query" name:"EnableVpcVipFlow"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
	VpcId                string           `position:"Query" name:"VpcId"`
	PayType              string           `position:"Query" name:"PayType"`
	PricingCycle         string           `position:"Query" name:"PricingCycle"`
	Ratio                requests.Integer `position:"Query" name:"Ratio"`
}

// CreateLoadBalancerResponse is the response struct for api CreateLoadBalancer
type CreateLoadBalancerResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	LoadBalancerId   string `json:"LoadBalancerId" xml:"LoadBalancerId"`
	ResourceGroupId  string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Address          string `json:"Address" xml:"Address"`
	LoadBalancerName string `json:"LoadBalancerName" xml:"LoadBalancerName"`
	VpcId            string `json:"VpcId" xml:"VpcId"`
	VSwitchId        string `json:"VSwitchId" xml:"VSwitchId"`
	NetworkType      string `json:"NetworkType" xml:"NetworkType"`
	OrderId          int64  `json:"OrderId" xml:"OrderId"`
	AddressIPVersion string `json:"AddressIPVersion" xml:"AddressIPVersion"`
}

// CreateCreateLoadBalancerRequest creates a request to invoke CreateLoadBalancer API
func CreateCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancer", "slb", "openAPI")
	return
}

// CreateCreateLoadBalancerResponse creates a response to parse from CreateLoadBalancer response
func CreateCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
