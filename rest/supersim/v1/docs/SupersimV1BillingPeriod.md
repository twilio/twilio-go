# SupersimV1BillingPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The SID of the Billing Period. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) the Super SIM belongs to. |
**SimSid** | Pointer to **string** | The SID of the Super SIM the Billing Period belongs to. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start time of the Billing Period specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end time of the Billing Period specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**PeriodType** | Pointer to [**string**](BillingPeriodEnumBpType.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


