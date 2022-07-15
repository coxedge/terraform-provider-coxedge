package apiclient

import (
	"encoding/json"
	"net/http"
)

//GetRoles Get organizations in account
func (c *Client) GetRoles() ([]Roles, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/roles", nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedRolesData
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}
