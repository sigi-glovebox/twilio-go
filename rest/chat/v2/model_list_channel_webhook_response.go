/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListChannelWebhookResponse struct for ListChannelWebhookResponse
type ListChannelWebhookResponse struct {
	Meta     ListCredentialResponseMeta           `json:"meta,omitempty"`
	Webhooks []ChatV2ServiceChannelChannelWebhook `json:"webhooks,omitempty"`
}
