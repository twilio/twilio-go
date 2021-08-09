# ApiV2010Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**IceServers** | Pointer to [**[]ApiV2010AccountTokenIceServers**](ApiV2010AccountTokenIceServers.md) | An array representing the ephemeral credentials |
**Password** | Pointer to **string** | The temporary password used for authenticating |
**Ttl** | Pointer to **string** | The duration in seconds the credentials are valid |
**Username** | Pointer to **string** | The temporary username that uniquely identifies a Token |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


