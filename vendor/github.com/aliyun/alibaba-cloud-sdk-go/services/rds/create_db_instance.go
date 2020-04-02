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

// CreateDBInstance invokes the rds.CreateDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/createdbinstance.html
func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
	response = CreateCreateDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBInstanceWithChan invokes the rds.CreateDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceWithChan(request *CreateDBInstanceRequest) (<-chan *CreateDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBInstance(request)
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

// CreateDBInstanceWithCallback invokes the rds.CreateDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceWithCallback(request *CreateDBInstanceRequest, callback func(response *CreateDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateDBInstance(request)
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

// CreateDBInstanceRequest is the request struct for api CreateDBInstance
type CreateDBInstanceRequest struct {
	*requests.RpcRequest
	DBParamGroupId                 string           `position:"Query" name:"DBParamGroupId"`
	ResourceOwnerId                requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage              requests.Integer `position:"Query" name:"DBInstanceStorage"`
	SystemDBCharset                string           `position:"Query" name:"SystemDBCharset"`
	EngineVersion                  string           `position:"Query" name:"EngineVersion"`
	ResourceGroupId                string           `position:"Query" name:"ResourceGroupId"`
	TargetDedicatedHostIdForMaster string           `position:"Query" name:"TargetDedicatedHostIdForMaster"`
	DBInstanceDescription          string           `position:"Query" name:"DBInstanceDescription"`
	BusinessInfo                   string           `position:"Query" name:"BusinessInfo"`
	Period                         string           `position:"Query" name:"Period"`
	EncryptionKey                  string           `position:"Query" name:"EncryptionKey"`
	OwnerId                        requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceClass                string           `position:"Query" name:"DBInstanceClass"`
	SecurityIPList                 string           `position:"Query" name:"SecurityIPList"`
	VSwitchId                      string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress               string           `position:"Query" name:"PrivateIpAddress"`
	TargetDedicatedHostIdForLog    string           `position:"Query" name:"TargetDedicatedHostIdForLog"`
	AutoRenew                      string           `position:"Query" name:"AutoRenew"`
	RoleARN                        string           `position:"Query" name:"RoleARN"`
	TunnelId                       string           `position:"Query" name:"TunnelId"`
	ZoneId                         string           `position:"Query" name:"ZoneId"`
	InstanceNetworkType            string           `position:"Query" name:"InstanceNetworkType"`
	ConnectionMode                 string           `position:"Query" name:"ConnectionMode"`
	ClientToken                    string           `position:"Query" name:"ClientToken"`
	TargetDedicatedHostIdForSlave  string           `position:"Query" name:"TargetDedicatedHostIdForSlave"`
	ZoneIdSlave1                   string           `position:"Query" name:"ZoneIdSlave1"`
	ZoneIdSlave2                   string           `position:"Query" name:"ZoneIdSlave2"`
	DBIsIgnoreCase                 string           `position:"Query" name:"DBIsIgnoreCase"`
	Engine                         string           `position:"Query" name:"Engine"`
	DBTimeZone                     string           `position:"Query" name:"DBTimeZone"`
	DBInstanceStorageType          string           `position:"Query" name:"DBInstanceStorageType"`
	DedicatedHostGroupId           string           `position:"Query" name:"DedicatedHostGroupId"`
	DBInstanceNetType              string           `position:"Query" name:"DBInstanceNetType"`
	ResourceOwnerAccount           string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string           `position:"Query" name:"OwnerAccount"`
	UsedTime                       string           `position:"Query" name:"UsedTime"`
	VPCId                          string           `position:"Query" name:"VPCId"`
	Category                       string           `position:"Query" name:"Category"`
	PayType                        string           `position:"Query" name:"PayType"`
}

// CreateDBInstanceResponse is the response struct for api CreateDBInstance
type CreateDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId          string `json:"OrderId" xml:"OrderId"`
	ConnectionString string `json:"ConnectionString" xml:"ConnectionString"`
	Port             string `json:"Port" xml:"Port"`
}

// CreateCreateDBInstanceRequest creates a request to invoke CreateDBInstance API
func CreateCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
	request = &CreateDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDBInstance", "rds", "openAPI")
	return
}

// CreateCreateDBInstanceResponse creates a response to parse from CreateDBInstance response
func CreateCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
	response = &CreateDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
