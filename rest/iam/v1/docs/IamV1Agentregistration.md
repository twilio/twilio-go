# IamV1Agentregistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationSid** | **string** | The agent's identity in the organization. |
**FriendlyName** | Pointer to **string** | The agent's name, from its CIMD document. Null when there is none, and then the screens show the CIMD URL. It cannot be set or changed. |
**ClientUrl** | **string** | The agent's CIMD document URL, as registered. |
**Origin** | **string** | The host of the CIMD document URL. |
**Enabled** | **bool** | Whether users in the organization can connect the agent. |
**CreatedBy** | Pointer to **string** | The name of the user who registered the agent, looked up when the registration is read. Null when the registration was made before the user was recorded, or the name could not be looked up. A user who has since been deleted shows as (Deleted User). |
**CreatedBySid** | Pointer to **string** | The sid of the user who registered the agent. Null when the registration was made before the user was recorded. |
**DateCreated** | [**time.Time**](time.Time.md) |  |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


