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

// Creates an evaluation for a bundle
func (c *ApiService) CreateEvaluation(BundleSid string) (*NumbersV2RegulatoryComplianceBundleEvaluation, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceBundleEvaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch specific Evaluation Instance.
func (c *ApiService) FetchEvaluation(BundleSid string, Sid string) (*NumbersV2RegulatoryComplianceBundleEvaluation, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceBundleEvaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEvaluation'
type ListEvaluationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListEvaluationParams) SetPageSize(PageSize int) *ListEvaluationParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Evaluations associated to the Bundle resource.
func (c *ApiService) ListEvaluation(BundleSid string, params *ListEvaluationParams) (*ListEvaluationResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"
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

	ps := &ListEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
