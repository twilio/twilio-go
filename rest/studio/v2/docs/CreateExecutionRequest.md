# CreateExecutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow&#39;s Execution. Available as variable &#x60;{{flow.channel.address}}&#x60;. | 
**Parameters** | [**map[string]interface{}**](.md) | JSON data that will be added to the Flow&#39;s context and that can be accessed as variables inside your Flow. For example, if you pass in &#x60;Parameters&#x3D;{\&quot;name\&quot;:\&quot;Zeke\&quot;}&#x60;, a widget in your Flow can reference the variable &#x60;{{flow.data.name}}&#x60;, which returns \&quot;Zeke\&quot;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string. | [optional] 
**To** | **string** | The Contact phone number to start a Studio Flow Execution, available as variable &#x60;{{contact.channel.address}}&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


