/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
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

// IntelligenceV2Service struct for IntelligenceV2Service
type IntelligenceV2Service struct {
	// The unique SID identifier of the Account the Service belongs to.
	AccountSid *string `json:"account_sid,omitempty"`
	// Instructs the Speech Recognition service to automatically redact PII from all transcripts made on this service.
	AutoRedaction *bool `json:"auto_redaction,omitempty"`
	// Instructs the Speech Recognition service to automatically redact PII from all transcripts media made on this service. The auto_redaction flag must be enabled, results in error otherwise.
	MediaRedaction *bool `json:"media_redaction,omitempty"`
	// Instructs the Speech Recognition service to automatically transcribe all recordings made on the account.
	AutoTranscribe *bool `json:"auto_transcribe,omitempty"`
	// Data logging allows Twilio to improve the quality of the speech recognition & language understanding services through using customer data to refine, fine tune and evaluate machine learning models. Note: Data logging cannot be activated via API, only via www.twilio.com, as it requires additional consent.
	DataLogging *bool `json:"data_logging,omitempty"`
	// The date that this Service was created, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Service was updated, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A human readable description of this resource, up to 64 characters.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The language code set during Service creation determines the Transcription language for all call recordings processed by that Service. The default is en-US if no language code is set. A Service can only support one language code, and it cannot be updated once it's set.
	LanguageCode *string `json:"language_code,omitempty"`
	// A 34 character string that uniquely identifies this Service.
	Sid *string `json:"sid,omitempty"`
	// Provides a unique and addressable name to be assigned to this Service, assigned by the developer, to be optionally used in addition to SID.
	UniqueName *string `json:"unique_name,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// The URL Twilio will request when executing the Webhook.
	WebhookUrl        *string `json:"webhook_url,omitempty"`
	WebhookHttpMethod *string `json:"webhook_http_method,omitempty"`
	// Operator sids attached to this service, read only
	ReadOnlyAttachedOperatorSids *[]string `json:"read_only_attached_operator_sids,omitempty"`
	// The version number of this Service.
	Version int `json:"version,omitempty"`
}
