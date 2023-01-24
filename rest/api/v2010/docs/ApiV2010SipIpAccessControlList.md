# ApiV2010SipIpAccessControlList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**AccountSid** | Pointer to **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) that owns this resource. |
**FriendlyName** | Pointer to **string** | A human readable descriptive text, up to 255 characters long. |
**DateCreated** | Pointer to **string** | The date that this resource was created, given as GMT in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**DateUpdated** | Pointer to **string** | The date that this resource was last updated, given as GMT in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of the IpAddress resources associated with this IP access control list resource. |
**Uri** | Pointer to **string** | The URI for this resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


