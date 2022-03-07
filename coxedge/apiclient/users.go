package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type UserCreateRequest struct {
	UserName       string         `json:"userName"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Email          string         `json:"email"`
	OrganizationId IdOnlyHelper   `json:"organization,omitempty"`
	Roles          []IdOnlyHelper `json:"roles,omitempty"`
}

//GetUsers Get users in account
func (c *Client) GetUsers() ([]User, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/users", nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedUsers
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetUser Get user in account by id
func (c *Client) GetUser(id string) (*User, error) {
	//Create the request
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/users/"+id, nil)
	if err != nil {
		return nil, err
	}

	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	//Unmarshal, unwrap, and return
	var wrappedAPIStruct WrappedUser
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//CreateUser Create the user
func (c *Client) CreateUser(newUser UserCreateRequest) (*User, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST", CoxEdgeAPIBase+"/users", bReader)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedUser
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//UpdateUser Update a user
func (c *Client) UpdateUser(userId string, newUser UserCreateRequest) (*User, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PUT", CoxEdgeAPIBase+"/users/"+userId, bReader)
	request.Header.Set("Content-Type", "application/json")
	//Execute request
	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	//Return struct
	var wrappedAPIStruct WrappedUser
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//DeleteUser Delete user in account by id
func (c *Client) DeleteUser(id string) error {
	//Create the request
	request, err := http.NewRequest("DELETE", CoxEdgeAPIBase+"/users/"+id, nil)
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

//UnlockUser Unlock user in account by id
func (c *Client) UnlockUser(id string) error {
	//Create the request
	request, err := http.NewRequest("DELETE", CoxEdgeAPIBase+"/users/"+id+"/unlock", nil)
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
