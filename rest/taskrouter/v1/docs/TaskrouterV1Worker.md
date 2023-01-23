# TaskrouterV1Worker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Worker resource. |
**ActivityName** | Pointer to **string** | The `friendly_name` of the Worker's current Activity. |
**ActivitySid** | Pointer to **string** | The SID of the Worker's current Activity. |
**Attributes** | Pointer to **string** | The JSON string that describes the Worker. For example: `{ \"email\": \"Bob@example.com\", \"phone\": \"+5095551234\" }`. **Note** If this property has been assigned a value, it will only be displayed in FETCH actions that return a single resource. Otherwise, this property will be null, even if it has a value. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. |
**Available** | Pointer to **bool** | Whether the Worker is available to perform tasks. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateStatusChanged** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT of the last change to the Worker's activity specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Used to calculate Workflow statistics. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. Friendly names are case insensitive, and unique within the TaskRouter Workspace. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Worker resource. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Worker. |
**Url** | Pointer to **string** | The absolute URL of the Worker resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


