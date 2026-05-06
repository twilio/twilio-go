# ProfileEventRequestEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**time.Time**](time.Time.md) | The time the event occurred in ISO8601 format with millisecond resolution. Defaults to received time if not provided. |
**Lifecycle** | **string** | The lifecycle event type of the communication. |
**ConversationId** | **string** | The unique identifier for the conversation. |
**CommunicationId** | **string** | The unique identifier for the communication. |[optional] 
**CommunicationType** | **string** | The communication channel type. |
**CommunicationStatus** | **string** | The lifecycle status of the communication. |[optional] 
**Direction** | **string** | The direction of the communication. |[optional] 
**Sender** | [**CommunicationLifecycleEventSender**](CommunicationLifecycleEventSender.md) |  |[optional] 
**Recipient** | [**CommunicationLifecycleEventRecipient**](CommunicationLifecycleEventRecipient.md) |  |[optional] 
**ErrorCode** | **string** | Error code for FAILED communication status. |[optional] 
**ErrorMessage** | **string** | Error message for FAILED communication status. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


