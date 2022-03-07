package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//GetImages Get images in account
func (c *Client) GetImages(environmentName string) ([]Image, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/images", nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedImages
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetImage Get images in account by id
func (c *Client) GetImage(environmentName string, id string) (*Image, error) {
	request, err := http.NewRequest("GET", CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/images/"+id, nil)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)

	fmt.Println(string(respBytes))
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedImage
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}
