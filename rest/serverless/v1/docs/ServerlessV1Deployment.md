# ServerlessV1Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Deployment resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Deployment resource. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Deployment resource is associated with. |
**EnvironmentSid** | Pointer to **string** | The SID of the Environment for the Deployment. |
**BuildSid** | Pointer to **string** | The SID of the Build for the deployment. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Deployment resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Deployment resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Deployment resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


