/*
 * Twilio - Wireless
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

// Optional parameters for the method 'ListDataSession'
type ListDataSessionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListDataSessionParams) SetPageSize(PageSize int) *ListDataSessionParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListDataSession(SimSid string, params *ListDataSessionParams) (*ListDataSessionResponse, error) {
	path := "/v1/Sims/{SimSid}/DataSessions"
	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

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

	ps := &ListDataSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
