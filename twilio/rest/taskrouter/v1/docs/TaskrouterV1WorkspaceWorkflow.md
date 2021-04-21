# TaskrouterV1WorkspaceWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AssignmentCallbackUrl** | Pointer to **string** | The URL that we call when a task managed by the Workflow is assigned to a Worker |
**Configuration** | Pointer to **string** | A JSON string that contains the Workflow's configuration |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**DocumentContentType** | Pointer to **string** | The MIME type of the document |
**FallbackAssignmentCallbackUrl** | Pointer to **string** | The URL that we call when a call to the `assignment_callback_url` fails |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Workflow resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TaskReservationTimeout** | Pointer to **int32** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker |
**Url** | Pointer to **string** | The absolute URL of the Workflow resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workflow |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


