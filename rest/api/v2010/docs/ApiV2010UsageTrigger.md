# ApiV2010UsageTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that this trigger monitors |[optional] 
**ApiVersion** | **string** | The API version used to create the resource |[optional] 
**CallbackMethod** | **string** | The HTTP method we use to call callback_url |[optional] 
**CallbackUrl** | **string** | he URL we call when the trigger fires |[optional] 
**CurrentValue** | **string** | The current value of the field the trigger is watching |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateFired** | **string** | The RFC 2822 date and time in GMT that the trigger was last fired |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the trigger |[optional] 
**Recurring** | Pointer to [**string**](UsageTriggerEnumRecurring.md) |  |
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**TriggerBy** | Pointer to [**string**](UsageTriggerEnumTriggerField.md) |  |
**TriggerValue** | **string** | The value at which the trigger will fire |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 
**UsageCategory** | Pointer to [**string**](UsageTriggerEnumUsageCategory.md) |  |
**UsageRecordUri** | **string** | The URI of the UsageRecord resource this trigger watches |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


