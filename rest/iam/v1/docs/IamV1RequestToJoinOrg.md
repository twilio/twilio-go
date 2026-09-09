# IamV1RequestToJoinOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationSid** | **string** |  |[optional] 
**Email** | **string** | Email address of the requesting user |[optional] 
**FirstName** | **string** | First name of the requesting user |[optional] 
**LastName** | **string** | Last name of the requesting user |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | Date the join request was created |[optional] 
**Scope** | **string** | The SID of the entity requested to join (OR…, AC…, etc) |[optional] 
**ScopeType** | **string** | Type of the scope entity for the requested role |[optional] 
**ScopeName** | **string** | Scope Name e.g. Account or Organization friendly name |[optional] 
**Roles** | [**[]IamV1RequestToJoinOrgRoles**](IamV1RequestToJoinOrgRoles.md) | The requested roles |[optional] 
**Status** | **string** | Display status of the join request (e.g. Active, Expired), derived by the downstream service from the underlying notification's state and expiry. Does not map 1:1 to the notification's internal status. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


