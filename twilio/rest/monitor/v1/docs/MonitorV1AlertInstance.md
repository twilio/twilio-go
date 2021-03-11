# MonitorV1AlertInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AlertText** | Pointer to **string** | The text of the alert |
**ApiVersion** | Pointer to **string** | The API version used when the alert was generated |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateGenerated** | Pointer to [**time.Time**](time.Time.md) | The date and time when the alert was generated specified in ISO 8601 format |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**ErrorCode** | Pointer to **string** | The error code for the condition that generated the alert |
**LogLevel** | Pointer to **string** | The log level |
**MoreInfo** | Pointer to **string** | The URL of the page in our Error Dictionary with more information about the error condition |
**RequestHeaders** | Pointer to **string** | The request headers of the request that generated the alert |
**RequestMethod** | Pointer to **string** | The method used by the request that generated the alert |
**RequestUrl** | Pointer to **string** | The URL of the request that generated the alert |
**RequestVariables** | Pointer to **string** | The variables passed in the request that generated the alert |
**ResourceSid** | Pointer to **string** | The SID of the resource for which the alert was generated |
**ResponseBody** | Pointer to **string** | The response body of the request that generated the alert |
**ResponseHeaders** | Pointer to **string** | The response headers of the request that generated the alert |
**ServiceSid** | Pointer to **string** | The SID of the service or resource that generated the alert |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Alert resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


