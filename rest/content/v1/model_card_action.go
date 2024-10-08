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

// CardAction struct for CardAction
type CardAction struct {
	Type  CardActionType `json:"type"`
	Title string         `json:"title"`
	Url   string         `json:"url,omitempty"`
	Phone string         `json:"phone,omitempty"`
	Id    string         `json:"id,omitempty"`
	Code  string         `json:"code,omitempty"`
}
