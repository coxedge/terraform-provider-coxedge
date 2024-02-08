package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ComputeFirewallRequest struct {
	Description string `json:"description"`
}

type UpdateComputeFirewallRequest struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

func (c *Client) GetComputeFirewalls(environmentName string, organizationId string) ([]ComputeFirewall, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/firewalls?&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewalls
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeFirewallById(environmentName string, organizationId string, firewallId string) (*ComputeFirewall, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/firewalls/"+firewallId+"?&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewall
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeFirewall(firewallRequest ComputeFirewallRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(firewallRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-firewall-request?org_id="+organizationId,
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

func (c *Client) UpdateComputeFirewall(firewallUpdateRequest UpdateComputeFirewallRequest, environmentName string, organizationId string, firewallId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(firewallUpdateRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/edit-firewall-request/"+firewallId+"?operation=edit-firewall&org_id="+organizationId,
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

func (c *Client) DeleteComputeFirewallById(environmentName string, organizationId string, firewallId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/firewalls/"+firewallId+"?operation=delete-firewall&org_id="+organizationId,
		nil)
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