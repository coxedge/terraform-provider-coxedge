package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ComputeStorageRequest struct {
	Region    string `json:"region"`
	SizeGB    string `json:"size_gb"`
	Label     string `json:"label,omitempty"`
	BlockType string `json:"block_type"`
}

func (c *Client) GetComputeStorages(environmentName string, organizationId string) ([]ComputeStorage, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/storages?&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeStorages
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

func (c *Client) GetComputeStorageById(environmentName string, organizationId string, storageId string) (*ComputeStorage, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/storages/"+storageId+"?&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrapperComputeStorage
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) CreateComputeStorage(storageRequest ComputeStorageRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(storageRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/add-storage-request?org_id="+organizationId,
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
