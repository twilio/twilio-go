/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// SupersimV1Network struct for SupersimV1Network
type SupersimV1Network struct {
	// The unique string that we created to identify the Network resource.
	Sid *string `json:"sid,omitempty"`
	// A human readable identifier of this resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The absolute URL of the Network resource.
	Url *string `json:"url,omitempty"`
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resource.
	IsoCountry *string `json:"iso_country,omitempty"`
	// Array of objects identifying the [MCC-MNCs](https://en.wikipedia.org/wiki/Mobile_country_code) that are included in the Network resource.
	Identifiers *[]map[string]interface{} `json:"identifiers,omitempty"`
}
