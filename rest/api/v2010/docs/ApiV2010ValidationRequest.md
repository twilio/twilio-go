# ApiV2010ValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the Caller ID. |
**CallSid** | Pointer to **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Caller ID is associated with. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**PhoneNumber** | Pointer to **string** | The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**ValidationCode** | Pointer to **string** | The 6 digit validation code that someone must enter to validate the Caller ID  when `phone_number` is called. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


