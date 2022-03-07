package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//GetOrganizations Get organizations in account
func (c *Client) GetOrganizations() ([]Organization, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/organizations", nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedOrganizations
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetOrganization Get organizations in account by id
func (c *Client) GetOrganization(id string) (*Organization, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/organizations/"+id, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)

	fmt.Println(string(respBytes))
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedOrganization
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}
