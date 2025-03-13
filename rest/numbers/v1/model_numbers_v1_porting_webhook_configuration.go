/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// NumbersV1PortingWebhookConfiguration struct for NumbersV1PortingWebhookConfiguration
type NumbersV1PortingWebhookConfiguration struct {
	// The URL of the webhook configuration request
	Url *string `json:"url,omitempty"`
	// The complete webhook url that will be called when a notification event for port in request or port in phone number happens
	PortInTargetUrl *string `json:"port_in_target_url,omitempty"`
	// The complete webhook url that will be called when a notification event for a port out phone number happens.
	PortOutTargetUrl *string `json:"port_out_target_url,omitempty"`
	// A list to filter what notification events to receive for this account and its sub accounts. If it is an empty list, then it means that there are no filters for the notifications events to send in each webhook and all events will get sent.
	NotificationsOf *[]string `json:"notifications_of,omitempty"`
}
