# VoiceV1ConnectionPolicyTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Target** | **string** | The SIP address you want Twilio to route your calls to |[optional] 
**Priority** | **int** | The relative importance of the target |[optional] 
**Weight** | **int** | The value that determines the relative load the Target should receive compared to others with the same priority |[optional] 
**Enabled** | **bool** | Whether the target is enabled |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


