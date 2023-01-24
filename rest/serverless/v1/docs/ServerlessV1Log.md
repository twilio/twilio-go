# ServerlessV1Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Log resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Log resource. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Log resource is associated with. |
**EnvironmentSid** | Pointer to **string** | The SID of the environment in which the log occurred. |
**BuildSid** | Pointer to **string** | The SID of the build that corresponds to the log. |
**DeploymentSid** | Pointer to **string** | The SID of the deployment that corresponds to the log. |
**FunctionSid** | Pointer to **string** | The SID of the function whose invocation produced the log. |
**RequestSid** | Pointer to **string** | The SID of the request associated with the log. |
**Level** | Pointer to [**string**](LogEnumLevel.md) |  |
**Message** | Pointer to **string** | The log message. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Log resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Log resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


