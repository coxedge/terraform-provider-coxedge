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

type ComputeWorkloadIPv6ReverseDNSRequest struct {
	Ip      string `json:"ip"`
	Reverse string `json:"reverse"`
}

type ComputeWorkloadFirewallGroupRequest struct {
	FirewallId string `json:"firewallId"`
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

func (c *Client) GetComputeWorkloadById(environmentName string, organizationId string, workloadId string) (*ComputeWorkload, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/workloads/"+workloadId+"?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeWorkload
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
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

func (c *Client) DeleteComputeWorkloadById(environmentName string, organizationId string, workloadId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/workloads/"+workloadId+"?operation=delete-workload&org_id="+organizationId,
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

func (c *Client) GetComputeWorkloadIPv4ById(environmentName string, organizationId string, workloadId string) ([]IPv4Configuration, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv4?id="+workloadId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperIPv4Configuration
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeWorkloadIPv6ById(environmentName string, organizationId string, workloadId string) ([]IPv6Configuration, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv6?id="+workloadId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperIPv6Configuration
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeWorkloadIPv6ReverseDNSById(environmentName string, organizationId string, workloadId string) ([]IPv6ReverseDNSConfiguration, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/reverse-dns?id="+workloadId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperIPv6ReverseDNSConfiguration
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeWorkloadIPv6ReverseDNSById(reverseDNSRequest ComputeWorkloadIPv6ReverseDNSRequest, environmentName string, organizationId string, workloadId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(reverseDNSRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-reverse-dns-request?workloadId="+workloadId+"&org_id="+organizationId,
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

func (c *Client) DeleteComputeWorkloadIPv6ReverseDNSById(environmentName string, organizationId string, workloadId string, ip string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/reverse-dns/"+ip+"?workloadId="+workloadId+"&org_id="+organizationId+"&ip="+ip+"&operation=delete-reverse-dns",
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

func (c *Client) GetComputeWorkloadFirewallGroupById(environmentName string, organizationId string, workloadId string) (*FirewallGroup, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/workload-firewall/"+workloadId+"?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperFirewallGroup
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) UpdateComputeWorkloadFirewallGroupById(firewallGroupRequest ComputeWorkloadFirewallGroupRequest, environmentName string, organizationId string, workloadId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(firewallGroupRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PATCH",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/workload-firewall/"+workloadId+"?org_id="+organizationId,
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
