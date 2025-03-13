/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
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

// IpMessagingV2Binding struct for IpMessagingV2Binding
type IpMessagingV2Binding struct {
	Sid           *string                 `json:"sid,omitempty"`
	AccountSid    *string                 `json:"account_sid,omitempty"`
	ServiceSid    *string                 `json:"service_sid,omitempty"`
	DateCreated   *time.Time              `json:"date_created,omitempty"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	Endpoint      *string                 `json:"endpoint,omitempty"`
	Identity      *string                 `json:"identity,omitempty"`
	CredentialSid *string                 `json:"credential_sid,omitempty"`
	BindingType   *string                 `json:"binding_type,omitempty"`
	MessageTypes  *[]string               `json:"message_types,omitempty"`
	Url           *string                 `json:"url,omitempty"`
	Links         *map[string]interface{} `json:"links,omitempty"`
}
