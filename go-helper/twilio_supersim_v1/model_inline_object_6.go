/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	// The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource's status is new.
	AccountSid string `json:"AccountSid,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_method` after an asynchronous update has finished.
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	// The SID or unique name of the Fleet to which the SIM resource should be assigned.
	Fleet string `json:"Fleet,omitempty"`
	// The new status of the resource. Can be: `ready`, `active`, or `inactive`. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info.
	Status string `json:"Status,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
}
