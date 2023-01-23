# VoiceV1IpRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**IpAddress** | **string** | An IP address in dotted decimal notation, IPv4 only. |[optional] 
**CidrPrefixLength** | **int** | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


