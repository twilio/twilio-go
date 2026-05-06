# InsightsV1Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**CallSid** | Pointer to **string** | The unique SID identifier of the Call. |
**CallType** | Pointer to [**string**](SummaryEnumCallType.md) |  |
**CallState** | Pointer to [**string**](SummaryEnumCallState.md) |  |
**AnsweredBy** | Pointer to [**string**](SummaryEnumAnsweredBy.md) |  |
**ProcessingState** | Pointer to [**string**](SummaryEnumProcessingState.md) |  |
**CreatedTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the Call was created, given in ISO 8601 format. Can be different from `start_time` in the event of queueing due to CPS |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the Call was started, given in ISO 8601 format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the Call was ended, given in ISO 8601 format. |
**Duration** | Pointer to **int** | Duration between when the call was initiated and the call was ended |
**ConnectDuration** | Pointer to **int** | Duration between when the call was answered and when it ended |
**From** | Pointer to **interface{}** | `object` The calling party. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#tofrom-object) for the object properties. |
**To** | Pointer to **interface{}** | `object` The called party. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#tofrom-object) for the object properties. |
**CarrierEdge** | Pointer to **interface{}** | `object` Contains metrics and properties for the Twilio media gateway of a PSTN call. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |
**ClientEdge** | Pointer to **interface{}** | `object` Contains metrics and properties for the Twilio media gateway of a Client call. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |
**SdkEdge** | Pointer to **interface{}** | `object` Contains metrics and properties for the SDK sensor library for Client calls. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |
**SipEdge** | Pointer to **interface{}** | `object` Contains metrics and properties for the Twilio media gateway of a SIP Interface or Trunking call. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#edges-and-their-properties) for the object properties. |
**Tags** | Pointer to **[]string** | Tags applied to calls by Voice Insights analysis indicating a condition that could result in subjective degradation of the call quality. |
**Url** | Pointer to **string** | The URL of this resource. |
**Attributes** | Pointer to **interface{}** | `object` Attributes capturing call-flow-specific details. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#attributes-object) for the object properties. |
**Properties** | Pointer to **interface{}** | `object` Contains edge-agnostic call-level details. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#properties-object) for the object properties. |
**Trust** | Pointer to **interface{}** | `object` Contains trusted communications details including Branded Call and verified caller ID. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#trust-object) for the object properties. |
**Annotation** | Pointer to **interface{}** | `object` Programmatically labeled annotations for the Call. Developers can update the Call Summary records with Annotation during or after a Call. Annotations can be updated as long as the Call Summary record is addressable via the API. See [Details: Call Summary](https://www.twilio.com/docs/voice/voice-insights/api/call/details-call-summary#annotation-object) for the object properties. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


