# ApiV2010UsageTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that this trigger monitors |
**ApiVersion** | Pointer to **string** | The API version used to create the resource |
**CallbackMethod** | Pointer to **string** | The HTTP method we use to call callback_url |
**CallbackUrl** | Pointer to **string** | he URL we call when the trigger fires |
**CurrentValue** | Pointer to **string** | The current value of the field the trigger is watching |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateFired** | Pointer to **string** | The RFC 2822 date and time in GMT that the trigger was last fired |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the trigger |
**Recurring** | Pointer to **string** | The frequency of a recurring UsageTrigger |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TriggerBy** | Pointer to **string** | The field in the UsageRecord resource that fires the trigger |
**TriggerValue** | Pointer to **string** | The value at which the trigger will fire |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**UsageCategory** | Pointer to **string** | The usage category the trigger watches |
**UsageRecordUri** | Pointer to **string** | The URI of the UsageRecord resource this trigger watches |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


