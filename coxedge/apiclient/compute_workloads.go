package apiclient

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetComputeWorkloads(environmentName string, organizationId string) ([]ComputeWorkload, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/workloads?&org_id="+organizationId,
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