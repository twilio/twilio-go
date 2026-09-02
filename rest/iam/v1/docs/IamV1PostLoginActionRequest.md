# IamV1PostLoginActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | The Whitney event name (e.g. auth_flow.checkpoint) |[optional] 
**AccountSid** | **string** | The Twilio account SID associated with the login |[optional] 
**NodeId** | **string** | The Whitney auth flow node identifier for this action |[optional] 
**Timestamp** | **int64** | Unix timestamp (seconds) of the login event |[optional] 
**UserContext** | [**IamV1PostLoginActionUserContext**](IamV1PostLoginActionUserContext.md) |  |
**RequestContext** | [**IamV1PostLoginActionRequestContext**](IamV1PostLoginActionRequestContext.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


