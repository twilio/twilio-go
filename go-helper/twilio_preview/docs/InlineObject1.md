# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The optional email to send the completion notification to | [optional] 
**EndDay** | **string** | The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day. | 
**FriendlyName** | **string** | The friendly name specified when creating the job | 
**StartDay** | **string** | The start day for the custom export specified as a string in the format of yyyy-mm-dd | 
**WebhookMethod** | **string** | This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied. | [optional] 
**WebhookUrl** | **string** | The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


