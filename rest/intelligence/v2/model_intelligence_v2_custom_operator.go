/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
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

// IntelligenceV2CustomOperator struct for IntelligenceV2CustomOperator
type IntelligenceV2CustomOperator struct {
	// The unique SID identifier of the Account the Custom Operator belongs to.
	AccountSid *string `json:"account_sid,omitempty"`
	// A 34 character string that uniquely identifies this Custom Operator.
	Sid *string `json:"sid,omitempty"`
	// A human-readable name of this resource, up to 64 characters.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A human-readable description of this resource, longer than the friendly name.
	Description *string `json:"description,omitempty"`
	// The creator of the Custom Operator. Custom Operators can only be created by a Twilio Account.
	Author *string `json:"author,omitempty"`
	// Operator Type for this Operator. References an existing Operator Type resource.
	OperatorType *string `json:"operator_type,omitempty"`
	// Numeric Custom Operator version. Incremented with each update on the resource, used to ensure integrity when updating the Custom Operator.
	Version      int     `json:"version,omitempty"`
	Availability *string `json:"availability,omitempty"`
	// Operator configuration, following the schema defined by the Operator Type. Only available on Operators created by the Account.
	Config *map[string]interface{} `json:"config,omitempty"`
	// The date that this Custom Operator was created, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Custom Operator was updated, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
