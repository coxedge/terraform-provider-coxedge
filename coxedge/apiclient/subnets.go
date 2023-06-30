package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SubnetRequest struct {
	Id    string `json:"id,omitempty"`
	VpcId string `json:"vpcId"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Cidr  string `json:"cidr"`
}

func (c *Client) GetAllSubnets(vpcId string, environmentName string, organizationId string) ([]Subnets, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/subnets?vpc_id="+vpcId+"&org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedSubnetsData
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) CreateSubnet(subnetRequest SubnetRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(subnetRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/subnets?org_id="+organizationId,
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

func (c *Client) GetSubnet(subnetRequestId string, environmentName string, organizationId string) (*Subnets, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/subnets/"+subnetRequestId+"?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedSubnet
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) DeleteSubnet(subnetRequest SubnetRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(subnetRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/subnets/"+subnetRequest.Id+"?operation=delete&org_id="+organizationId,
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
