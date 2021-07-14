/*
 * Twilio - Proxy
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

// Optional parameters for the method 'CreateMessageInteraction'
type CreateMessageInteractionParams struct {
	// The message to send to the participant
	Body *string `json:"Body,omitempty"`
	// Reserved. Not currently supported.
	MediaUrl *[]string `json:"MediaUrl,omitempty"`
}

func (params *CreateMessageInteractionParams) SetBody(Body string) *CreateMessageInteractionParams {
	params.Body = &Body
	return params
}
func (params *CreateMessageInteractionParams) SetMediaUrl(MediaUrl []string) *CreateMessageInteractionParams {
	params.MediaUrl = &MediaUrl
	return params
}

// Create a new message Interaction to send directly from your system to one [Participant](https://www.twilio.com/docs/proxy/api/participant).  The &#x60;inbound&#x60; properties for the Interaction will always be empty.
func (c *ApiService) CreateMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, params *CreateMessageInteractionParams) (*ProxyV1ServiceSessionParticipantMessageInteraction, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.MediaUrl != nil {
		for _, item := range *params.MediaUrl {
			data.Add("MediaUrl", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServiceSessionParticipantMessageInteraction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, Sid string) (*ProxyV1ServiceSessionParticipantMessageInteraction, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServiceSessionParticipantMessageInteraction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessageInteraction'
type ListMessageInteractionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListMessageInteractionParams) SetPageSize(PageSize int) *ListMessageInteractionParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, params *ListMessageInteractionParams) (*ListMessageInteractionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

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

	ps := &ListMessageInteractionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
