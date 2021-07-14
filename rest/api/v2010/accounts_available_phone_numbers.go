/*
 * Twilio - Api
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

// Optional parameters for the method 'FetchAvailablePhoneNumberCountry'
type FetchAvailablePhoneNumberCountryParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchAvailablePhoneNumberCountryParams) SetPathAccountSid(PathAccountSid string) *FetchAvailablePhoneNumberCountryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchAvailablePhoneNumberCountry(CountryCode string, params *FetchAvailablePhoneNumberCountryParams) (*ApiV2010AccountAvailablePhoneNumberCountry, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CountryCode"+"}", CountryCode, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountAvailablePhoneNumberCountry{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAvailablePhoneNumberCountry'
type ListAvailablePhoneNumberCountryParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListAvailablePhoneNumberCountryParams) SetPathAccountSid(PathAccountSid string) *ListAvailablePhoneNumberCountryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAvailablePhoneNumberCountryParams) SetPageSize(PageSize int) *ListAvailablePhoneNumberCountryParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListAvailablePhoneNumberCountry(params *ListAvailablePhoneNumberCountryParams) (*ListAvailablePhoneNumberCountryResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListAvailablePhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
