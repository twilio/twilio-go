# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackEvents** | **string** | Reserved. | [optional] 
**CallbackUrl** | **string** | Reserved. | [optional] 
**Defaults** | [**map[string]interface{}**](.md) | A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. | [optional] 
**LogQueries** | **bool** | Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored. | [optional] 
**StyleSheet** | [**map[string]interface{}**](.md) | The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet) | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


