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

// DescribeRouteEntryList invokes the vpc.DescribeRouteEntryList API synchronously
// api document: https://help.aliyun.com/api/vpc/describerouteentrylist.html
func (client *Client) DescribeRouteEntryList(request *DescribeRouteEntryListRequest) (response *DescribeRouteEntryListResponse, err error) {
	response = CreateDescribeRouteEntryListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRouteEntryListWithChan invokes the vpc.DescribeRouteEntryList API asynchronously
// api document: https://help.aliyun.com/api/vpc/describerouteentrylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouteEntryListWithChan(request *DescribeRouteEntryListRequest) (<-chan *DescribeRouteEntryListResponse, <-chan error) {
	responseChan := make(chan *DescribeRouteEntryListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouteEntryList(request)
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

// DescribeRouteEntryListWithCallback invokes the vpc.DescribeRouteEntryList API asynchronously
// api document: https://help.aliyun.com/api/vpc/describerouteentrylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouteEntryListWithCallback(request *DescribeRouteEntryListRequest, callback func(response *DescribeRouteEntryListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouteEntryListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouteEntryList(request)
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

// DescribeRouteEntryListRequest is the request struct for api DescribeRouteEntryList
type DescribeRouteEntryListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouteEntryName       string           `position:"Query" name:"RouteEntryName"`
	NextToken            string           `position:"Query" name:"NextToken"`
	RouteEntryType       string           `position:"Query" name:"RouteEntryType"`
	IpVersion            string           `position:"Query" name:"IpVersion"`
	NextHopId            string           `position:"Query" name:"NextHopId"`
	NextHopType          string           `position:"Query" name:"NextHopType"`
	RouteTableId         string           `position:"Query" name:"RouteTableId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DestinationCidrBlock string           `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MaxResult            requests.Integer `position:"Query" name:"MaxResult"`
	RouteEntryId         string           `position:"Query" name:"RouteEntryId"`
}

// DescribeRouteEntryListResponse is the response struct for api DescribeRouteEntryList
type DescribeRouteEntryListResponse struct {
	*responses.BaseResponse
	RequestId   string                              `json:"RequestId" xml:"RequestId"`
	NextToken   string                              `json:"NextToken" xml:"NextToken"`
	RouteEntrys RouteEntrysInDescribeRouteEntryList `json:"RouteEntrys" xml:"RouteEntrys"`
}

// CreateDescribeRouteEntryListRequest creates a request to invoke DescribeRouteEntryList API
func CreateDescribeRouteEntryListRequest() (request *DescribeRouteEntryListRequest) {
	request = &DescribeRouteEntryListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouteEntryList", "Vpc", "openAPI")
	return
}

// CreateDescribeRouteEntryListResponse creates a response to parse from DescribeRouteEntryList response
func CreateDescribeRouteEntryListResponse() (response *DescribeRouteEntryListResponse) {
	response = &DescribeRouteEntryListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
