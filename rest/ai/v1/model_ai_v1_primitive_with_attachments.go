/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ai
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

// AiV1PrimitiveWithAttachments struct for AiV1PrimitiveWithAttachments
type AiV1PrimitiveWithAttachments struct {
	// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique SID identifier of the Primitive.
	Sid *string `json:"sid,omitempty"`
	// Provides a unique and addressable name to be assigned to this Primitive, assigned by the developer, to be optionally used in addition to SID.
	UniqueName *string `json:"unique_name,omitempty"`
	// Defines the underlying model for the primitive.
	ModelUrn *string `json:"model_urn,omitempty"`
	// A human readable description of this resource, up to 64 characters.
	FriendlyName    *string `json:"friendly_name,omitempty"`
	Type            *string `json:"type,omitempty"`
	PrimitiveStatus *string `json:"primitive_status,omitempty"`
	// The date that this Primitive was created, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Primitive was updated, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The date that this Primitive was last built, given in ISO 8601 format.
	LastBuilt *time.Time `json:"last_built,omitempty"`
	// Defines the configuration and examples of either the Literal Phrase Set or Semantic Phrase Set.
	Definition *interface{} `json:"definition,omitempty"`
	// List of Services attached to the Primitive, either active or inactive.
	Attachments *interface{} `json:"attachments,omitempty"`
	// Configuration of the primitive.
	Config *interface{} `json:"config,omitempty"`
}
