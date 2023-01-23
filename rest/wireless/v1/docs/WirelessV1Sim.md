# WirelessV1Sim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Sim resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource belongs. |
**RatePlanSid** | Pointer to **string** | The SID of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource is assigned. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Sim resource. |
**Iccid** | Pointer to **string** | The [ICCID](https://en.wikipedia.org/wiki/SIM_card#ICCID) associated with the SIM. |
**EId** | Pointer to **string** | Deprecated. |
**Status** | Pointer to [**string**](SimEnumStatus.md) |  |
**ResetStatus** | Pointer to [**string**](SimEnumResetStatus.md) |  |
**CommandsCallbackUrl** | Pointer to **string** | The URL we call using the `commands_callback_method` when the SIM originates a machine-to-machine [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. |
**CommandsCallbackMethod** | Pointer to **string** | The HTTP method we use to call `commands_callback_url`.  Can be: `POST` or `GET`. Default is `POST`. |
**SmsFallbackMethod** | Pointer to **string** | Deprecated. |
**SmsFallbackUrl** | Pointer to **string** | Deprecated. |
**SmsMethod** | Pointer to **string** | Deprecated. |
**SmsUrl** | Pointer to **string** | Deprecated. |
**VoiceFallbackMethod** | Pointer to **string** | Deprecated. The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`. Default is `POST`. |
**VoiceFallbackUrl** | Pointer to **string** | Deprecated. The URL we call using the `voice_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `voice_url`. |
**VoiceMethod** | Pointer to **string** | Deprecated. The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`. Default is `POST`. |
**VoiceUrl** | Pointer to **string** | Deprecated. The URL we call using the `voice_method` when the SIM-connected device makes a voice call. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Sim resource was last updated specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related subresources. |
**IpAddress** | Pointer to **string** | Deprecated. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


