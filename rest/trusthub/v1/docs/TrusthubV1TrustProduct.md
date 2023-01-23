# TrusthubV1TrustProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource. |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**PolicySid** | **string** | The unique string of a policy. |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Status** | Pointer to [**string**](TrustProductEnumStatus.md) |  |
**ValidUntil** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource will be valid until. |[optional] 
**Email** | **string** | The email address |[optional] 
**StatusCallback** | **string** | The URL we call to inform your application of status changes. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Customer-Profile resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of the Assigned Items of the Customer-Profile resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


