# WirelessV1AccountUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AccountUsageRecord resource. |
**Period** | Pointer to **interface{}** | The time period for which usage is reported. Contains `start` and `end` properties that describe the period using GMT date-time values specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**Commands** | Pointer to **interface{}** | An object that describes the aggregated Commands usage for all SIMs during the specified period. See [Commands Usage Object](https://www.twilio.com/docs/wireless/api/account-usagerecord-resource#commands-usage-object). |
**Data** | Pointer to **interface{}** | An object that describes the aggregated Data usage for all SIMs over the period. See [Data Usage Object](https://www.twilio.com/docs/wireless/api/account-usagerecord-resource#data-usage-object). |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


