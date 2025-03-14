/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// MessagingV1BrandRegistrations struct for MessagingV1BrandRegistrations
type MessagingV1BrandRegistrations struct {
	// The unique string to identify Brand Registration.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Brand Registration resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// A2P Messaging Profile Bundle BundleSid.
	CustomerProfileBundleSid *string `json:"customer_profile_bundle_sid,omitempty"`
	// A2P Messaging Profile Bundle BundleSid.
	A2pProfileBundleSid *string `json:"a2p_profile_bundle_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Type of brand. One of: \"STANDARD\", \"SOLE_PROPRIETOR\". SOLE_PROPRIETOR is for the low volume, SOLE_PROPRIETOR campaign use case. There can only be one SOLE_PROPRIETOR campaign created per SOLE_PROPRIETOR brand. STANDARD is for all other campaign use cases. Multiple campaign use cases can be created per STANDARD brand.
	BrandType *string `json:"brand_type,omitempty"`
	Status    *string `json:"status,omitempty"`
	// Campaign Registry (TCR) Brand ID. Assigned only after successful brand registration.
	TcrId *string `json:"tcr_id,omitempty"`
	// DEPRECATED. A reason why brand registration has failed. Only applicable when status is FAILED.
	FailureReason *string `json:"failure_reason,omitempty"`
	// A list of errors that occurred during the brand registration process.
	Errors *[]map[string]interface{} `json:"errors,omitempty"`
	// The absolute URL of the Brand Registration resource.
	Url *string `json:"url,omitempty"`
	// The secondary vetting score if it was done. Otherwise, it will be the brand score if it's returned from TCR. It may be null if no score is available.
	BrandScore *int `json:"brand_score,omitempty"`
	// DEPRECATED. Feedback on how to improve brand score
	BrandFeedback  *[]string `json:"brand_feedback,omitempty"`
	IdentityStatus *string   `json:"identity_status,omitempty"`
	// Publicly traded company identified in the Russell 3000 Index
	Russell3000 *bool `json:"russell_3000,omitempty"`
	// Identified as a government entity
	GovernmentEntity *bool `json:"government_entity,omitempty"`
	// Nonprofit organization tax-exempt status per section 501 of the U.S. tax code.
	TaxExemptStatus *string `json:"tax_exempt_status,omitempty"`
	// A flag to disable automatic secondary vetting for brands which it would otherwise be done.
	SkipAutomaticSecVet *bool `json:"skip_automatic_sec_vet,omitempty"`
	// A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided.
	Mock  *bool                   `json:"mock,omitempty"`
	Links *map[string]interface{} `json:"links,omitempty"`
}
