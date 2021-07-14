/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateItemAssignment'
type CreateItemAssignmentParams struct {
	// The SID of an object bag that holds information of the different items.
	ObjectSid *string `json:"ObjectSid,omitempty"`
}

func (params *CreateItemAssignmentParams) SetObjectSid(ObjectSid string) *CreateItemAssignmentParams {
	params.ObjectSid = &ObjectSid
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateItemAssignment(BundleSid string, params *CreateItemAssignmentParams) (*NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ObjectSid != nil {
		data.Set("ObjectSid", *params.ObjectSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceBundleItemAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteItemAssignment(BundleSid string, Sid string) error {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchItemAssignment(BundleSid string, Sid string) (*NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceBundleItemAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListItemAssignment'
type ListItemAssignmentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListItemAssignmentParams) SetPageSize(PageSize int) *ListItemAssignmentParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Assigned Items for an account.
func (c *ApiService) ListItemAssignment(BundleSid string, params *ListItemAssignmentParams) (*ListItemAssignmentResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListItemAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
