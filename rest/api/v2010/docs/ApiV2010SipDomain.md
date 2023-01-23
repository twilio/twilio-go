# ApiV2010SipDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource. |
**ApiVersion** | Pointer to **string** | The API version used to process the call. |
**AuthType** | Pointer to **string** | The types of authentication you have mapped to your domain. Can be: `IP_ACL` and `CREDENTIAL_LIST`. If you have both defined for your domain, both will be returned in a comma delimited string. If `auth_type` is not defined, the domain will not be able to receive any traffic. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DomainName** | Pointer to **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \"-\" and must end with `sip.twilio.com`. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the SipDomain resource. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`. |
**VoiceFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML requested from `voice_url`. |
**VoiceMethod** | Pointer to **string** | The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`. |
**VoiceStatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call `voice_status_callback_url`. Either `GET` or `POST`. |
**VoiceStatusCallbackUrl** | Pointer to **string** | The URL that we call to pass status parameters (such as call ended) to your application. |
**VoiceUrl** | Pointer to **string** | The URL we call using the `voice_method` when the domain receives a call. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of mapping resources associated with the SIP Domain resource identified by their relative URIs. |
**SipRegistration** | Pointer to **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. |
**EmergencyCallingEnabled** | Pointer to **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. |
**Secure** | Pointer to **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. |
**ByocTrunkSid** | Pointer to **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. |
**EmergencyCallerSid** | Pointer to **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


