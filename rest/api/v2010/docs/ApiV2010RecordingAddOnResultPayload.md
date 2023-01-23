# ApiV2010RecordingAddOnResultPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that that we created to identify the Recording AddOnResult Payload resource. |
**AddOnResultSid** | Pointer to **string** | The SID of the AddOnResult to which the payload belongs. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource. |
**Label** | Pointer to **string** | The string provided by the vendor that describes the payload. |
**AddOnSid** | Pointer to **string** | The SID of the Add-on to which the result belongs. |
**AddOnConfigurationSid** | Pointer to **string** | The SID of the Add-on configuration. |
**ContentType** | Pointer to **string** | The MIME type of the payload. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**ReferenceSid** | Pointer to **string** | The SID of the recording to which the AddOnResult resource that contains the payload belongs. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


