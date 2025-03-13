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

// BulkexportsV1Day struct for BulkexportsV1Day
type BulkexportsV1Day struct {
	// The ISO 8601 format date of the resources in the file, for a UTC day
	Day *string `json:"day,omitempty"`
	// The size of the day's data file in bytes
	Size int `json:"size,omitempty"`
	// The ISO 8601 format date when resources is created
	CreateDate *string `json:"create_date,omitempty"`
	// The friendly name specified when creating the job
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
}
