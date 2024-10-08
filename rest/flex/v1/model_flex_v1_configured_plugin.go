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
	"time"
)

// FlexV1ConfiguredPlugin struct for FlexV1ConfiguredPlugin
type FlexV1ConfiguredPlugin struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the Flex Plugin resource is installed for.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Flex Plugin Configuration that this Flex Plugin belongs to.
	ConfigurationSid *string `json:"configuration_sid,omitempty"`
	// The SID of the Flex Plugin.
	PluginSid *string `json:"plugin_sid,omitempty"`
	// The SID of the Flex Plugin Version.
	PluginVersionSid *string `json:"plugin_version_sid,omitempty"`
	// The phase this Flex Plugin would initialize at runtime.
	Phase int `json:"phase,omitempty"`
	// The URL of where the Flex Plugin Version JavaScript bundle is hosted on.
	PluginUrl *string `json:"plugin_url,omitempty"`
	// The name that uniquely identifies this Flex Plugin resource.
	UniqueName *string `json:"unique_name,omitempty"`
	// The friendly name of this Flex Plugin resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A descriptive string that you create to describe the plugin resource. It can be up to 500 characters long
	Description *string `json:"description,omitempty"`
	// Whether the Flex Plugin is archived. The default value is false.
	PluginArchived *bool `json:"plugin_archived,omitempty"`
	// The latest version of this Flex Plugin Version.
	Version *string `json:"version,omitempty"`
	// A changelog that describes the changes this Flex Plugin Version brings.
	Changelog *string `json:"changelog,omitempty"`
	// Whether the Flex Plugin Version is archived. The default value is false.
	PluginVersionArchived *bool `json:"plugin_version_archived,omitempty"`
	// Whether to validate the request is authorized to access the Flex Plugin Version.
	Private *bool `json:"private,omitempty"`
	// The date and time in GMT when the Flex Plugin was installed specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The absolute URL of the Flex Plugin resource.
	Url *string `json:"url,omitempty"`
}
