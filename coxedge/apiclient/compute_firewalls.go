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

type ComputeFirewallRuleRequest struct {
	CIDR         string `json:"cidr"`
	Protocol     string `json:"protocol"`
	SourceOption string `json:"source_option"`
	Port         string `json:"port"`
	Notes        string `json:"notes"`
}

type ComputeFirewallLinkedInstanceRequest struct {
	Id         string `json:"id"`
	WorkloadId string `json:"workloadId"`
}

type ComputeVPC2Request struct {
	LocationID   string `json:"locationId"`
	PrefixLength string `json:"prefix_length"`
	IPRange      string `json:"ipRange"`
	IPBlock      string `json:"ip_block"`
	Description  string `json:"description"`
}

type ComputeVPCRequest struct {
	LocationID    string                   `json:"locationId"`
	V4SubnetMask  int                      `json:"v4SubnetMask"`
	NetworkPrefix int                      `json:"networkPrefix"`
	Routes        []ComputeVPCRouteRequest `json:"routes"`
	IPRange       string                   `json:"ipRange"`
	RouteID       string                   `json:"routeId"`
	V4Subnet      string                   `json:"v4Subnet"`
	Description   string                   `json:"description"`
}

type ComputeVPCRouteRequest struct {
	Destination   string `json:"destination"`
	NetworkPrefix string `json:"networkPrefix"`
	TargetAddress string `json:"targetAddress"`
}

type ComputeReservedIPRequest struct {
	Region string `json:"region"`
	IpType string `json:"ipType"`
	Label  string `json:"label"`
}

type UpdateComputeReservedIPRequest struct {
	Label string `json:"label"`
}

type ComputeReservedIPConvertRequest struct {
	IpType    string `json:"ipType"`
	IpAddress string `json:"ipAddress"`
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

func (c *Client) CreateComputeFirewallIPv4Rule(firewallIPv4Request ComputeFirewallRuleRequest, environmentName string, organizationId string, firewallId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(firewallIPv4Request)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-ipv4-rule-request?firewallId="+firewallId+"&org_id="+organizationId,
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

func (c *Client) GetComputeFirewallsIPv4Rules(environmentName string, organizationId string, firewallId string) ([]ComputeFirewallRule, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv4-rule?firewallId="+firewallId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewallRules
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeFirewallsIPv4RuleById(environmentName string, organizationId string, firewallId string, ipv4RuleId string) (*ComputeFirewallRule, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv4-rule/"+ipv4RuleId+"?firewallId="+firewallId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewallRule
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) DeleteComputeFirewallIPv4RuleById(environmentName string, organizationId string, firewallId string, ipv4RuleId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv4-rule/"+ipv4RuleId+"?firewallId="+firewallId+"&operation=delete-ipv4-rule&org_id="+organizationId,
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

func (c *Client) GetComputeFirewallsIPv6Rules(environmentName string, organizationId string, firewallId string) ([]ComputeFirewallRule, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv6-rule?firewallId="+firewallId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewallRules
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeFirewallsIPv6RuleById(environmentName string, organizationId string, firewallId string, ipv6RuleId string) (*ComputeFirewallRule, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv6-rule/"+ipv6RuleId+"?firewallId="+firewallId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewallRule
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeFirewallIPv6Rule(firewallIPv6Request ComputeFirewallRuleRequest, environmentName string, organizationId string, firewallId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(firewallIPv6Request)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-ipv6-rule-request?firewallId="+firewallId+"&org_id="+organizationId,
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

func (c *Client) DeleteComputeFirewallIPv6RuleById(environmentName string, organizationId string, firewallId string, ipv6RuleId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/ipv6-rule/"+ipv6RuleId+"?firewallId="+firewallId+"&operation=delete-ipv6-rule&org_id="+organizationId,
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

func (c *Client) GetComputeFirewallLinkedInstances(environmentName string, organizationId string, firewallId string) ([]ComputeFirewallLinkedInstance, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/linked_instances?firewallId="+firewallId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewallLinkedInstances
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeFirewallLinkedInstanceById(environmentName string, organizationId string, firewallId string, linkedInstanceId string) (*ComputeFirewallLinkedInstance, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/linked_instances/"+linkedInstanceId+"?firewallId="+firewallId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeFirewallLinkedInstance
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeFirewallLinkedInstance(firewallLinkedInstanceRequest ComputeFirewallLinkedInstanceRequest, environmentName string, organizationId string, firewallId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(firewallLinkedInstanceRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/link-instance-request?id="+firewallId+"&org_id="+organizationId,
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

func (c *Client) DeleteComputeFirewallLinkedInstanceById(environmentName string, organizationId string, firewallId string, linkedInstanceId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/linked_instances/"+linkedInstanceId+"?firewallId="+firewallId+"&operation=unlink-instance&org_id="+organizationId,
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

func (c *Client) GetComputeVPC2(environmentName string, organizationId string) ([]ComputeVPC2, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/vpc2?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeVPC2s
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeVPC2ById(environmentName string, organizationId string, vpc2Id string) (*ComputeVPC2, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/vpc2/"+vpc2Id+"?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeVPC2
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeVPC2(vpc2Request ComputeVPC2Request, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(vpc2Request)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-vpc2-request?org_id="+organizationId,
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

func (c *Client) DeleteComputeVPC2ById(environmentName string, organizationId string, vpc2Id string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/vpc2/"+vpc2Id+"?operation=delete-vpc2&org_id="+organizationId,
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

func (c *Client) GetComputeVPC(environmentName string, organizationId string) ([]ComputeVPC, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/vpc?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeVPCs
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeVPCById(environmentName string, organizationId string, vpcId string) (*ComputeVPC, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/vpc/"+vpcId+"?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeVPC
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeVPC(vpc2Request ComputeVPCRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(vpc2Request)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-vpc-request?org_id="+organizationId,
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

func (c *Client) DeleteComputeVPCById(environmentName string, organizationId string, vpcId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/vpc/"+vpcId+"?operation=delete-vpc&org_id="+organizationId,
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

func (c *Client) GetComputeReservedIPs(environmentName string, organizationId string) ([]ComputeReservedIP, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/reserved-ip?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeReservedIPs
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeReservedIPById(environmentName string, organizationId string, reservedIpId string) (*ComputeReservedIP, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/reserved-ip/"+reservedIpId+"?org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeReservedIP
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeReservedIP(reservedIPRequest ComputeReservedIPRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(reservedIPRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-reserved-ip-request?org_id="+organizationId,
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

func (c *Client) DeleteComputeReservedIPById(environmentName string, organizationId string, reservedIPId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/reserved-ip/"+reservedIPId+"?operation=delete-reserved-ip&org_id="+organizationId,
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

func (c *Client) UpdateComputeReservedIP(reservedIPRequest UpdateComputeReservedIPRequest, environmentName string, organizationId string, reservedIPId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(reservedIPRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/reserved-ip/"+reservedIPId+"?operation=edit-reserved-ip&org_id="+organizationId,
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

func (c *Client) CreateComputeReservedIPConvert(reservedIPRequest ComputeReservedIPConvertRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(reservedIPRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/convert-ip-request?org_id="+organizationId,
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
