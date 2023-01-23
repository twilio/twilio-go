# ApiV2010UsageTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the trigger monitors. |
**ApiVersion** | Pointer to **string** | The API version used to create the resource. |
**CallbackMethod** | Pointer to **string** | The HTTP method we use to call `callback_url`. Can be: `GET` or `POST`. |
**CallbackUrl** | Pointer to **string** | The URL we call using the `callback_method` when the trigger fires. |
**CurrentValue** | Pointer to **string** | The current value of the field the trigger is watching. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateFired** | Pointer to **string** | The date and time in GMT that the trigger was last fired specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the trigger. |
**Recurring** | Pointer to [**string**](UsageTriggerEnumRecurring.md) |  |
**Sid** | Pointer to **string** | The unique string that that we created to identify the UsageTrigger resource. |
**TriggerBy** | Pointer to [**string**](UsageTriggerEnumTriggerField.md) |  |
**TriggerValue** | Pointer to **string** | The value at which the trigger will fire.  Must be a positive, numeric value. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**UsageCategory** | Pointer to [**string**](UsageTriggerEnumUsageCategory.md) |  |
**UsageRecordUri** | Pointer to **string** | The URI of the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource this trigger watches, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


