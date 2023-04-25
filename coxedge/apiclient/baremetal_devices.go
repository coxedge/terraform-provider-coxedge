/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CreateBareMetalDeviceRequest struct {
	Quantity        int      `json:"quantity"`
	LocationName    string   `json:"locationName"`
	HasUserData     *bool    `json:"hasUserData"`
	HasSshData      *bool    `json:"hasSshData"`
	ProductOptionId int      `json:"productOptionId"`
	ProductId       string   `json:"productId"`
	OsName          string   `json:"osName"`
	Server          []Server `json:"server"`
	SshKey          string   `json:"sshKey,omitempty"`
	SshKeyName      string   `json:"sshKeyName,omitempty"`
	SshKeyId        string   `json:"sshKeyId,omitempty"`
	UserData        string   `json:"user_data,omitempty"`
}

type Server struct {
	Hostname string `json:"hostname"`
}

/*
GetBareMetalDevices get all BareMetal devices
*/
func (c *Client) GetBareMetalDevices(environmentName string, organizationId string) ([]BareMetalDevice, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/devices?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalDevices
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

/*
GetBareMetalDeviceById get BareMetal device by Id
*/
func (c *Client) GetBareMetalDeviceById(environmentName string, organizationId string, requestedId string) (*BareMetalDevice, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/devices/"+requestedId+"?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalDevice
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

/*
CreateBareMetalDevice create BareMetal device(s)
*/
func (c *Client) CreateBareMetalDevice(createRequest CreateBareMetalDeviceRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-create-request?org_id="+organizationId,
		bReader)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct TaskStatusResponse
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct, nil
}
