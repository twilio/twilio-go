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

// BulkexportsV1Export struct for BulkexportsV1Export
type BulkexportsV1Export struct {
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// Contains a dictionary of URL links to nested resources of this Export.
	Links *map[string]interface{} `json:"links,omitempty"`
}
