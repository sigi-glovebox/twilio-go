/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Marketplace
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Fetch an instance of an Extension for the Available Add-on.
func (c *ApiService) FetchMarketplaceAvailableAddOnExtension(AvailableAddOnSid string, Sid string) (*MarketplaceAvailableAddOnAvailableAddOnExtension, error) {
	path := "/v1/AvailableAddOns/{AvailableAddOnSid}/Extensions/{Sid}"
	path = strings.Replace(path, "{"+"AvailableAddOnSid"+"}", AvailableAddOnSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MarketplaceAvailableAddOnAvailableAddOnExtension{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMarketplaceAvailableAddOnExtension'
type ListMarketplaceAvailableAddOnExtensionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMarketplaceAvailableAddOnExtensionParams) SetPageSize(PageSize int) *ListMarketplaceAvailableAddOnExtensionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMarketplaceAvailableAddOnExtensionParams) SetLimit(Limit int) *ListMarketplaceAvailableAddOnExtensionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of MarketplaceAvailableAddOnExtension records from the API. Request is executed immediately.
func (c *ApiService) PageMarketplaceAvailableAddOnExtension(AvailableAddOnSid string, params *ListMarketplaceAvailableAddOnExtensionParams, pageToken, pageNumber string) (*ListMarketplaceAvailableAddOnExtensionResponse, error) {
	path := "/v1/AvailableAddOns/{AvailableAddOnSid}/Extensions"

	path = strings.Replace(path, "{"+"AvailableAddOnSid"+"}", AvailableAddOnSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMarketplaceAvailableAddOnExtensionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists MarketplaceAvailableAddOnExtension records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMarketplaceAvailableAddOnExtension(AvailableAddOnSid string, params *ListMarketplaceAvailableAddOnExtensionParams) ([]MarketplaceAvailableAddOnAvailableAddOnExtension, error) {
	response, errors := c.StreamMarketplaceAvailableAddOnExtension(AvailableAddOnSid, params)

	records := make([]MarketplaceAvailableAddOnAvailableAddOnExtension, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams MarketplaceAvailableAddOnExtension records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMarketplaceAvailableAddOnExtension(AvailableAddOnSid string, params *ListMarketplaceAvailableAddOnExtensionParams) (chan MarketplaceAvailableAddOnAvailableAddOnExtension, chan error) {
	if params == nil {
		params = &ListMarketplaceAvailableAddOnExtensionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MarketplaceAvailableAddOnAvailableAddOnExtension, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMarketplaceAvailableAddOnExtension(AvailableAddOnSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMarketplaceAvailableAddOnExtension(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamMarketplaceAvailableAddOnExtension(response *ListMarketplaceAvailableAddOnExtensionResponse, params *ListMarketplaceAvailableAddOnExtensionParams, recordChannel chan MarketplaceAvailableAddOnAvailableAddOnExtension, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Extensions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListMarketplaceAvailableAddOnExtensionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListMarketplaceAvailableAddOnExtensionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListMarketplaceAvailableAddOnExtensionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMarketplaceAvailableAddOnExtensionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
