/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
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

// ChatV2Message struct for ChatV2Message
type ChatV2Message struct {
	// The unique string that we created to identify the Message resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The JSON string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes *string `json:"attributes,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the Message resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Channel](https://www.twilio.com/docs/chat/channels) that the message was sent to.
	To *string `json:"to,omitempty"`
	// The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource belongs to.
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable.
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// Whether the message has been edited since it was created.
	WasEdited *bool `json:"was_edited,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the message's author. The default value is `system`.
	From *string `json:"from,omitempty"`
	// The content of the message.
	Body *string `json:"body,omitempty"`
	// The index of the message within the [Channel](https://www.twilio.com/docs/chat/channels). Indices may skip numbers, but will always be in order of when the message was received.
	Index int `json:"index,omitempty"`
	// The Message type. Can be: `text` or `media`.
	Type *string `json:"type,omitempty"`
	// An object that describes the Message's media, if the message contains media. The object contains these fields: `content_type` with the MIME type of the media, `filename` with the name of the media, `sid` with the SID of the Media resource, and `size` with the media object's file size in bytes. If the Message has no media, this value is `null`.
	Media *interface{} `json:"media,omitempty"`
	// The absolute URL of the Message resource.
	Url *string `json:"url,omitempty"`
}
