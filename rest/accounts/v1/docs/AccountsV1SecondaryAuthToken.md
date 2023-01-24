# AccountsV1SecondaryAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the secondary Auth Token was created for. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in UTC when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in UTC when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**SecondaryAuthToken** | Pointer to **string** | The generated secondary Auth Token that can be used to authenticate future API requests. |
**Url** | Pointer to **string** | The URI for this resource, relative to `https://accounts.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


