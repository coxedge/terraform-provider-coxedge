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

//GetPredefinedEdgeLogics Get predefined edge logic of sites by site Id
func (c *Client) GetPredefinedEdgeLogics(environmentName string, organizationId string, siteId string) (*EdgeLogic, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/predefinededgerules/"+siteId+"?org_id="+organizationId,
		nil,
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedEdgeLogic
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}
