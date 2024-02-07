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

type UpdateComputeStorageSizeRequest struct {
	SizeGB string `json:"size_gb"`
}

type UpdateComputeStorageLabelRequest struct {
	Label string `json:"label"`
}

type AttachComputeStorageInstanceRequest struct {
	Live       bool   `json:"live"`
	InstanceId string `json:"instance_id"`
}

type DetachComputeStorageInstanceRequest struct {
	Live bool `json:"live"`
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

func (c *Client) UpdateComputeStorageSize(sizeRequest UpdateComputeStorageSizeRequest, environmentName string, organizationId string, storageId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(sizeRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/update-storage-size-request/"+storageId+"?operation=update-storage-size&org_id="+organizationId+"&storageId="+storageId,
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

func (c *Client) UpdateComputeStorageLabel(labelRequest UpdateComputeStorageLabelRequest, environmentName string, organizationId string, storageId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(labelRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/update-storage-label-request/"+storageId+"?operation=update-storage-label&org_id="+organizationId+"&storageId="+storageId,
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

func (c *Client) AttachComputeStorageInstance(attachRequest AttachComputeStorageInstanceRequest, environmentName string, organizationId string, storageId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(attachRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/update-storage-attach-to-instance-request/"+storageId+"?operation=update-storage-attach-to-instance&org_id="+organizationId+"&storageId="+storageId,
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

func (c *Client) DetachComputeStorageInstance(detachRequest DetachComputeStorageInstanceRequest, environmentName string, organizationId string, storageId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(detachRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/update-storage-detach-from-instance-request/"+storageId+"?operation=update-storage-detach-from-instance&org_id="+organizationId+"&storageId="+storageId,
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

func (c *Client) DeleteComputeStorageById(environmentName string, organizationId string, storageId string) error {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeComputeServiceCode+"/"+environmentName+"/storages/"+storageId+"?operation=delete-storage&org_id="+organizationId,
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
