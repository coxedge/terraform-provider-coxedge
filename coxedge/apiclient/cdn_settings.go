package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//GetCDNSettings Get cdnSettings in account by id
func (c *Client) GetCDNSettings(environmentName string, id string) (*CDNSettings, error) {
	//Create the request
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/cdnSettings/"+id,
		nil,
	)
	if err != nil {
		return nil, err
	}

	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	//Unmarshal, unwrap, and return
	var wrappedAPIStruct WrappedCDNSettings
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//UpdateCDNSettings Update a cdnSettings
func (c *Client) UpdateCDNSettings(cdnSettingsId string, newCDNSettings CDNSettings) (*CDNSettings, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newCDNSettings)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+newCDNSettings.EnvironmentName+"/cdnsettings/"+cdnSettingsId,
		bReader,
	)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedCDNSettings
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

func (c *Client) PurgeCDN(environmentName string, siteId string, options CDNPurgeOptions) (*TaskStatusResponse, error) {
	//Derive the operation type
	operationType := "purge"
	if len(options.Items) == 0 {
		operationType = "purgeAll"
	}
	//Marshal the request
	jsonBytes, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/cdnsettings/"+siteId+"?operation="+operationType,
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
