# ServerlessV1Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Environment resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Environment resource. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Environment resource is associated with. |
**BuildSid** | Pointer to **string** | The SID of the build deployed in the environment. |
**UniqueName** | Pointer to **string** | A user-defined string that uniquely identifies the Environment resource. |
**DomainSuffix** | Pointer to **string** | A URL-friendly name that represents the environment and forms part of the domain name. |
**DomainName** | Pointer to **string** | The domain name for all Functions and Assets deployed in the Environment, using the Service unique name, a randomly-generated Service suffix, and an optional Environment domain suffix. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Environment resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Environment resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Environment resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Environment resource's nested resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


