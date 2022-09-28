# ApiV2010Recording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ApiVersion** | Pointer to **string** | The API version used during the recording. |
**CallSid** | Pointer to **string** | The SID of the Call the resource is associated with |
**ConferenceSid** | Pointer to **string** | The unique ID for the conference associated with the recording. |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**StartTime** | Pointer to **string** | The start time of the recording, given in RFC 2822 format |
**Duration** | Pointer to **string** | The length of the recording in seconds. |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Price** | Pointer to **string** | The one-time cost of creating the recording. |
**PriceUnit** | Pointer to **string** | The currency used in the price property. |
**Status** | Pointer to [**string**](RecordingEnumStatus.md) |  |
**Channels** | Pointer to **int** | The number of channels in the final recording file as an integer. |
**Source** | Pointer to [**string**](RecordingEnumSource.md) |  |
**ErrorCode** | Pointer to **int** | More information about why the recording is missing, if status is `absent`. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**EncryptionDetails** | Pointer to **interface{}** | How to decrypt the recording. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs |
**MediaUrl** | Pointer to **string** | The URL of the media file. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


