# InsightsV1CallSummaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**CallSid** | Pointer to **string** | The unique SID identifier of the Call. |
**AnsweredBy** | Pointer to [**string**](CallSummariesEnumAnsweredBy.md) |  |
**CallType** | Pointer to [**string**](CallSummariesEnumCallType.md) |  |
**CallState** | Pointer to [**string**](CallSummariesEnumCallState.md) |  |
**ProcessingState** | Pointer to [**string**](CallSummariesEnumProcessingState.md) |  |
**CreatedTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the Call was created, given in ISO 8601 format. Can be different from `start_time` in the event of queueing due to CPS |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the Call was started, given in ISO 8601 format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the Call was ended, given in ISO 8601 format. |
**Duration** | Pointer to **int** | Duration between when the call was initiated and the call was ended |
**ConnectDuration** | Pointer to **int** | Duration between when the call was answered and when it ended |
**From** | Pointer to **map[string]interface{}** | The calling party. |
**To** | Pointer to **map[string]interface{}** | The called party. |
**CarrierEdge** | Pointer to **map[string]interface{}** | Contains metrics and properties for the Twilio media gateway of a PSTN call. |
**ClientEdge** | Pointer to **map[string]interface{}** | Contains metrics and properties for the Twilio media gateway of a Client call. |
**SdkEdge** | Pointer to **map[string]interface{}** | Contains metrics and properties for the SDK sensor library for Client calls. |
**SipEdge** | Pointer to **map[string]interface{}** | Contains metrics and properties for the Twilio media gateway of a SIP Interface or Trunking call. |
**Tags** | Pointer to **[]string** | Tags applied to calls by Voice Insights analysis indicating a condition that could result in subjective degradation of the call quality. |
**Url** | Pointer to **string** | The URL of this resource. |
**Attributes** | Pointer to **map[string]interface{}** | Attributes capturing call-flow-specific details. |
**Properties** | Pointer to **map[string]interface{}** | Contains edge-agnostic call-level details. |
**Trust** | Pointer to **map[string]interface{}** | Contains trusted communications details including Branded Call and verified caller ID. |
**Annotation** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


