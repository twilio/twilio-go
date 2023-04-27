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
	"time"
)
// VerifyV2NewFactor struct for VerifyV2NewFactor
type VerifyV2NewFactor struct {
		// A 34 character string that uniquely identifies this Factor.
	Sid *string `json:"sid,omitempty"`
		// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
		// The unique SID identifier of the Service.
	ServiceSid *string `json:"service_sid,omitempty"`
		// The unique SID identifier of the Entity.
	EntitySid *string `json:"entity_sid,omitempty"`
		// Customer unique identity for the Entity owner of the Factor. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters.
	Identity *string `json:"identity,omitempty"`
		// Contains the `factor_type` specific secret and metadata. For push, this is `binding.public_key` and `binding.alg`. For totp, this is `binding.secret` and `binding.uri`. The `binding.uri` property is generated following the [google authenticator key URI format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format), and `Factor.friendly_name` is used for the “accountname” value and `Service.friendly_name` or `Service.totp.issuer` is used for the `issuer` value.   The Binding property is ONLY returned upon Factor creation.
	Binding *interface{} `json:"binding,omitempty"`
		// The date that this Factor was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date that this Factor was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The friendly name of this Factor. This can be any string up to 64 characters, meant for humans to distinguish between Factors. For `factor_type` `push`, this could be a device name. For `factor_type` `totp`, this value is used as the “account name” in constructing the `binding.uri` property. At the same time, we recommend avoiding providing PII.
	FriendlyName *string `json:"friendly_name,omitempty"`
	Status *string `json:"status,omitempty"`
	FactorType *string `json:"factor_type,omitempty"`
		// An object that contains configurations specific to a `factor_type`.
	Config *interface{} `json:"config,omitempty"`
		// Custom metadata associated with the factor. This is added by the Device/SDK directly to allow for the inclusion of device information. It must be a stringified JSON with only strings values eg. `{\"os\": \"Android\"}`. Can be up to 1024 characters in length.
	Metadata *interface{} `json:"metadata,omitempty"`
		// The URL of this resource.
	Url *string `json:"url,omitempty"`
}


