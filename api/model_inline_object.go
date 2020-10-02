/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio
// InlineObject struct for InlineObject
type InlineObject struct {
	// The SID of the Application resource that will handle the call, if the call will be handled by an application.
	ApplicationSid string `json:"ApplicationSid,omitempty"`
	// Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: `true` or `false`.
	AsyncAmd string `json:"AsyncAmd,omitempty"`
	// The URL that we should call using the `async_amd_status_callback_method` to notify customer application whether the call was answered by human, machine or fax.
	AsyncAmdStatusCallback string `json:"AsyncAmdStatusCallback,omitempty"`
	// The HTTP method we should use when calling the `async_amd_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`.
	AsyncAmdStatusCallbackMethod string `json:"AsyncAmdStatusCallbackMethod,omitempty"`
	// The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta)
	Byoc string `json:"Byoc,omitempty"`
	// The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta)
	CallReason string `json:"CallReason,omitempty"`
	// The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as `name@company.com`.
	CallerId string `json:"CallerId,omitempty"`
	// The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackMethod string `json:"FallbackMethod,omitempty"`
	// The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored.
	FallbackUrl string `json:"FallbackUrl,omitempty"`
	// The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `From` must also be a phone number.
	From string `json:"From"`
	// Whether to detect if a human, answering machine, or fax has picked up the call. Can be: `Enable` or `DetectMessageEnd`. Use `Enable` if you would like us to return `AnsweredBy` as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine. If `send_digits` is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection).
	MachineDetection string `json:"MachineDetection,omitempty"`
	// The number of milliseconds of initial silence after which an `unknown` AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000.
	MachineDetectionSilenceTimeout int32 `json:"MachineDetectionSilenceTimeout,omitempty"`
	// The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200.
	MachineDetectionSpeechEndThreshold int32 `json:"MachineDetectionSpeechEndThreshold,omitempty"`
	// The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400.
	MachineDetectionSpeechThreshold int32 `json:"MachineDetectionSpeechThreshold,omitempty"`
	// The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with `AnsweredBy` of `unknown`. The default timeout is 30 seconds.
	MachineDetectionTimeout int32 `json:"MachineDetectionTimeout,omitempty"`
	// The HTTP method we should use when calling the `url` parameter's value. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	Method string `json:"Method,omitempty"`
	// Whether to record the call. Can be `true` to record the phone call, or `false` to not. The default is `false`. The `recording_url` is sent to the `status_callback` URL.
	Record bool `json:"Record,omitempty"`
	// The number of channels in the final recording. Can be: `mono` or `dual`. The default is `mono`. `mono` records both legs of the call in a single channel of the recording file. `dual` records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call.
	RecordingChannels string `json:"RecordingChannels,omitempty"`
	// The URL that we call when the recording is available to be accessed.
	RecordingStatusCallback string `json:"RecordingStatusCallback,omitempty"`
	// The recording status events that will trigger calls to the URL specified in `recording_status_callback`. Can be: `in-progress`, `completed` and `absent`. Defaults to `completed`. Separate  multiple values with a space.
	RecordingStatusCallbackEvent []string `json:"RecordingStatusCallbackEvent,omitempty"`
	// The HTTP method we should use when calling the `recording_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`.
	RecordingStatusCallbackMethod string `json:"RecordingStatusCallbackMethod,omitempty"`
	// A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (`0`-`9`), '`#`', '`*`' and '`w`', to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be `ww1234#`. Remember to URL-encode this string, since the '`#`' character has special meaning in a URL. If both `SendDigits` and `MachineDetection` parameters are provided, then `MachineDetection` will be ignored.
	SendDigits string `json:"SendDigits,omitempty"`
	// The password required to authenticate the user account specified in `sip_auth_username`.
	SipAuthPassword string `json:"SipAuthPassword,omitempty"`
	// The username used to authenticate the caller making a SIP call.
	SipAuthUsername string `json:"SipAuthUsername,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
	StatusCallback string `json:"StatusCallback,omitempty"`
	// The call progress events that we will send to the `status_callback` URL. Can be: `initiated`, `ringing`, `answered`, and `completed`. If no event is specified, we send the `completed` status. If you want to receive multiple events, specify each one in a separate `status_callback_event` parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample=code-create-a-call-resource-and-specify-a-statuscallbackevent&code-sdk-version=json). If an `application_sid` is present, this parameter is ignored.
	StatusCallbackEvent []string `json:"StatusCallbackEvent,omitempty"`
	// The HTTP method we should use when calling the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	// The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is `60` seconds and the maximum is `600` seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as `15` seconds, to hang up before reaching an answering machine or voicemail.
	Timeout int32 `json:"Timeout,omitempty"`
	// The phone number, SIP address, or client identifier to call.
	To string `json:"To"`
	// Whether to trim any leading and trailing silence from the recording. Can be: `trim-silence` or `do-not-trim` and the default is `trim-silence`.
	Trim string `json:"Trim,omitempty"`
	// TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both `twiml` and `url` are provided then `twiml` parameter will be ignored.
	Twiml string `json:"Twiml,omitempty"`
	// The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).
	Url string `json:"Url,omitempty"`
}
