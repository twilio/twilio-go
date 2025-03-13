/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
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

// VerifyV2Safelist struct for VerifyV2Safelist
type VerifyV2Safelist struct {
	// The unique string that we created to identify the SafeList resource.
	Sid *string `json:"sid,omitempty"`
	// The phone number in SafeList.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The absolute URL of the SafeList resource.
	Url *string `json:"url,omitempty"`
}
