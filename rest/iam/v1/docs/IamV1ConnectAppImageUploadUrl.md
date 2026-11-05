# IamV1ConnectAppImageUploadUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PresignedUrl** | **string** | Presigned S3 URL to upload the image to with a PUT. The request must send exactly contentType and the requested FileSize. Valid until expiresAt. |[optional] 
**ImageKey** | **string** | The key of the uploaded image. Pass it as imageKey when creating or updating the Connect App. |[optional] 
**ContentType** | **string** | The Content-Type the upload must be sent with. |[optional] 
**ExpiresAt** | [**time.Time**](time.Time.md) | When presignedUrl stops accepting uploads (5 minutes after it is issued). |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


