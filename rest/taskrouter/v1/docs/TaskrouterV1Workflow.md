# TaskrouterV1Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Workflow resource. |
**AssignmentCallbackUrl** | Pointer to **string** | The URL that we call when a task managed by the Workflow is assigned to a Worker. See Assignment Callback URL for more information. |
**Configuration** | Pointer to **string** | A JSON string that contains the Workflow's configuration. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DocumentContentType** | Pointer to **string** | The MIME type of the document. |
**FallbackAssignmentCallbackUrl** | Pointer to **string** | The URL that we call when a call to the `assignment_callback_url` fails. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Workflow resource. For example, `Customer Support` or `2014 Election Campaign`. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Workflow resource. |
**TaskReservationTimeout** | Pointer to **int** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workflow. |
**Url** | Pointer to **string** | The absolute URL of the Workflow resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


