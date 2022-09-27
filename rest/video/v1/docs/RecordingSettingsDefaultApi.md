# RecordingSettingsDefaultApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecordingSettings**](RecordingSettingsDefaultApi.md#CreateRecordingSettings) | **Post** /v1/RecordingSettings/Default | 
[**FetchRecordingSettings**](RecordingSettingsDefaultApi.md#FetchRecordingSettings) | **Get** /v1/RecordingSettings/Default | 



## CreateRecordingSettings

> VideoV1RecordingSettings CreateRecordingSettings(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRecordingSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource and be shown to users in the console
**AwsCredentialsSid** | **string** | The SID of the stored Credential resource.
**EncryptionKeySid** | **string** | The SID of the Public Key resource to use for encryption.
**AwsS3Url** | **string** | The URL of the AWS S3 bucket where the recordings should be stored. We only support DNS-compliant URLs like &#x60;https://documentation-example-twilio-bucket/recordings&#x60;, where &#x60;recordings&#x60; is the path in which you want the recordings to be stored. This URL accepts only URI-valid characters, as described in the &lt;a href&#x3D;&#39;https://tools.ietf.org/html/rfc3986#section-2&#39;&gt;RFC 3986&lt;/a&gt;.
**AwsStorageEnabled** | **bool** | Whether all recordings should be written to the &#x60;aws_s3_url&#x60;. When &#x60;false&#x60;, all recordings are stored in our cloud.
**EncryptionEnabled** | **bool** | Whether all recordings should be stored in an encrypted form. The default is &#x60;false&#x60;.

### Return type

[**VideoV1RecordingSettings**](VideoV1RecordingSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingSettings

> VideoV1RecordingSettings FetchRecordingSettings(ctx, )





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingSettingsParams struct


### Return type

[**VideoV1RecordingSettings**](VideoV1RecordingSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

