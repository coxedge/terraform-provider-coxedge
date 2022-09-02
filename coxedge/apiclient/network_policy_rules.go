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
	EnvironmentName string              `json:"-"`
	NetworkPolicy   []NetworkPolicyList `json:"network_policy"`
}

type NetworkPolicyList struct {
	EnvironmentName string `json:"-"`
	Id              string `json:"id"`
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

func (c *Client) GetNetworkPolicyRuleWorkload(environmentName string, id string, organizationId string) ([]NetworkPolicyRule, error) {
	//Create the request
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/networkpolicyrules?workloadId="+id+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}

	//Execute request
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
func (c *Client) CreateNetworkPolicyRule(newNetworkPolicyRule NetworkPolicyRuleCreateRequest, organizationId string) ([]NetworkPolicyRule, error) {
	var networkResponse []NetworkPolicyRule
	//Marshal the request
	for _, entry := range newNetworkPolicyRule.NetworkPolicy {
		jsonBytes, err := json.Marshal(entry)
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
		networkResponse = append(networkResponse, wrappedAPIStruct.Data)
	}
	return networkResponse, nil

}

//UpdateNetworkPolicyRule Update a networkPolicyRule
func (c *Client) UpdateNetworkPolicyRule(networkPolicyRuleId string, newNetworkPolicyRule NetworkPolicyRuleCreateRequest, organizationId string) ([]NetworkPolicyRule, error) {
	var networkPolicy []NetworkPolicyRule

	for _, entry := range newNetworkPolicyRule.NetworkPolicy {
		//Marshal the request
		jsonBytes, err := json.Marshal(entry)
		if err != nil {
			return nil, err
		}
		//Wrap bytes in reader
		bReader := bytes.NewReader(jsonBytes)
		//Create the request
		request, err := http.NewRequest("PUT",
			CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+newNetworkPolicyRule.EnvironmentName+"/networkpolicyrules/"+entry.Id+"?org_id="+organizationId,
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
		networkPolicy = append(networkPolicy, wrappedAPIStruct.Data)
	}

	return networkPolicy, nil
}

//DeleteNetworkPolicyRule Delete networkPolicyRule in account by id
func (c *Client) DeleteNetworkPolicyRule(environmentName string, id string, organizationId string, newNetworkPolicyRule NetworkPolicyRuleCreateRequest) error {
	for _, entry := range newNetworkPolicyRule.NetworkPolicy {
		//Create the request
		request, err := http.NewRequest("DELETE",
			CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/networkpolicyrules/"+entry.Id+"?org_id="+organizationId,
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
	}

	return nil
}
