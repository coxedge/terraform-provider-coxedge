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

type EnvironmentCreateRequest struct {
	EnvironmentName   string       `json:"name,omitempty"`
	Description       string       `json:"description,omitempty"`
	ServiceConnection IdOnlyHelper `json:"serviceConnection,omitempty"`
	Organization      IdOnlyHelper `json:"organization,omitempty"`
	Membership        string       `json:"membership,omitempty"`
	Roles             []Role       `json:"roles,omitempty"`
	//Roles             []struct {
	//	Name      string         `json:"name,omitempty"`
	//	Users     []IdOnlyHelper `json:"users,omitempty"`
	//	IsDefault bool           `json:"isDefault,omitempty"`
	//} `json:"roles,omitempty"`
}

type EnvironmentMembershipRequest struct {
	Membership string `json:"membership,omitempty"`
}

type EnvironmentMembersRequest struct {
	User IdOnlyHelper `json:"user,omitempty"`
	Role IdOnlyHelper `json:"role,omitempty"`
}

//GetEnvironments Get Environments in account
func (c *Client) GetEnvironments() ([]Environment, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/environments", nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedEnvironments
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetEnvironment Get Environment in account by id
func (c *Client) GetEnvironment(id string) (*Environment, error) {
	//Create the request
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/environments/"+id, nil)
	if err != nil {
		return nil, err
	}

	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	//Unmarshal, unwrap, and return
	var wrappedAPIStruct WrappedEnvironment
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//CreateEnvironment Create the Environment
func (c *Client) CreateEnvironment(newEnvironment EnvironmentCreateRequest) (*Environment, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newEnvironment)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/environments", bReader)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedEnvironment
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//UpdateEnvironment Update a Environment
func (c *Client) UpdateEnvironment(EnvironmentId string, newEnvironment EnvironmentCreateRequest) (*Environment, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newEnvironment)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT", CoxEdgeAPIBase+"/environments/"+EnvironmentId, bReader)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedEnvironment
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//UpdateEnvironmentMembership Update a Environment membership
func (c *Client) UpdateEnvironmentMembership(EnvironmentId string, newEnvironment EnvironmentMembershipRequest) (*Environment, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newEnvironment)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT", CoxEdgeAPIBase+"/environments/"+EnvironmentId+"/membership", bReader)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedEnvironment
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//UpdateEnvironmentMember Update a Environment members
func (c *Client) UpdateEnvironmentMember(EnvironmentId string, newEnvironment EnvironmentMembersRequest) (*Environment, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newEnvironment)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/environments/"+EnvironmentId+"/members", bReader)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedEnvironment
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//DeleteEnvironment Delete Environment in account by id
func (c *Client) DeleteEnvironment(id string) error {
	//Create the request
	request, err := http.NewRequest("DELETE", CoxEdgeAPIBase+"/environments/"+id, nil)
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
