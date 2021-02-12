/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateUsageTriggerRequest struct for CreateUsageTriggerRequest
type CreateUsageTriggerRequest struct {
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`.
	CallbackMethod string `json:"CallbackMethod,omitempty"`
	// The URL we should call using `callback_method` when the trigger fires.
	CallbackUrl string `json:"CallbackUrl"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The frequency of a recurring UsageTrigger.  Can be: `daily`, `monthly`, or `yearly` for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT.
	Recurring string `json:"Recurring,omitempty"`
	// The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is `usage`.
	TriggerBy string `json:"TriggerBy,omitempty"`
	// The usage value at which the trigger should fire.  For convenience, you can use an offset value such as `+30` to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a `+` as `%2B`.
	TriggerValue string `json:"TriggerValue"`
	// The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value.
	UsageCategory string `json:"UsageCategory"`
}
