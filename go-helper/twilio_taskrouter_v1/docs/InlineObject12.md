# InlineObject12

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivitySid** | **string** | The SID of a valid Activity that will describe the Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. | [optional] 
**Attributes** | **string** | The JSON string that describes the Worker. For example: &#x60;{ \&quot;email\&quot;: \&quot;Bob@example.com\&quot;, \&quot;phone\&quot;: \&quot;+5095551234\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the Worker. It can be up to 64 characters long. | [optional] 
**RejectPendingReservations** | **bool** | Whether to reject pending reservations. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


