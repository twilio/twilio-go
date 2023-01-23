# AutopilotV1Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**Results** | Pointer to **interface{}** | The natural language analysis results that include the Task recognized and a list of identified Fields |
**Language** | **string** | The ISO language-country string that specifies the language used by the Query |[optional] 
**ModelBuildSid** | **string** | The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried |[optional] 
**Query** | **string** | The end-user's natural language input |[optional] 
**SampleSid** | **string** | The SID of an optional reference to the Sample created from the query |[optional] 
**AssistantSid** | **string** | The SID of the Assistant that is the parent of the resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Status** | **string** | The status of the Query |[optional] 
**Url** | **string** | The absolute URL of the Query resource |[optional] 
**SourceChannel** | **string** | The communication channel from where the end-user input came |[optional] 
**DialogueSid** | **string** | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue). |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


