# CreateEngagementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable &#x60;{{flow.channel.address}}&#x60; | 
**Parameters** | [**map[string]interface{}**](.md) | A JSON string we will add to your flow&#39;s context and that you can access as variables inside your flow. For example, if you pass in &#x60;Parameters&#x3D;{&#39;name&#39;:&#39;Zeke&#39;}&#x60; then inside a widget you can reference the variable &#x60;{{flow.data.name}}&#x60; which will return the string &#39;Zeke&#39;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string. | [optional] 
**To** | **string** | The Contact phone number to start a Studio Flow Engagement, available as variable &#x60;{{contact.channel.address}}&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


