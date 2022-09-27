/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
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

// InsightsV1CallSummaries struct for InsightsV1CallSummaries
type InsightsV1CallSummaries struct {
	AccountSid      *string      `json:"account_sid,omitempty"`
	CallSid         *string      `json:"call_sid,omitempty"`
	CallType        *string      `json:"call_type,omitempty"`
	CallState       *string      `json:"call_state,omitempty"`
	ProcessingState *string      `json:"processing_state,omitempty"`
	CreatedTime     *time.Time   `json:"created_time,omitempty"`
	StartTime       *time.Time   `json:"start_time,omitempty"`
	EndTime         *time.Time   `json:"end_time,omitempty"`
	Duration        *int         `json:"duration,omitempty"`
	ConnectDuration *int         `json:"connect_duration,omitempty"`
	From            *interface{} `json:"from,omitempty"`
	To              *interface{} `json:"to,omitempty"`
	CarrierEdge     *interface{} `json:"carrier_edge,omitempty"`
	ClientEdge      *interface{} `json:"client_edge,omitempty"`
	SdkEdge         *interface{} `json:"sdk_edge,omitempty"`
	SipEdge         *interface{} `json:"sip_edge,omitempty"`
	Tags            *[]string    `json:"tags,omitempty"`
	Url             *string      `json:"url,omitempty"`
	Attributes      *interface{} `json:"attributes,omitempty"`
	Properties      *interface{} `json:"properties,omitempty"`
	Trust           *interface{} `json:"trust,omitempty"`
}
