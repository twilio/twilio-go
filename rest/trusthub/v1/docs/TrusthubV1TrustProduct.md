# TrusthubV1TrustProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Trust Product resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Trust Product resource. |
**PolicySid** | Pointer to **string** | The unique string of the policy that is associated with the Trust Product resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Status** | Pointer to [**string**](TrustProductEnumStatus.md) |  |
**ValidUntil** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format until which the resource will be valid. |
**Email** | Pointer to **string** | The email address that will receive updates when the Trust Product resource changes status. |
**StatusCallback** | Pointer to **string** | The URL we call to inform your application of status changes. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Trust Product resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Assigned Items of the Trust Product resource. |
**Errors** | Pointer to **[]interface{}** | The error codes associated with the rejection of the Trust Product. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


