# ApiV2010AccountCallCallRecording

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ApiVersion** | Pointer to **string** | The API version used to make the recording |
**CallSid** | Pointer to **string** | The SID of the Call the resource is associated with |
**Channels** | Pointer to **int32** | The number of channels in the final recording file |
**ConferenceSid** | Pointer to **string** | The Conference SID that identifies the conference associated with the recording |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**Duration** | Pointer to **string** | The length of the recording in seconds |
**EncryptionDetails** | Pointer to **map[string]interface{}** | How to decrypt the recording. |
**ErrorCode** | Pointer to **int32** | More information about why the recording is missing, if status is `absent`. |
**Price** | Pointer to **float32** | The one-time cost of creating the recording. |
**PriceUnit** | Pointer to **string** | The currency used in the price property. |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Source** | Pointer to **string** | How the recording was created |
**StartTime** | Pointer to **string** | The start time of the recording, given in RFC 2822 format |
**Status** | Pointer to **string** | The status of the recording |
**Track** | Pointer to **string** | The recorded track. Can be: `inbound`, `outbound`, or `both`. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


