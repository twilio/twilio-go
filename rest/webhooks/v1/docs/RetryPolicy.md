# RetryPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryCount** | **int** | Number of retry attempts (default: 0, no retries) |[default to 0]
**RetryOn** | **[]string** | Conditions that trigger retries: - HTTP_4XX: Retry on 4xx HTTP status codes - HTTP_5XX: Retry on 5xx HTTP status codes - CONNECTION_TIMEOUT: Retry on TCP/TLS connection failures - READ_TIMEOUT: Retry when server doesn't respond in time - ALL: Retry on all failure types  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


