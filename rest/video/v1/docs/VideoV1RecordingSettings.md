# VideoV1RecordingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the RecordingSettings resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource and show the user in the console |
**AwsCredentialsSid** | Pointer to **string** | The SID of the stored Credential resource. |
**AwsS3Url** | Pointer to **string** | The URL of the AWS S3 bucket where the recordings are stored. We only support DNS-compliant URLs like `https://documentation-example-twilio-bucket/recordings`, where `recordings` is the path in which you want the recordings to be stored. This URL accepts only URI-valid characters, as described in the <a href='https://tools.ietf.org/html/rfc3986#section-2'>RFC 3986</a>. |
**AwsStorageEnabled** | Pointer to **bool** | Whether all recordings are written to the `aws_s3_url`. When `false`, all recordings are stored in our cloud. |
**EncryptionKeySid** | Pointer to **string** | The SID of the Public Key resource used for encryption. |
**EncryptionEnabled** | Pointer to **bool** | Whether all recordings are stored in an encrypted form. The default is `false`. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


