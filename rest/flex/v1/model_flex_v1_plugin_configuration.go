/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// FlexV1PluginConfiguration struct for FlexV1PluginConfiguration
type FlexV1PluginConfiguration struct {
	// The unique string that we created to identify the Flex Plugin Configuration resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Plugin Configuration resource and owns this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The name of this Flex Plugin Configuration.
	Name *string `json:"name,omitempty"`
	// The description of the Flex Plugin Configuration resource.
	Description *string `json:"description,omitempty"`
	// Whether the Flex Plugin Configuration is archived. The default value is false.
	Archived *bool `json:"archived,omitempty"`
	// The date and time in GMT when the Flex Plugin Configuration was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The absolute URL of the Flex Plugin Configuration resource.
	Url   *string                 `json:"url,omitempty"`
	Links *map[string]interface{} `json:"links,omitempty"`
}
