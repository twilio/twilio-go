# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \&quot;{}\&quot; will be returned. | [optional] 
**Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;. | [optional] 
**Body** | **string** | The content of the message, can be up to 1,600 characters long. | [optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. | [optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited. | [optional] 
**MediaSid** | **string** | The Media SID to be attached to the new Message. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


