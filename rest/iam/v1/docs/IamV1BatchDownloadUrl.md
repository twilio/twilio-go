# IamV1BatchDownloadUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadUrl** | **string** | A pre-signed, time-limited S3 URL for the CSV. The lifetime is short and server-configured; read `expiresAt` rather than assuming a fixed duration. |
**ExpiresAt** | [**time.Time**](time.Time.md) | The time the pre-signed URL stops being valid, given in RFC 3339 format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


