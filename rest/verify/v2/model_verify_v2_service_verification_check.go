/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VerifyV2ServiceVerificationCheck struct for VerifyV2ServiceVerificationCheck
type VerifyV2ServiceVerificationCheck struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The amount of the associated PSD2 compliant transaction.
	Amount *string `json:"amount,omitempty"`
	// The verification method to use
	Channel *string `json:"channel,omitempty"`
	// The ISO 8601 date and time in GMT when the Verification Check resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Verification Check resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The payee of the associated PSD2 compliant transaction
	Payee *string `json:"payee,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the verification resource
	Status *string `json:"status,omitempty"`
	// The phone number or email being verified
	To *string `json:"to,omitempty"`
	// Whether the verification was successful
	Valid *bool `json:"valid,omitempty"`
}
