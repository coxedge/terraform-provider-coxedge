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

type ScriptCreateRequest struct {
	SiteId string   `json:"siteId,omitempty"`
	Name   string   `json:"name,omitempty"`
	Routes []string `json:"routes,omitempty"`
	Code   string   `json:"code,omitempty"`
}

//GetScripts Get Scripts in account
func (c *Client) GetScripts() ([]Script, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/scripts", nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedScripts
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetScript Get Script in account by id
func (c *Client) GetScript(id string) (*Script, error) {
	//Create the request
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/scripts/"+id, nil)
	if err != nil {
		return nil, err
	}

	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	//Unmarshal, unwrap, and return
	var wrappedAPIStruct WrappedScript
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//CreateScript Create the Script
func (c *Client) CreateScript(newScript ScriptCreateRequest) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newScript)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/scripts", bReader)
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

//UpdateScript Update a Script
func (c *Client) UpdateScript(ScriptId string, newScript ScriptCreateRequest) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newScript)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT", CoxEdgeAPIBase+"/scripts/"+ScriptId, bReader)
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

//DeleteScript Delete Script in account by id
func (c *Client) DeleteScript(id string) error {
	//Create the request
	request, err := http.NewRequest("DELETE", CoxEdgeAPIBase+"/scripts/"+id, nil)
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
