# TaskrouterV1WorkerChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Worker resource. |
**AssignedTasks** | Pointer to **int** | The total number of Tasks assigned to Worker for the TaskChannel type. |
**Available** | Pointer to **bool** | Whether the Worker should receive Tasks of the TaskChannel type. |
**AvailableCapacityPercentage** | Pointer to **int** | The current percentage of capacity the TaskChannel has available. Can be a number between `0` and `100`. A value of `0` indicates that TaskChannel has no capacity available and a value of `100` means the  Worker is available to receive any Tasks of this TaskChannel type. |
**ConfiguredCapacity** | Pointer to **int** | The current configured capacity for the WorkerChannel. TaskRouter will not create any reservations after the assigned Tasks for the Worker reaches the value. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Sid** | Pointer to **string** | The unique string that we created to identify the WorkerChannel resource. |
**TaskChannelSid** | Pointer to **string** | The SID of the TaskChannel. |
**TaskChannelUniqueName** | Pointer to **string** | The unique name of the TaskChannel, such as `voice` or `sms`. |
**WorkerSid** | Pointer to **string** | The SID of the Worker that contains the WorkerChannel. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the WorkerChannel. |
**Url** | Pointer to **string** | The absolute URL of the WorkerChannel resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


