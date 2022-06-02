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

type DeliveryDomainCreateRequest struct {
	EnvironmentName string `json:"-"`
	Domain          string `json:"domain"`
}

//GetDeliveryDomains Get deliveryDomains in account
func (c *Client) GetDeliveryDomains(environmentName string) ([]DeliveryDomain, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliverydomains",
		nil,
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedDeliveryDomains
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetDeliveryDomain Get deliveryDomain in account by id
func (c *Client) GetDeliveryDomain(environmentName string, id string) (*DeliveryDomain, error) {
	//Create the request
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliverydomains/"+id,
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
	var wrappedAPIStruct WrappedDeliveryDomain
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//CreateDeliveryDomain Create the deliveryDomain
func (c *Client) CreateDeliveryDomain(newDeliveryDomain DeliveryDomainCreateRequest) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newDeliveryDomain)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+newDeliveryDomain.EnvironmentName+"/deliverydomains",
		bReader,
	)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var taskStatusResp TaskStatusResponse
	err = json.Unmarshal(respBytes, &taskStatusResp)
	if err != nil {
		return nil, err
	}
	return &taskStatusResp, nil
}

//DeleteDeliveryDomain Delete deliveryDomain in account by id
func (c *Client) DeleteDeliveryDomain(environmentName string, id string) error {
	//Create the request
	request, err := http.NewRequest("DELETE",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliverydomains/"+id,
		nil,
	)
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
