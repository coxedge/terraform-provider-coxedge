package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//GetFirewallRules Get FirewallRules in account
func (c *Client) GetFirewallRules(siteId string) ([]FirewallRule, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/firewallrules?siteId="+siteId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedFirewallRules
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetFirewallRule Get FirewallRule in account by id
func (c *Client) GetFirewallRule(siteId string, id string) (*FirewallRule, error) {
	//Create the request
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/firewallrules/"+id+"?siteId="+siteId, nil)
	if err != nil {
		return nil, err
	}

	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	//Unmarshal, unwrap, and return
	var wrappedAPIStruct WrappedFirewallRule
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//CreateFirewallRule Create the FirewallRule
func (c *Client) CreateFirewallRule(newFirewallRule FirewallRule) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newFirewallRule)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/firewallrules?siteId="+newFirewallRule.SiteId, bReader)
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

//UpdateFirewallRule Update a FirewallRule
func (c *Client) UpdateFirewallRule(FirewallRuleId string, newFirewallRule FirewallRule) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newFirewallRule)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT", CoxEdgeAPIBase+"/firewallrules/"+FirewallRuleId+"?siteId="+newFirewallRule.SiteId, bReader)
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

//DeleteFirewallRule Delete FirewallRule in account by id
func (c *Client) DeleteFirewallRule(siteId string, id string) error {
	//Create the request
	request, err := http.NewRequest("DELETE", CoxEdgeAPIBase+"/firewallrules/"+id+"?siteId="+siteId, nil)
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
