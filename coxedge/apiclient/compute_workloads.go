package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ComputeWorkloadRequest struct {
	IsIPv6                 bool   `json:"isIPv6"`
	NoPublicIPv4           bool   `json:"noPublicIPv4"`
	IsVirtualPrivateClouds bool   `json:"isVirtualPrivateClouds"`
	IsVPC2                 bool   `json:"isVPC2"`
	OperatingSystemId      string `json:"operatingSystemId"`
	LocationId             string `json:"locationId"`
	PlanId                 string `json:"planId"`
	Hostname               string `json:"hostname"`
	Label                  string `json:"label"`
	Name                   string `json:"name"`
	FirstBootSshKey        string `json:"firstBootSshKey"`
	SshKeyName             string `json:"sshKeyName"`
	FirewallId             string `json:"firewallId"`
	UserData               string `json:"userData"`
}

func (c *Client) GetComputeWorkloads(environmentName string, organizationId string) ([]ComputeWorkload, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/workloads?&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeWorkloads
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeWorkload(workloadRequest ComputeWorkloadRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(workloadRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-workload-request?org_id="+organizationId,
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
