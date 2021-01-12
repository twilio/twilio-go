/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// EventsV1SchemaVersion struct for EventsV1SchemaVersion
type EventsV1SchemaVersion struct {
	DateCreated time.Time `json:"date_created,omitempty"`
	Id string `json:"id,omitempty"`
	Raw string `json:"raw,omitempty"`
	SchemaVersion int32 `json:"schema_version,omitempty"`
	Url string `json:"url,omitempty"`
}
