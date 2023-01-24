# ApiV2010SipCredentialList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) that owns this resource. |
**DateCreated** | Pointer to **string** | The date that this resource was created, given as GMT in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**DateUpdated** | Pointer to **string** | The date that this resource was last updated, given as GMT in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**FriendlyName** | Pointer to **string** | A human readable descriptive text that describes the CredentialList, up to 64 characters long. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of credentials associated with this credential list. |
**Uri** | Pointer to **string** | The URI for this resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


