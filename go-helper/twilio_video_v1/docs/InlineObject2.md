# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsCredentialsSid** | **string** | The SID of the stored Credential resource. | [optional] 
**AwsS3Url** | **string** | The URL of the AWS S3 bucket where the compositions should be stored. We only support DNS-compliant URLs like &#x60;https://documentation-example-twilio-bucket/compositions&#x60;, where &#x60;compositions&#x60; is the path in which you want the compositions to be stored. This URL accepts only URI-valid characters, as described in the &lt;a href&#x3D;&#39;https://tools.ietf.org/html/rfc3986#section-2&#39;&gt;RFC 3986&lt;/a&gt;. | [optional] 
**AwsStorageEnabled** | **bool** | Whether all compositions should be written to the &#x60;aws_s3_url&#x60;. When &#x60;false&#x60;, all compositions are stored in our cloud. | [optional] 
**EncryptionEnabled** | **bool** | Whether all compositions should be stored in an encrypted form. The default is &#x60;false&#x60;. | [optional] 
**EncryptionKeySid** | **string** | The SID of the Public Key resource to use for encryption. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource and show to the user in the console | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


