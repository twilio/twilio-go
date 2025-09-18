# IntelligenceV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account the Service belongs to. |
**AutoRedaction** | Pointer to **bool** | Instructs the Speech Recognition service to automatically redact PII from all transcripts made on this service. |
**MediaRedaction** | Pointer to **bool** | Instructs the Speech Recognition service to automatically redact PII from all transcripts media made on this service. The auto_redaction flag must be enabled, results in error otherwise. |
**AutoTranscribe** | Pointer to **bool** | Instructs the Speech Recognition service to automatically transcribe all recordings made on the account. |
**DataLogging** | Pointer to **bool** | Data logging allows Twilio to improve the quality of the speech recognition & language understanding services through using customer data to refine, fine tune and evaluate machine learning models. Note: Data logging cannot be activated via API, only via www.twilio.com, as it requires additional consent. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Service was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Service was updated, given in ISO 8601 format. |
**FriendlyName** | Pointer to **string** | A human readable description of this resource, up to 64 characters. |
**LanguageCode** | Pointer to **string** | The language code set during Service creation determines the Transcription language for all call recordings processed by that Service. The default is en-US if no language code is set. A Service can only support one language code, and it cannot be updated once it's set. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Service. |
**UniqueName** | Pointer to **string** | Provides a unique and addressable name to be assigned to this Service, assigned by the developer, to be optionally used in addition to SID. |
**Url** | Pointer to **string** | The URL of this resource. |
**WebhookUrl** | Pointer to **string** | The URL Twilio will request when executing the Webhook. |
**WebhookHttpMethod** | Pointer to [**string**](ServiceEnumHttpMethod.md) |  |
**ReadOnlyAttachedOperatorSids** | Pointer to **[]string** | Operator sids attached to this service, read only |
**Version** | **int** | The version number of this Service. |[optional] [default to 0]
**EncryptionCredentialSid** | Pointer to **string** | The unique SID identifier of the Public Key resource used to encrypt the sentences and operator results. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


