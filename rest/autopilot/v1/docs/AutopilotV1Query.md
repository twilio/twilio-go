# AutopilotV1Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Query resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Results** | Pointer to **interface{}** | The natural language analysis results that include the [Task](https://www.twilio.com/docs/autopilot/api/task) recognized and a list of identified [Fields](https://www.twilio.com/docs/autopilot/api/task-field). |
**Language** | Pointer to **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used by the Query. For example: `en-US`. |
**ModelBuildSid** | Pointer to **string** | The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried. |
**Query** | Pointer to **string** | The end-user's natural language input. |
**SampleSid** | Pointer to **string** | The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query. |
**AssistantSid** | Pointer to **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Query resource. |
**Status** | Pointer to **string** | The status of the Query. Can be: `pending-review`, `reviewed`, or `discarded` |
**Url** | Pointer to **string** | The absolute URL of the Query resource. |
**SourceChannel** | Pointer to **string** | The communication channel from where the end-user input came. |
**DialogueSid** | Pointer to **string** | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue). |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


