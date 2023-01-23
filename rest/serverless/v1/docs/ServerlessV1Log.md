# ServerlessV1Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the Log resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the Log resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the Log resource is associated with |[optional] 
**EnvironmentSid** | **string** | The SID of the environment in which the log occurred |[optional] 
**BuildSid** | **string** | The SID of the build that corresponds to the log |[optional] 
**DeploymentSid** | **string** | The SID of the deployment that corresponds to the log |[optional] 
**FunctionSid** | **string** | The SID of the function whose invocation produced the log |[optional] 
**RequestSid** | **string** | The SID of the request associated with the log |[optional] 
**Level** | Pointer to [**string**](LogEnumLevel.md) |  |
**Message** | **string** | The log message |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Log resource was created |[optional] 
**Url** | **string** | The absolute URL of the Log resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


