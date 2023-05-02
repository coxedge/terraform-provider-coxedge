/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package apiclient

import (
	"encoding/json"
	"net/http"
)

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
func (c *Client) GetBareMetalSSHKeyById(environmentName string, organizationId string, requestedId string) (*BareMetalSSHKey, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/ssh-key/"+requestedId+"?org_id="+organizationId, nil)
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
