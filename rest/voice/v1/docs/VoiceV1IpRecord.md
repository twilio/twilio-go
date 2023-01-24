# VoiceV1IpRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IP Record resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the IP Record resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**IpAddress** | Pointer to **string** | An IP address in dotted decimal notation, IPv4 only. |
**CidrPrefixLength** | Pointer to **int** | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


