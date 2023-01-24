# ApiV2010SipIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**AccountSid** | Pointer to **string** | The unique id of the Account that is responsible for this resource. |
**FriendlyName** | Pointer to **string** | A human readable descriptive text for this resource, up to 255 characters long. |
**IpAddress** | Pointer to **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. |
**CidrPrefixLength** | Pointer to **int** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. |
**IpAccessControlListSid** | Pointer to **string** | The unique id of the IpAccessControlList resource that includes this resource. |
**DateCreated** | Pointer to **string** | The date that this resource was created, given as GMT in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**DateUpdated** | Pointer to **string** | The date that this resource was last updated, given as GMT in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**Uri** | Pointer to **string** | The URI for this resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


