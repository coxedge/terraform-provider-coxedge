package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RouteRequest struct {
	Id               string   `json:"id,omitempty"`
	VpcId            string   `json:"vpcId"`
	Name             string   `json:"name"`
	DestinationCidrs []string `json:"destinationCidrs"`
	NextHops         []string `json:"nextHops"`
	StackId          string   `json:"stackId,omitempty"`
	Status           string   `json:"status"`
}

func (c *Client) GetAllRoutes(vpcId string, environmentName string, organizationId string) ([]Route, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/routes?vpc_id="+vpcId+"&org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedRoutesData
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) CreateRoute(routeRequest RouteRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(routeRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/routes?org_id="+organizationId,
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

func (c *Client) GetRoute(routeRequestId string, vpcId string, environmentName string, organizationId string) (*Route, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/routes/"+routeRequestId+"?vpc_id="+vpcId+"&org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedRoute
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) DeleteRoute(routeRequest RouteRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(routeRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/routes/"+routeRequest.Id+"?vpc_id="+routeRequest.VpcId+"&operation=delete&org_id="+organizationId,
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
