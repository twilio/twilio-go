# InlineObject50

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackEvents** | **string** | Space-separated list of callback events that will trigger callbacks. | [optional] 
**CallbackUrl** | **string** | A user-provided URL to send event callbacks to. | [optional] 
**FallbackActions** | [**map[string]interface{}**](.md) | The JSON actions to be executed when the user&#39;s input is not recognized as matching any Task. | [optional] 
**FriendlyName** | **string** | A text description for the Assistant. It is non-unique and can up to 255 characters long. | [optional] 
**InitiationActions** | [**map[string]interface{}**](.md) | The JSON actions to be executed on inbound phone calls when the Assistant has to say something first. | [optional] 
**LogQueries** | **bool** | A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided. | [optional] 
**StyleSheet** | [**map[string]interface{}**](.md) | The JSON object that holds the style sheet for the assistant | [optional] 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


