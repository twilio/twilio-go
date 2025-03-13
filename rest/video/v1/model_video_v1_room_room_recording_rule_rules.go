/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
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

// VideoV1RoomRoomRecordingRuleRules struct for VideoV1RoomRoomRecordingRuleRules
type VideoV1RoomRoomRecordingRuleRules struct {
	Type      string `json:"type,omitempty"`
	All       bool   `json:"all,omitempty"`
	Publisher string `json:"publisher,omitempty"`
	Track     string `json:"track,omitempty"`
	Kind      string `json:"kind,omitempty"`
}
