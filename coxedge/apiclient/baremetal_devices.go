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
