# InlineObject46

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | [optional] 
**CallbackUrl** | **string** | The URL we should call using &#x60;callback_method&#x60; when the trigger fires. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | [optional] 
**Recurring** | **string** | The frequency of a recurring UsageTrigger.  Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT. | [optional] 
**TriggerBy** | **string** | The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is &#x60;usage&#x60;. | [optional] 
**TriggerValue** | **string** | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as &#x60;+30&#x60; to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a &#x60;+&#x60; as &#x60;%2B&#x60;. | 
**UsageCategory** | **string** | The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


