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
)

// MessagingV1BrandRegistrationOtp struct for MessagingV1BrandRegistrationOtp
type MessagingV1BrandRegistrationOtp struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Brand Registration resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique string to identify Brand Registration of Sole Proprietor Brand
	BrandRegistrationSid *string `json:"brand_registration_sid,omitempty"`
}
