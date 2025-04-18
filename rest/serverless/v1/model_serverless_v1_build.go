/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Serverless
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ServerlessV1Build struct for ServerlessV1Build
type ServerlessV1Build struct {
	// The unique string that we created to identify the Build resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Build resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Service that the Build resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	Status     *string `json:"status,omitempty"`
	// The list of Asset Version resource SIDs that are included in the Build.
	AssetVersions *[]map[string]interface{} `json:"asset_versions,omitempty"`
	// The list of Function Version resource SIDs that are included in the Build.
	FunctionVersions *[]map[string]interface{} `json:"function_versions,omitempty"`
	// A list of objects that describe the Dependencies included in the Build. Each object contains the `name` and `version` of the dependency.
	Dependencies *[]map[string]interface{} `json:"dependencies,omitempty"`
	Runtime      *string                   `json:"runtime,omitempty"`
	// The date and time in GMT when the Build resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the Build resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Build resource.
	Url   *string                 `json:"url,omitempty"`
	Links *map[string]interface{} `json:"links,omitempty"`
}
