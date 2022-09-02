package apiclient

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetWorkloadInstances(environmentName string, organizationId string, workloadId string) ([]WorkloadInstance, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/instances?workloadId="+workloadId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperWorkloadInstances
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}
