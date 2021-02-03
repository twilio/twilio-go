/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// EventsV1Sink struct for EventsV1Sink
type EventsV1Sink struct {
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Description string `json:"Description,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SinkConfiguration map[string]interface{} `json:"SinkConfiguration,omitempty"`
	SinkType SinkType `json:"SinkType,omitempty"`
	Status Status `json:"Status,omitempty"`
	Url string `json:"Url,omitempty"`
}
