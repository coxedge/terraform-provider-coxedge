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

type NetworkPolicyRuleCreateRequest struct {
	EnvironmentName string `json:"-"`
	WorkloadId      string `json:"workloadId"`
	Description     string `json:"description"`
	Protocol        string `json:"protocol"`
	Type            string `json:"type"`
	Action          string `json:"action"`
	Source          string `json:"source"`
	PortRange       string `json:"portRange"`
}

//GetNetworkPolicyRules Get networkPolicyRules in account
func (c *Client) GetNetworkPolicyRules(environmentName string, organizationId string) ([]NetworkPolicyRule, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/networkpolicyrules"+"?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedNetworkPolicyRules
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetNetworkPolicyRule Get networkPolicyRule in account by id
func (c *Client) GetNetworkPolicyRule(environmentName string, id string, organizationId string) (*NetworkPolicyRule, error) {
	//Create the request
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/networkpolicyrules/"+id+"?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}

	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	//Unmarshal, unwrap, and return
	var wrappedAPIStruct WrappedNetworkPolicyRule
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//CreateNetworkPolicyRule Create the networkPolicyRule
func (c *Client) CreateNetworkPolicyRule(newNetworkPolicyRule NetworkPolicyRuleCreateRequest, organizationId string) (*NetworkPolicyRule, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newNetworkPolicyRule)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+newNetworkPolicyRule.EnvironmentName+"/networkpolicyrules?org_id="+organizationId,
		bReader,
	)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedNetworkPolicyRule
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//UpdateNetworkPolicyRule Update a networkPolicyRule
func (c *Client) UpdateNetworkPolicyRule(networkPolicyRuleId string, newNetworkPolicyRule NetworkPolicyRuleCreateRequest, organizationId string) (*NetworkPolicyRule, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newNetworkPolicyRule)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+newNetworkPolicyRule.EnvironmentName+"/networkpolicyrules/"+networkPolicyRuleId+"?org_id="+organizationId,
		bReader,
	)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedNetworkPolicyRule
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//DeleteNetworkPolicyRule Delete networkPolicyRule in account by id
func (c *Client) DeleteNetworkPolicyRule(environmentName string, id string, organizationId string) error {
	//Create the request
	request, err := http.NewRequest("DELETE",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/networkpolicyrules/"+id+"?org_id="+organizationId,
		nil,
	)
	if err != nil {
		return err
	}

	//Execute request
	_, err = c.doRequest(request)
	if err != nil {
		return err
	}
	return nil
}
