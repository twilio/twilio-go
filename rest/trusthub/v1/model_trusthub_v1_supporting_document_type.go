/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// TrusthubV1SupportingDocumentType struct for TrusthubV1SupportingDocumentType
type TrusthubV1SupportingDocumentType struct {
	// The unique string that identifies the Supporting Document Type resource.
	Sid *string `json:"sid,omitempty"`
	// A human-readable description of the Supporting Document Type resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The machine-readable description of the Supporting Document Type resource.
	MachineName *string `json:"machine_name,omitempty"`
	// The required information for creating a Supporting Document. The required fields will change as regulatory needs change and will differ for businesses and individuals.
	Fields *[]map[string]interface{} `json:"fields,omitempty"`
	// The absolute URL of the Supporting Document Type resource.
	Url *string `json:"url,omitempty"`
}
