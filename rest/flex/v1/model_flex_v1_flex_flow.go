/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FlexV1FlexFlow struct for FlexV1FlexFlow
type FlexV1FlexFlow struct {
	AccountSid      *string                  `json:"account_sid,omitempty"`
	ChannelType     *FlexFlowChannelType     `json:"channel_type,omitempty"`
	ChatServiceSid  *string                  `json:"chat_service_sid,omitempty"`
	ContactIdentity *string                  `json:"contact_identity,omitempty"`
	DateCreated     *time.Time               `json:"date_created,omitempty"`
	DateUpdated     *time.Time               `json:"date_updated,omitempty"`
	Enabled         *bool                    `json:"enabled,omitempty"`
	FriendlyName    *string                  `json:"friendly_name,omitempty"`
	Integration     *map[string]interface{}  `json:"integration,omitempty"`
	IntegrationType *FlexFlowIntegrationType `json:"integration_type,omitempty"`
	JanitorEnabled  *bool                    `json:"janitor_enabled,omitempty"`
	LongLived       *bool                    `json:"long_lived,omitempty"`
	Sid             *string                  `json:"sid,omitempty"`
	Url             *string                  `json:"url,omitempty"`
}
