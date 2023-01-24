# ApiV2010Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Token resource. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**IceServers** | Pointer to [**[]ApiV2010AccountTokenIceServers**](ApiV2010AccountTokenIceServers.md) | An array representing the ephemeral credentials and the STUN and TURN server URIs. |
**Password** | Pointer to **string** | The temporary password that the username will use when authenticating with Twilio. |
**Ttl** | Pointer to **string** | The duration in seconds for which the username and password are valid. |
**Username** | Pointer to **string** | The temporary username that uniquely identifies a Token. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


