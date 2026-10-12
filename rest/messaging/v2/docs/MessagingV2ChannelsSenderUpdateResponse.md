# MessagingV2ChannelsSenderUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The SID of the sender. |
**Status** | Pointer to [**string**](ChannelsSenderEnumStatus.md) |  |
**SenderId** | Pointer to **string** | The ID of the sender in `whatsapp:<E.164_PHONE_NUMBER>` format. |
**FriendlyName** | Pointer to **string** | Optional display label for the sender in the Twilio Console. |
**Configuration** | Pointer to [**MessagingV2Configuration**](MessagingV2Configuration.md) |  |
**Webhook** | Pointer to [**MessagingV2Webhook**](MessagingV2Webhook.md) |  |
**Profile** | Pointer to [**MessagingV2ProfileGenericResponse**](MessagingV2ProfileGenericResponse.md) |  |
**DisplayNameStatus** | Pointer to **string** | WhatsApp only. The outcome of the display name operation in this request. Present only when the request included `profile.name`. `updating` — accepted; either submitted to Meta for review, or, when Meta had already approved this exact name, routed straight to re-registration. `no_change` — the name already matches the sender's active display name; nothing was submitted to Meta. `pending_review` — the same name is already under review at Meta; the existing request continues unchanged. `error` — the display name could not be processed, while other profile fields in the same request were still applied. Returned with a 202 and carries no error code or message. This covers every failure mode, including the case where Meta accepted the name but tracking could not be started — poll `pending_display_name_status` to establish the real state rather than assuming the name was rejected. When `profile.name` is the only field in the request, the failure is returned as an error response with a specific code instead of this status.  |
**Properties** | Pointer to [**MessagingV2Properties**](MessagingV2Properties.md) |  |
**OfflineReasons** | Pointer to [**[]MessagingV2Items**](MessagingV2Items.md) | The reasons why the sender is offline. |
**Compliance** | Pointer to [**MessagingV2RcsComplianceResponse**](MessagingV2RcsComplianceResponse.md) |  |
**Url** | Pointer to **string** | The URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


