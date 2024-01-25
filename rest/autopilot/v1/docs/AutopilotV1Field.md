# AutopilotV1Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Field resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FieldType** | Pointer to **string** | The Field Type of the field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the unique_name, or the SID of a custom Field Type. |
**TaskSid** | Pointer to **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with this Field. |
**AssistantSid** | Pointer to **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Field resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**Url** | Pointer to **string** | The absolute URL of the Field resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


