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
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)
// MessagingV1TollfreeVerification struct for MessagingV1TollfreeVerification
type MessagingV1TollfreeVerification struct {
		// The unique string to identify Tollfree Verification.
	Sid *string `json:"sid,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Tollfree Verification resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// Customer's Profile Bundle BundleSid.
	CustomerProfileSid *string `json:"customer_profile_sid,omitempty"`
		// Tollfree TrustProduct Bundle BundleSid.
	TrustProductSid *string `json:"trust_product_sid,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The SID of the Regulated Item.
	RegulatedItemSid *string `json:"regulated_item_sid,omitempty"`
		// The name of the business or organization using the Tollfree number.
	BusinessName *string `json:"business_name,omitempty"`
		// The address of the business or organization using the Tollfree number.
	BusinessStreetAddress *string `json:"business_street_address,omitempty"`
		// The address of the business or organization using the Tollfree number.
	BusinessStreetAddress2 *string `json:"business_street_address2,omitempty"`
		// The city of the business or organization using the Tollfree number.
	BusinessCity *string `json:"business_city,omitempty"`
		// The state/province/region of the business or organization using the Tollfree number.
	BusinessStateProvinceRegion *string `json:"business_state_province_region,omitempty"`
		// The postal code of the business or organization using the Tollfree number.
	BusinessPostalCode *string `json:"business_postal_code,omitempty"`
		// The country of the business or organization using the Tollfree number.
	BusinessCountry *string `json:"business_country,omitempty"`
		// The website of the business or organization using the Tollfree number.
	BusinessWebsite *string `json:"business_website,omitempty"`
		// The first name of the contact for the business or organization using the Tollfree number.
	BusinessContactFirstName *string `json:"business_contact_first_name,omitempty"`
		// The last name of the contact for the business or organization using the Tollfree number.
	BusinessContactLastName *string `json:"business_contact_last_name,omitempty"`
		// The email address of the contact for the business or organization using the Tollfree number.
	BusinessContactEmail *string `json:"business_contact_email,omitempty"`
		// The phone number of the contact for the business or organization using the Tollfree number.
	BusinessContactPhone *string `json:"business_contact_phone,omitempty"`
		// The email address to receive the notification about the verification result. .
	NotificationEmail *string `json:"notification_email,omitempty"`
		// The category of the use case for the Tollfree Number. List as many are applicable..
	UseCaseCategories *[]string `json:"use_case_categories,omitempty"`
		// Use this to further explain how messaging is used by the business or organization.
	UseCaseSummary *string `json:"use_case_summary,omitempty"`
		// An example of message content, i.e. a sample message.
	ProductionMessageSample *string `json:"production_message_sample,omitempty"`
		// Link to an image that shows the opt-in workflow. Multiple images allowed and must be a publicly hosted URL.
	OptInImageUrls *[]string `json:"opt_in_image_urls,omitempty"`
	OptInType *string `json:"opt_in_type,omitempty"`
		// Estimate monthly volume of messages from the Tollfree Number.
	MessageVolume *string `json:"message_volume,omitempty"`
		// Additional information to be provided for verification.
	AdditionalInformation *string `json:"additional_information,omitempty"`
		// The SID of the Phone Number associated with the Tollfree Verification.
	TollfreePhoneNumberSid *string `json:"tollfree_phone_number_sid,omitempty"`
	Status *string `json:"status,omitempty"`
		// The absolute URL of the Tollfree Verification resource.
	Url *string `json:"url,omitempty"`
		// The rejection reason given when a Tollfree Verification has been rejected.
	RejectionReason *string `json:"rejection_reason,omitempty"`
		// The error code given when a Tollfree Verification has been rejected.
	ErrorCode *int `json:"error_code,omitempty"`
		// The URLs of the documents associated with the Tollfree Verification resource.
	ResourceLinks *interface{} `json:"resource_links,omitempty"`
		// An optional external reference ID supplied by customer and echoed back on status retrieval.
	ExternalReferenceId *string `json:"external_reference_id,omitempty"`
}


