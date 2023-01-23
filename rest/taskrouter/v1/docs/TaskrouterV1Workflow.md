# TaskrouterV1Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AssignmentCallbackUrl** | **string** | The URL that we call when a task managed by the Workflow is assigned to a Worker |[optional] 
**Configuration** | **string** | A JSON string that contains the Workflow's configuration |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**DocumentContentType** | **string** | The MIME type of the document |[optional] 
**FallbackAssignmentCallbackUrl** | **string** | The URL that we call when a call to the `assignment_callback_url` fails |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the Workflow resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**TaskReservationTimeout** | **int** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the Workflow |[optional] 
**Url** | **string** | The absolute URL of the Workflow resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


