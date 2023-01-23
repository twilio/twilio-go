# AutopilotV1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Task resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. It is not unique and can be up to 255 characters long. |
**Links** | Pointer to **map[string]interface{}** | A list of the URLs of related resources. |
**AssistantSid** | Pointer to **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Task resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**ActionsUrl** | Pointer to **string** | The URL from which the Assistant can fetch actions. |
**Url** | Pointer to **string** | The absolute URL of the Task resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


