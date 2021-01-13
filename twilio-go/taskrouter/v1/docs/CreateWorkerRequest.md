# CreateWorkerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivitySid** | **string** | The SID of a valid Activity that will describe the new Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker&#39;s initial state is the &#x60;default_activity_sid&#x60; configured on the Workspace. | [optional] 
**Attributes** | **string** | A valid JSON string that describes the new Worker. For example: &#x60;{ \&quot;email\&quot;: \&quot;Bob@example.com\&quot;, \&quot;phone\&quot;: \&quot;+5095551234\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the new Worker. It can be up to 64 characters long. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


