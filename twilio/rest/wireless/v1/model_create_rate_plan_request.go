/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateRatePlanRequest struct for CreateRatePlanRequest
type CreateRatePlanRequest struct {
	// Whether SIMs can use GPRS/3G/4G/LTE data connectivity.
	DataEnabled bool `json:"DataEnabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is `1000`.
	DataLimit int32 `json:"DataLimit,omitempty"`
	// The model used to meter data usage. Can be: `payg` and `quota-1`, `quota-10`, and `quota-50`. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans).
	DataMetering string `json:"DataMetering,omitempty"`
	// A descriptive string that you create to describe the resource. It does not have to be unique.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can be: `data`, `voice`, and `messaging`.
	InternationalRoaming []string `json:"InternationalRoaming,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB.
	InternationalRoamingDataLimit int32 `json:"InternationalRoamingDataLimit,omitempty"`
	// Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource).
	MessagingEnabled bool `json:"MessagingEnabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info.
	NationalRoamingDataLimit int32 `json:"NationalRoamingDataLimit,omitempty"`
	// Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming).
	NationalRoamingEnabled bool `json:"NationalRoamingEnabled,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
	// Whether SIMs can make and receive voice calls.
	VoiceEnabled bool `json:"VoiceEnabled,omitempty"`
}
