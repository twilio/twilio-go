# InlineObject9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \&quot;{}\&quot; will be returned. | [optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. | [optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. | [optional] 
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional. | [optional] 
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to. | [optional] 
**State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60; | [optional] 
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes. | [optional] 
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


