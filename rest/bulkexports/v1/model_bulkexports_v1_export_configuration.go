/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Bulkexports
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

// BulkexportsV1ExportConfiguration struct for BulkexportsV1ExportConfiguration
type BulkexportsV1ExportConfiguration struct {
	// If true, Twilio will automatically generate every day's file when the day is over.
	Enabled *bool `json:"enabled,omitempty"`
	// Stores the URL destination for the method specified in webhook_method.
	WebhookUrl *string `json:"webhook_url,omitempty"`
	// Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url
	WebhookMethod *string `json:"webhook_method,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
