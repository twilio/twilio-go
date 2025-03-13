/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
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
	"time"
)

// ProxyV1MessageInteraction struct for ProxyV1MessageInteraction
type ProxyV1MessageInteraction struct {
	// The unique string that we created to identify the MessageInteraction resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
	SessionSid *string `json:"session_sid,omitempty"`
	// The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the MessageInteraction resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// A JSON string that includes the message body sent to the participant. (e.g. `{\"body\": \"hello\"}`)
	Data *string `json:"data,omitempty"`
	Type *string `json:"type,omitempty"`
	// The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// Always empty for created Message Interactions.
	InboundParticipantSid *string `json:"inbound_participant_sid,omitempty"`
	// Always empty for created Message Interactions.
	InboundResourceSid    *string `json:"inbound_resource_sid,omitempty"`
	InboundResourceStatus *string `json:"inbound_resource_status,omitempty"`
	// Always empty for created Message Interactions.
	InboundResourceType *string `json:"inbound_resource_type,omitempty"`
	// Always empty for created Message Interactions.
	InboundResourceUrl *string `json:"inbound_resource_url,omitempty"`
	// The SID of the outbound [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.
	OutboundParticipantSid *string `json:"outbound_participant_sid,omitempty"`
	// The SID of the outbound [Message](https://www.twilio.com/docs/sms/api/message-resource) resource.
	OutboundResourceSid    *string `json:"outbound_resource_sid,omitempty"`
	OutboundResourceStatus *string `json:"outbound_resource_status,omitempty"`
	// The outbound resource type. This value is always `Message`.
	OutboundResourceType *string `json:"outbound_resource_type,omitempty"`
	// The URL of the Twilio message resource.
	OutboundResourceUrl *string `json:"outbound_resource_url,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the MessageInteraction resource.
	Url *string `json:"url,omitempty"`
}
