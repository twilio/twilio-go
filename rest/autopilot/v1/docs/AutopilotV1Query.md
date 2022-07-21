# AutopilotV1Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Results** | Pointer to **interface{}** | The natural language analysis results that include the Task recognized and a list of identified Fields |
**Language** | Pointer to **string** | The ISO language-country string that specifies the language used by the Query |
**ModelBuildSid** | Pointer to **string** | The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried |
**Query** | Pointer to **string** | The end-user's natural language input |
**SampleSid** | Pointer to **string** | The SID of an optional reference to the Sample created from the query |
**AssistantSid** | Pointer to **string** | The SID of the Assistant that is the parent of the resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the Query |
**Url** | Pointer to **string** | The absolute URL of the Query resource |
**SourceChannel** | Pointer to **string** | The communication channel from where the end-user input came |
**DialogueSid** | Pointer to **string** | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue). |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


