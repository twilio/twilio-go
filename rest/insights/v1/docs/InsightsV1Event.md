# InsightsV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | Event time. |
**CallSid** | Pointer to **string** | The unique SID identifier of the Call. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**Edge** | Pointer to [**string**](EventEnumTwilioEdge.md) |  |
**Group** | Pointer to **string** | Event group. |
**Level** | Pointer to [**string**](EventEnumLevel.md) |  |
**Name** | Pointer to **string** | Event name. |
**CarrierEdge** | Pointer to **interface{}** | `object` Represents the connection between Twilio and our immediate carrier partners. The events here describe the call lifecycle as reported by Twilio's carrier media gateways. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |
**SipEdge** | Pointer to **interface{}** | `object` Represents the Twilio media gateway for SIP interface and SIP trunking calls. The events here describe the call lifecycle as reported by Twilio's public media gateways. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |
**SdkEdge** | Pointer to **interface{}** | `object` Represents the Voice SDK running locally in the browser or in the Android/iOS application. The events here are emitted by the Voice SDK in response to certain call progress events, network changes, or call quality conditions. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |
**ClientEdge** | Pointer to **interface{}** | `object` Represents the Twilio media gateway for Client calls. The events here describe the call lifecycle as reported by Twilio's Voice SDK media gateways. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


