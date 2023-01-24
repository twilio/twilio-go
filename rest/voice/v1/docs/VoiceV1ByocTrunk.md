# VoiceV1ByocTrunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the BYOC Trunk resource. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the BYOC Trunk resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**VoiceUrl** | Pointer to **string** | The URL we call using the `voice_method` when the BYOC Trunk receives a call. |
**VoiceMethod** | Pointer to **string** | The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`. |
**VoiceFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML requested from `voice_url`. |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`. |
**StatusCallbackUrl** | Pointer to **string** | The URL that we call to pass status parameters (such as call ended) to your application. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call `status_callback_url`. Either `GET` or `POST`. |
**CnamLookupEnabled** | Pointer to **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. |
**ConnectionPolicySid** | Pointer to **string** | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. |
**FromDomainSid** | Pointer to **string** | The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \"call back\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \"sip.twilio.com\". |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


