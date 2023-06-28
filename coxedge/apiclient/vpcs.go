package apiclient

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetAllVPCs(environmentName string) ([]VPCs, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/vpcs", nil)
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
