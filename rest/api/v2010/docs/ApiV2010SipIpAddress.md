# ApiV2010SipIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource. |[optional] 
**FriendlyName** | **string** | A human readable descriptive text for this resource, up to 255 characters long. |[optional] 
**IpAddress** | **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. |[optional] 
**CidrPrefixLength** | **int** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. |[optional] 
**IpAccessControlListSid** | **string** | The unique id of the IpAccessControlList resource that includes this resource. |[optional] 
**DateCreated** | **string** | The date that this resource was created, given as GMT in RFC 2822 format. |[optional] 
**DateUpdated** | **string** | The date that this resource was last updated, given as GMT in RFC 2822 format. |[optional] 
**Uri** | **string** | The URI for this resource, relative to https://api.twilio.com |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


