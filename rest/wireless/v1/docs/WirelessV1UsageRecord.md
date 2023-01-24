# WirelessV1UsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimSid** | Pointer to **string** | The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource) that this Usage Record is for. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resource. |
**Period** | Pointer to **interface{}** | The time period for which the usage is reported. Contains `start` and `end` datetime values given as GMT in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**Commands** | Pointer to **interface{}** | An object that describes the SIM's usage of Commands during the specified period. See [Commands Usage Object](https://www.twilio.com/docs/wireless/api/sim-usagerecord-resource#commands-usage-object). |
**Data** | Pointer to **interface{}** | An object that describes the SIM's data usage during the specified period. See [Data Usage Object](https://www.twilio.com/docs/wireless/api/sim-usagerecord-resource#data-usage-object). |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


