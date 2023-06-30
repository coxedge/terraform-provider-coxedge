package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type VPCRequest struct {
	Id         string   `json:"id,omitempty"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
	StackId    string   `json:"stackId"`
	Cidr       string   `json:"cidr"`
	DefaultVpc bool     `json:"defaultVpc"`
	Status     string   `json:"status"`
	Created    string   `json:"created"`
	Subnets    []string `json:"subnets,omitempty"`
	Routes     []string `json:"routes,omitempty"`
}

func (c *Client) GetAllVPCs(environmentName string, organizationId string) ([]VPCs, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/vpcs?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedVPCsData
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) CreateVPCNetwork(vpcRequest VPCRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(vpcRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/vpcs?org_id="+organizationId,
		bReader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
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

func (c *Client) GetVPCNetwork(vpcRequestId string, environmentName string, organizationId string) (*VPCs, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/vpcs/"+vpcRequestId+"?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedVPCs
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) DeleteVPCNetwork(vpcRequest VPCRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(vpcRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/vpcs/"+vpcRequest.Id+"?operation=delete&org_id="+organizationId,
		bReader)
	if err != nil {
		return nil, err
	}

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
