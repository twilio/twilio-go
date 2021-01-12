/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackMethod string `json:"FallbackMethod,omitempty"`
	// The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackUrl string `json:"FallbackUrl,omitempty"`
	// The HTTP method we should use when calling the `url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	Method string `json:"Method,omitempty"`
	// The new status of the resource. Can be: `canceled` or `completed`. Specifying `canceled` will attempt to hang up calls that are queued or ringing; however, it will not affect calls already in progress. Specifying `completed` will attempt to hang up a call even if it's already in progress.
	Status string `json:"Status,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
	StatusCallback string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use when requesting the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	// TwiML instructions for the call Twilio will use without fetching Twiml from url. Twiml and url parameters are mutually exclusive
	Twiml string `json:"Twiml,omitempty"`
	// The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).
	Url string `json:"Url,omitempty"`
}
