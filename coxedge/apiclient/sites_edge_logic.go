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

type PredefinedEdgeLogicRequest struct {
	ForceWwwEnabled           *bool    `json:"forceWwwEnabled,omitempty"`
	RobotsTxtEnabled          *bool    `json:"robotsTxtEnabled,omitempty"`
	RobotsTxtFile             string   `json:"robotsTxtFile,omitempty"`
	PseudoStreamingEnabled    *bool    `json:"pseudoStreamingEnabled,omitempty"`
	ReferrerProtectionEnabled *bool    `json:"referrerProtectionEnabled,omitempty"`
	AllowEmptyReferrer        *bool    `json:"allowEmptyReferrer,omitempty"`
	ReferrerList              []string `json:"referrerList,omitempty"`
}

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

//UpdatePredefinedEdgeLogic Update predefined edge logic of sites by site Id
func (c *Client) UpdatePredefinedEdgeLogic(predefinedEdgeLogicRequest PredefinedEdgeLogicRequest, environmentName string, organizationId string, siteId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(predefinedEdgeLogicRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PATCH",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/predefinededgerules/"+siteId+"?org_id="+organizationId,
		bReader,
	)
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
