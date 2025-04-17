/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Content
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// TwilioCard twilio/card is a structured template which can be used to send a series of related information. It must include a title and at least one additional field.
type TwilioCard struct {
	Title    string       `json:"title"`
	Subtitle string       `json:"subtitle,omitempty"`
	Media    []string     `json:"media,omitempty"`
	Actions  []CardAction `json:"actions,omitempty"`
}
