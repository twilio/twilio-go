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
	"time"
)

// NumbersV1BulkEligibility struct for NumbersV1BulkEligibility
type NumbersV1BulkEligibility struct {
	// The SID of the bulk eligibility check that you want to know about.
	RequestId *string `json:"request_id,omitempty"`
	// This is the url of the request that you're trying to reach out to locate the resource.
	Url *string `json:"url,omitempty"`
	// The result set that contains the eligibility check response for each requested number, each result has at least the following attributes:  phone_number: The requested phone number ,hosting_account_sid: The account sid where the phone number will be hosted, country: Phone number’s country, eligibility_status: Indicates the eligibility status of the PN (Eligible/Ineligible), eligibility_sub_status: Indicates the sub status of the eligibility , ineligibility_reason: Reason for number's ineligibility (if applicable), next_step: Suggested next step in the hosting process based on the eligibility status.
	Results *[]map[string]interface{} `json:"results,omitempty"`
	// This is the string that you assigned as a friendly name for describing the eligibility check request.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// This is the status of the bulk eligibility check request. (Example: COMPLETE, IN_PROGRESS)
	Status        *string    `json:"status,omitempty"`
	DateCreated   *time.Time `json:"date_created,omitempty"`
	DateCompleted *time.Time `json:"date_completed,omitempty"`
}
