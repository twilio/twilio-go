/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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
	"time"
)

// SupersimV1SettingsUpdate struct for SupersimV1SettingsUpdate
type SupersimV1SettingsUpdate struct {
	// The unique identifier of this Settings Update.
	Sid *string `json:"sid,omitempty"`
	// The [ICCID](https://en.wikipedia.org/wiki/SIM_card#ICCID) associated with the SIM.
	Iccid *string `json:"iccid,omitempty"`
	// The SID of the Super SIM to which this Settings Update was applied.
	SimSid *string `json:"sim_sid,omitempty"`
	Status *string `json:"status,omitempty"`
	// Array containing the different Settings Packages that will be applied to the SIM after the update completes. Each object within the array indicates the name and the version of the Settings Package that will be on the SIM if the update is successful.
	Packages *[]map[string]interface{} `json:"packages,omitempty"`
	// The time, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, when the update successfully completed and the new settings were applied to the SIM.
	DateCompleted *time.Time `json:"date_completed,omitempty"`
	// The date that this Settings Update was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Settings Update was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}
