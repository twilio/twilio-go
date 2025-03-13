/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// InsightsV1Event struct for InsightsV1Event
type InsightsV1Event struct {
	// Event time.
	Timestamp *string `json:"timestamp,omitempty"`
	// The unique SID identifier of the Call.
	CallSid *string `json:"call_sid,omitempty"`
	// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
	Edge       *string `json:"edge,omitempty"`
	// Event group.
	Group *string `json:"group,omitempty"`
	Level *string `json:"level,omitempty"`
	// Event name.
	Name *string `json:"name,omitempty"`
	// Represents the connection between Twilio and our immediate carrier partners. The events here describe the call lifecycle as reported by Twilio's carrier media gateways.
	CarrierEdge *map[string]interface{} `json:"carrier_edge,omitempty"`
	// Represents the Twilio media gateway for SIP interface and SIP trunking calls. The events here describe the call lifecycle as reported by Twilio's public media gateways.
	SipEdge *map[string]interface{} `json:"sip_edge,omitempty"`
	// Represents the Voice SDK running locally in the browser or in the Android/iOS application. The events here are emitted by the Voice SDK in response to certain call progress events, network changes, or call quality conditions.
	SdkEdge *map[string]interface{} `json:"sdk_edge,omitempty"`
	// Represents the Twilio media gateway for Client calls. The events here describe the call lifecycle as reported by Twilio's Voice SDK media gateways.
	ClientEdge *map[string]interface{} `json:"client_edge,omitempty"`
}
