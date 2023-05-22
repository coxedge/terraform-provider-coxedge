/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package apiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type PredefinedEdgeLogicRequest struct {
	ForceWwwEnabled           *bool    `json:"forceWwwEnabled,omitempty"`
	RobotsTxtEnabled          *bool    `json:"robotsTxtEnabled,omitempty"`
	RobotsTxtFile             string   `json:"robotsTxtFile,omitempty"`
	PseudoStreamingEnabled    *bool    `json:"pseudoStreamingEnabled,omitempty"`
	ReferrerProtectionEnabled *bool    `json:"referrerProtectionEnabled,omitempty"`
	AllowEmptyReferrer        *bool    `json:"allowEmptyReferrer,omitempty"`
	ReferrerList              []string `json:"referrerList,omitempty"`
}

type DeliveryRuleRequest struct {
	Id        string             `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Condition []ConditionRequest `json:"conditions,omitempty"`
	Action    []ActionRequest    `json:"actions,omitempty"`
	ScopeId   string             `json:"scopeId"`
	SiteId    string             `json:"siteId"`
	Slug      string             `json:"slug"`
	StackId   string             `json:"stackId"`
}

type ConditionRequest struct {
	Trigger     string   `json:"trigger,omitempty"`
	Operator    string   `json:"operator,omitempty"`
	HTTPMethods []string `json:"httpMethods,omitempty"`
	Target      string   `json:"target,omitempty"`
}

type ActionRequest struct {
	ActionType             string          `json:"actionType,omitempty"`
	ResponseHeaders        []HeaderRequest `json:"responseHeaders"`
	OriginHeaders          []HeaderRequest `json:"originHeaders"`
	CDNHeaders             []HeaderRequest `json:"cdnHeaders"`
	CacheTtl               int             `json:"cacheTtl,omitempty"`
	RedirectUrl            string          `json:"redirectUrl,omitempty"`
	HeaderPattern          string          `json:"headerPattern,omitempty"`
	Passphrase             string          `json:"passphrase,omitempty"`
	PassphraseField        string          `json:"passphraseField,omitempty"`
	MD5TokenField          string          `json:"md5TokenField,omitempty"`
	TTLField               string          `json:"ttlField,omitempty"`
	IPAddressFilter        string          `json:"ipAddressFilter,omitempty"`
	URLSignaturePathLength string          `json:"urlSignaturePathLength,omitempty"`
}

type HeaderRequest struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

//GetPredefinedEdgeLogics Get predefined edge logic of sites by site Id
func (c *Client) GetPredefinedEdgeLogics(environmentName string, organizationId string, siteId string) (*EdgeLogic, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/predefinededgerules/"+siteId+"?org_id="+organizationId,
		nil,
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedEdgeLogic
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//UpdatePredefinedEdgeLogic Update predefined edge logic of sites by site Id
func (c *Client) UpdatePredefinedEdgeLogic(predefinedEdgeLogicRequest PredefinedEdgeLogicRequest, environmentName string, organizationId string, siteId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(predefinedEdgeLogicRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("PATCH",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/predefinededgerules/"+siteId+"?org_id="+organizationId,
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

//GetDeliveryRules Get delivery rules from edge logic of sites by site Id
func (c *Client) GetDeliveryRules(environmentName string, organizationId string, siteId string) ([]DeliveryRule, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliveryrules?siteId="+siteId+"&org_id="+organizationId,
		nil,
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedDeliveryRules
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return wrappedAPIStruct.Data, nil
}

//GetDeliveryRule Get delivery rule by Id from edge logic of sites by site Id
func (c *Client) GetDeliveryRule(environmentName string, organizationId string, deliveryRuleId string, siteId string) (*DeliveryRule, error) {
	request, err := http.NewRequest("GET",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliveryrules/"+deliveryRuleId+"?siteId="+siteId+"&org_id="+organizationId,
		nil,
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}

	var wrappedAPIStruct WrappedDeliveryRule
	err = json.Unmarshal(respBytes, &wrappedAPIStruct)
	if err != nil {
		return nil, err
	}
	return &wrappedAPIStruct.Data, nil
}

//AddDeliveryRule add delivery rule to edge logic of sites by site Id
func (c *Client) AddDeliveryRule(deliveryRuleRequest DeliveryRuleRequest, environmentName string, organizationId string, siteId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(deliveryRuleRequest)
	if err != nil {
		return nil, err
	}

	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliveryrules?siteId="+siteId+"&org_id="+organizationId,
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

//UpdateDeliveryRule edit delivery rule by rule Id in edge logic of sites by site Id
func (c *Client) UpdateDeliveryRule(ctx context.Context, deliveryRuleRequest DeliveryRuleRequest, environmentName string, organizationId string, deliveryRuleId string, siteId string) (*TaskStatusResponse, error) {
	//Marshal the request
	jsonBytes, err := json.Marshal(deliveryRuleRequest)
	if err != nil {
		return nil, err
	}
	//Wrap bytes in reader
	bReader := bytes.NewReader(jsonBytes)
	//Create the request
	request, err := http.NewRequest("POST",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliveryrules/"+deliveryRuleId+"?siteId="+siteId+"&org_id="+organizationId+"&operation=edit",
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

//DeleteDeliveryRule delete delivery rule by rule Id in edge logic of sites by site Id
func (c *Client) DeleteDeliveryRule(environmentName string, organizationId string, deliveryRuleId string, siteId string) (*TaskStatusResponse, error) {

	//Create the request
	request, err := http.NewRequest("DELETE",
		CoxEdgeAPIBase+"/services/"+CoxEdgeServiceCode+"/"+environmentName+"/deliveryrules/"+deliveryRuleId+"?siteId="+siteId+"&org_id="+organizationId,
		nil,
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
