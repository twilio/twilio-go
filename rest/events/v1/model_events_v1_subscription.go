/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Events
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

// EventsV1Subscription struct for EventsV1Subscription
type EventsV1Subscription struct {
	// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
	// A 34 character string that uniquely identifies this Subscription.
	Sid *string `json:"sid,omitempty"`
	// The date that this Subscription was created, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Subscription was updated, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A human readable description for the Subscription
	Description *string `json:"description,omitempty"`
	// The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
	SinkSid *string `json:"sink_sid,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// Contains a dictionary of URL links to nested resources of this Subscription.
	Links *map[string]interface{} `json:"links,omitempty"`
	// Receive events from all children accounts in the parent account subscription.
	ReceiveEventsFromSubaccounts *bool `json:"receive_events_from_subaccounts,omitempty"`
}
