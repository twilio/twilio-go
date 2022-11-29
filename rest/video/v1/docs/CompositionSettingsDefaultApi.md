# CompositionSettingsDefaultApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompositionSettings**](CompositionSettingsDefaultApi.md#CreateCompositionSettings) | **Post** /v1/CompositionSettings/Default | 
[**FetchCompositionSettings**](CompositionSettingsDefaultApi.md#FetchCompositionSettings) | **Get** /v1/CompositionSettings/Default | 



## CreateCompositionSettings

> VideoV1CompositionSettings CreateCompositionSettings(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCompositionSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource and show to the user in the console
**AwsCredentialsSid** | **string** | The SID of the stored Credential resource.
**EncryptionKeySid** | **string** | The SID of the Public Key resource to use for encryption.
**AwsS3Url** | **string** | The URL of the AWS S3 bucket where the compositions should be stored. We only support DNS-compliant URLs like `https://documentation-example-twilio-bucket/compositions`, where `compositions` is the path in which you want the compositions to be stored. This URL accepts only URI-valid characters, as described in the <a href='https://tools.ietf.org/html/rfc3986#section-2'>RFC 3986</a>.
**AwsStorageEnabled** | **bool** | Whether all compositions should be written to the `aws_s3_url`. When `false`, all compositions are stored in our cloud.
**EncryptionEnabled** | **bool** | Whether all compositions should be stored in an encrypted form. The default is `false`.

### Return type

[**VideoV1CompositionSettings**](VideoV1CompositionSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCompositionSettings

> VideoV1CompositionSettings FetchCompositionSettings(ctx, )





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchCompositionSettingsParams struct


### Return type

[**VideoV1CompositionSettings**](VideoV1CompositionSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

