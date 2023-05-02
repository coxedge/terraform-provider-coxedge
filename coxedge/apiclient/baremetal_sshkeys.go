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

type CreateBareMetalSSHKeyRequest struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

/*
GetBareMetalSSHKeys get all SSH Keys
*/
func (c *Client) GetBareMetalSSHKeys(environmentName string, organizationId string) ([]BareMetalSSHKey, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/ssh-key?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalSSHKeys
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

/*
GetBareMetalSSHKeyById get SSH Key by Id
*/
func (c *Client) GetBareMetalSSHKeyById(environmentName string, organizationId string, resourceId string) (*BareMetalSSHKey, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/ssh-key/"+resourceId+"?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalSSHKey
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

/*
CreateBareMetalSSHKey create SSH key in BareMetal
*/
func (c *Client) CreateBareMetalSSHKey(createRequest CreateBareMetalSSHKeyRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/ssh-key?org_id="+organizationId,
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

/*
DeleteBareMetalSSHKeyById delete SSH key in BareMetal by Id
*/
func (c *Client) DeleteBareMetalSSHKeyById(environmentName string, organizationId string, resourceId string) (*TaskStatusResponse, error) {

	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/ssh-key/"+resourceId+"?operation=delete&org_id="+organizationId,
		nil)
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
