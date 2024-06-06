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

	"github.com/twilio/twilio-go/client"
)

// MarketplaceInstalledAddOnBillingUsageResponseBillableItems struct for MarketplaceInstalledAddOnBillingUsageResponseBillableItems
type MarketplaceInstalledAddOnBillingUsageResponseBillableItems struct {
	//
	Quantity float32 `json:"quantity,omitempty"`
	//
	Sid string `json:"sid,omitempty"`
	// Whether this billable item was successfully submitted for billing.
	Submitted bool `json:"submitted,omitempty"`
}

func (response *MarketplaceInstalledAddOnBillingUsageResponseBillableItems) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		Quantity  interface{} `json:"quantity"`
		Sid       string      `json:"sid"`
		Submitted bool        `json:"submitted"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = MarketplaceInstalledAddOnBillingUsageResponseBillableItems{
		Sid:       raw.Sid,
		Submitted: raw.Submitted,
	}

	responseQuantity, err := client.UnmarshalFloat32(&raw.Quantity)
	if err != nil {
		return err
	}
	response.Quantity = *responseQuantity

	return
}
