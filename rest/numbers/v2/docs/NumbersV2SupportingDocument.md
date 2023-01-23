# NumbersV2SupportingDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string created by Twilio to identify the Supporting Document resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Document resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**MimeType** | Pointer to **string** | The image type uploaded in the Supporting Document container. |
**Status** | Pointer to [**string**](SupportingDocumentEnumStatus.md) |  |
**FailureReason** | Pointer to **string** | The failure reason of the Supporting Document Resource. |
**Type** | Pointer to **string** | The type of the Supporting Document. |
**Attributes** | Pointer to **interface{}** | The set of parameters that are the attributes of the Supporting Documents resource which are listed in the Supporting Document Types. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Supporting Document resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


