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
**CarrierEdge** | Pointer to **interface{}** | Represents the connection between Twilio and our immediate carrier partners. The events here describe the call lifecycle as reported by Twilio's carrier media gateways. |
**SipEdge** | Pointer to **interface{}** | Represents the Twilio media gateway for SIP interface and SIP trunking calls. The events here describe the call lifecycle as reported by Twilio's public media gateways. |
**SdkEdge** | Pointer to **interface{}** | Represents the Voice SDK running locally in the browser or in the Android/iOS application. The events here are emitted by the Voice SDK in response to certain call progress events, network changes, or call quality conditions. |
**ClientEdge** | Pointer to **interface{}** | Represents the Twilio media gateway for Client calls. The events here describe the call lifecycle as reported by Twilio's Voice SDK media gateways. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


