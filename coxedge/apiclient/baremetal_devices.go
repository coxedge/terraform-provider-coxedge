/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CreateBareMetalDeviceRequest struct {
	Quantity        int      `json:"quantity"`
	LocationName    string   `json:"locationName"`
	HasUserData     *bool    `json:"hasUserData"`
	HasSshData      *bool    `json:"hasSshData"`
	ProductOptionId int      `json:"productOptionId"`
	ProductId       string   `json:"productId"`
	OsName          string   `json:"osName"`
	Server          []Server `json:"server"`
	SshKey          string   `json:"sshKey,omitempty"`
	SshKeyName      string   `json:"sshKeyName,omitempty"`
	SshKeyId        string   `json:"sshKeyId,omitempty"`
	UserData        string   `json:"user_data,omitempty"`
	Name            string   `json:"name,omitempty"`
}

type Server struct {
	Hostname string `json:"hostname"`
}

type EditBareMetalDeviceRequest struct {
	Name        string   `json:"name,omitempty"`
	Hostname    string   `json:"hostname,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	PowerStatus string   `json:"power_status,omitempty"`
}

type CustomChartRequest struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

type ConnectIPMIRequest struct {
	CustomIP string `json:"customIP,omitempty"`
}

/*
GetBareMetalDevices get all BareMetal devices
*/
func (c *Client) GetBareMetalDevices(environmentName string, organizationId string) ([]BareMetalDevice, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/devices?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalDevices
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

/*
GetBareMetalDeviceById get BareMetal device by Id
*/
func (c *Client) GetBareMetalDeviceById(environmentName string, organizationId string, requestedId string) (*BareMetalDevice, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/devices/"+requestedId+"?org_id="+organizationId, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalDevice
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

/*
CreateBareMetalDevice create BareMetal device(s)
*/
func (c *Client) CreateBareMetalDevice(createRequest CreateBareMetalDeviceRequest, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-create-request?org_id="+organizationId,
		bReader)
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

/*
DeleteBareMetalDeviceById delete BareMetal device by Id
*/
func (c *Client) DeleteBareMetalDeviceById(deviceId string, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/devices/"+deviceId+"?operation=delete&org_id="+organizationId,
		nil)
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

/*
EditBareMetalDeviceById edit BareMetal device by Id
*/
func (c *Client) EditBareMetalDeviceById(editRequest EditBareMetalDeviceRequest, deviceId string, environmentName string, organizationId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(editRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PATCH",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-setting/"+deviceId+"?org_id="+organizationId,
		bReader)
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

/*
EditBareMetalDevicePowerById edit BareMetal device power by Id
*/
func (c *Client) EditBareMetalDevicePowerById(deviceId string, operation string, environmentName string, organizationId string) (*TaskStatusResponse, error) {

	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/devices/"+deviceId+"?operation="+operation+"&org_id="+organizationId,
		nil)
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

/*
GetBareMetalDeviceChartsById get BareMetal device charts by Id
*/
func (c *Client) GetBareMetalDeviceChartsById(environmentName string, organizationId string, requestedId string) ([]BareMetalDeviceChart, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-charts?id="+requestedId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalDeviceCharts
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

/*
PostBareMetalDeviceCustomChartsById request BareMetal device custom charts by Id
*/
func (c *Client) PostBareMetalDeviceCustomChartsById(customRequest CustomChartRequest, environmentName string, organizationId string, requestedId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(customRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)

	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-charts/"+requestedId+"?operation=custom&org_id="+organizationId,
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

/*
GetBareMetalDeviceSensorsById get BareMetal device sensors by Id
*/
func (c *Client) GetBareMetalDeviceSensorsById(environmentName string, organizationId string, requestedId string) ([]BareMetalDeviceSensor, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-sensors-list?id="+requestedId+"&org_id="+organizationId,
		nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedBareMetalDeviceSensors
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

/*
PostBareMetalDeviceConnectToIPMIById request BareMetal device connect to ipmi
*/
func (c *Client) PostBareMetalDeviceConnectToIPMIById(ipmiRequest ConnectIPMIRequest, environmentName string, organizationId string, requestedId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(ipmiRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)

	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-sensors-list/"+requestedId+"?operation=connect&org_id="+organizationId,
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

/*
PostBareMetalDeviceClearIPMIById request BareMetal device clear ipmi address
*/
func (c *Client) PostBareMetalDeviceClearIPMIById(environmentName string, organizationId string, requestedId string) (*TaskStatusResponse, error) {

	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeBareMetalServiceCode+"/"+environmentName+"/device-sensors-list/"+requestedId+"?operation=clear&org_id="+organizationId,
		nil)
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
