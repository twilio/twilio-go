# VerifyV2Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A string that uniquely identifies this Bucket. |
**RateLimitSid** | Pointer to **string** | Rate Limit Sid. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Max** | Pointer to **int** | Max number of requests. |
**Interval** | Pointer to **int** | Number of seconds that the rate limit will be enforced over. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


