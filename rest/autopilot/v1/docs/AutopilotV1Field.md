# AutopilotV1Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AssistantSid** | Pointer to **string** | The SID of the Assistant that is the parent of the Task associated with the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**FieldType** | Pointer to **string** | The Field Type of the field |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TaskSid** | Pointer to **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with this Field |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Field resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


