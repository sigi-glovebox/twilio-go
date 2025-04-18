/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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

// Optional parameters for the method 'ListUsageRecordToday'
type ListUsageRecordTodayParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
	Category *string `json:"Category,omitempty"`
	// Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date.
	StartDate *string `json:"StartDate,omitempty"`
	// Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date.
	EndDate *string `json:"EndDate,omitempty"`
	// Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account.
	IncludeSubaccounts *bool `json:"IncludeSubaccounts,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsageRecordTodayParams) SetPathAccountSid(PathAccountSid string) *ListUsageRecordTodayParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListUsageRecordTodayParams) SetCategory(Category string) *ListUsageRecordTodayParams {
	params.Category = &Category
	return params
}
func (params *ListUsageRecordTodayParams) SetStartDate(StartDate string) *ListUsageRecordTodayParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListUsageRecordTodayParams) SetEndDate(EndDate string) *ListUsageRecordTodayParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListUsageRecordTodayParams) SetIncludeSubaccounts(IncludeSubaccounts bool) *ListUsageRecordTodayParams {
	params.IncludeSubaccounts = &IncludeSubaccounts
	return params
}
func (params *ListUsageRecordTodayParams) SetPageSize(PageSize int) *ListUsageRecordTodayParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsageRecordTodayParams) SetLimit(Limit int) *ListUsageRecordTodayParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsageRecordToday records from the API. Request is executed immediately.
func (c *ApiService) PageUsageRecordToday(params *ListUsageRecordTodayParams, pageToken, pageNumber string) (*ListUsageRecordTodayResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Records/Today.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Category != nil {
		data.Set("Category", fmt.Sprint(*params.Category))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.IncludeSubaccounts != nil {
		data.Set("IncludeSubaccounts", fmt.Sprint(*params.IncludeSubaccounts))
	}
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

	ps := &ListUsageRecordTodayResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsageRecordToday records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsageRecordToday(params *ListUsageRecordTodayParams) ([]ApiV2010UsageRecordToday, error) {
	response, errors := c.StreamUsageRecordToday(params)

	records := make([]ApiV2010UsageRecordToday, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams UsageRecordToday records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsageRecordToday(params *ListUsageRecordTodayParams) (chan ApiV2010UsageRecordToday, chan error) {
	if params == nil {
		params = &ListUsageRecordTodayParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010UsageRecordToday, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageUsageRecordToday(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamUsageRecordToday(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamUsageRecordToday(response *ListUsageRecordTodayResponse, params *ListUsageRecordTodayParams, recordChannel chan ApiV2010UsageRecordToday, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.UsageRecords
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListUsageRecordTodayResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListUsageRecordTodayResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListUsageRecordTodayResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordTodayResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
