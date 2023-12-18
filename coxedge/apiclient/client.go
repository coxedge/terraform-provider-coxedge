/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//const CoxEdgeAPIBase = "https://portal.coxedge.com/api/v2"
const CoxEdgeServiceCode = "edge-services"
const CoxEdgeBareMetalServiceCode = "baremetal-services"

//const CoxEdgeAPIBase = "https://cox.uat.cloudmc.io/api/v2"
//const CoxEdgeServiceCode = "stackpath-cox-uat"

type Client struct {
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) Client {
	return Client{
		HTTPClient: &http.Client{Timeout: 180 * time.Second},
		apiKey:     apiKey,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("MC-Api-Key", c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
