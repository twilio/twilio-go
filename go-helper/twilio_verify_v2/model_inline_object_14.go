/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject14 struct for InlineObject14
type InlineObject14 struct {
	// The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
	Amount string `json:"Amount,omitempty"`
	// Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: `<#> Your AppName verification code is: 1234 He42w354ol9`.
	AppHash string `json:"AppHash,omitempty"`
	// The verification method to use. Can be: [`email`](https://www.twilio.com/docs/verify/email), `sms` or `call`.
	Channel string `json:"Channel"`
	// [`email`](https://www.twilio.com/docs/verify/email) channel configuration in json format. Must include 'from' and 'from_name'.
	ChannelConfiguration map[string]interface{} `json:"ChannelConfiguration,omitempty"`
	// A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive.
	CustomCode string `json:"CustomCode,omitempty"`
	// A custom user defined friendly name that overwrites the existing one in the verification message
	CustomFriendlyName string `json:"CustomFriendlyName,omitempty"`
	// The text of a custom message to use for the verification.
	CustomMessage string `json:"CustomMessage,omitempty"`
	// The locale to use for the verification SMS or call. Can be: `af`, `ar`, `ca`, `cs`, `da`, `de`, `el`, `en`, `en-GB`, `es`, `fi`, `fr`, `he`, `hi`, `hr`, `hu`, `id`, `it`, `ja`, `ko`, `ms`, `nb`, `nl`, `pl`, `pt`, `pr-BR`, `ro`, `ru`, `sv`, `th`, `tl`, `tr`, `vi`, `zh`, `zh-CN`, or `zh-HK.`
	Locale string `json:"Locale,omitempty"`
	// The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
	Payee string `json:"Payee,omitempty"`
	// The custom key-value pairs of Programmable Rate Limits. Keys correspond to `unique_name` fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request.
	RateLimits map[string]interface{} `json:"RateLimits,omitempty"`
	// The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits).
	SendDigits string `json:"SendDigits,omitempty"`
	// The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
	To string `json:"To"`
}
