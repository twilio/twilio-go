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
	"github.com/twilio/twilio-go/client"
)

// ApiV2010Recording struct for ApiV2010Recording
type ApiV2010Recording struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used during the recording.
	ApiVersion *string `json:"api_version,omitempty"`
	// The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Recording resource is associated with. This will always refer to the parent leg of a two-leg call.
	CallSid *string `json:"call_sid,omitempty"`
	// The Conference SID that identifies the conference associated with the recording, if a conference recording.
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// The start time of the recording in GMT and in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format.
	StartTime *string `json:"start_time,omitempty"`
	// The length of the recording in seconds.
	Duration *string `json:"duration,omitempty"`
	// The unique string that that we created to identify the Recording resource.
	Sid *string `json:"sid,omitempty"`
	// The one-time cost of creating the recording in the `price_unit` currency.
	Price *string `json:"price,omitempty"`
	// The currency used in the `price` property. Example: `USD`.
	PriceUnit *string `json:"price_unit,omitempty"`
	Status    *string `json:"status,omitempty"`
	// The number of channels in the final recording file. Can be: `1` or `2`.
	Channels int     `json:"channels,omitempty"`
	Source   *string `json:"source,omitempty"`
	// The error code that describes why the recording is `absent`. The error code is described in our [Error Dictionary](https://www.twilio.com/docs/api/errors). This value is null if the recording `status` is not `absent`.
	ErrorCode *int `json:"error_code,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
	// How to decrypt the recording if it was encrypted using [Call Recording Encryption](https://www.twilio.com/docs/voice/tutorials/voice-recording-encryption) feature.
	EncryptionDetails *map[string]interface{} `json:"encryption_details,omitempty"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URL of the media file associated with this recording resource. When stored externally, this is the full URL location of the media file.
	MediaUrl *string `json:"media_url,omitempty"`
}
