# MessagingV1Broadcast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BroadcastSid** | **string** | Numeric ID indentifying individual Broadcast requests |[optional] 
**CreatedDate** | [**time.Time**](time.Time.md) | Timestamp of when the Broadcast was created |[optional] 
**UpdatedDate** | [**time.Time**](time.Time.md) | Timestamp of when the Broadcast was last updated |[optional] 
**BroadcastStatus** | **string** | Status of the Broadcast request. Valid values are None, Pending-Upload, Uploaded, Queued, Executing, Execution-Failure, Execution-Completed, Cancelation-Requested, and Canceled |[optional] 
**ExecutionDetails** | [**MessagingV1BroadcastExecutionDetails**](MessagingV1BroadcastExecutionDetails.md) |  |[optional] 
**ErrorsFile** | **string** | Path to a file detailing errors from Broadcast execution |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


