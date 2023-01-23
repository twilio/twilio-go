# ServerlessV1Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Variable resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Variable resource. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Variable resource is associated with. |
**EnvironmentSid** | Pointer to **string** | The SID of the Environment in which the Variable exists. |
**Key** | Pointer to **string** | A string by which the Variable resource can be referenced. |
**Value** | Pointer to **string** | A string that contains the actual value of the Variable. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Variable resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Variable resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Variable resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


