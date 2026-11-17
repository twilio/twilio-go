# CreateComplianceSignalsCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signals** | **[]string** | The signals to evaluate. Only `QUIET_HOURS` is supported. |
**To** | **string** | The recipient phone number in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). |
**Body** | **string** | The content of the message, at most 1600 characters. Required if `useCase` is not provided. |[optional] 
**UseCase** | **string** | The intent of the message, one of `OTP`, `NOTIFICATIONS`, `FRAUD`, `SECURITY`, `CUSTOMERCARE`, `DELIVERY`, `EDUCATION`, `EVENTS`, `POLLING`, `ANNOUNCEMENTS` or `MARKETING`. Required if `body` is not provided. It is used instead of the body when both are given. |[optional] 
**Channel** | **string** | The channel the message is sent on, one of `SMS`, `MMS`, `RCS` or `CALL`. Defaults to `SMS`. |[optional] 
**CountryIsoCode** | **string** | The ISO 3166-1 alpha-2 country of the recipient. Only `US` is supported. |
**ZipCode** | **string** | The US zip code of the recipient, 5 digits or ZIP+4. When it is not given or not known, the location is inferred from the phone number. |[optional] 
**SendAt** | **string** | An [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339) timestamp, in UTC or with an offset, to evaluate a scheduled delivery instead of the current time. It must not be in the past. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


