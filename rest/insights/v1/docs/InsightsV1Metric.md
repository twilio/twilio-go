# InsightsV1Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | Timestamp of metric sample. Samples are taken every 10 seconds and contain the metrics for the previous 10 seconds. |
**CallSid** | Pointer to **string** | The unique SID identifier of the Call. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**Edge** | Pointer to [**string**](MetricEnumTwilioEdge.md) |  |
**Direction** | Pointer to [**string**](MetricEnumStreamDirection.md) |  |
**CarrierEdge** | Pointer to **interface{}** | Contains metrics and properties for the Twilio media gateway of a PSTN call. |
**SipEdge** | Pointer to **interface{}** | Contains metrics and properties for the Twilio media gateway of a SIP Interface or Trunking call. |
**SdkEdge** | Pointer to **interface{}** | Contains metrics and properties for the SDK sensor library for Client calls. |
**ClientEdge** | Pointer to **interface{}** | Contains metrics and properties for the Twilio media gateway of a Client call. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


