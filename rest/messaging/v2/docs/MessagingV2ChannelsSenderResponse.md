# MessagingV2ChannelsSenderResponse

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
**PendingDisplayName** | Pointer to **string** | WhatsApp only. The display name the most recent change applies to — awaiting Meta review, approved by Meta and awaiting re-registration, or, once `pending_display_name_status` is `COMPLETED`, the name now in effect (identical to `name`). Absent when no display name change has been made, and once a completed change stops being reported.  |
**PendingDisplayNameStatus** | Pointer to **string** | WhatsApp only. The status of the most recent display name change. `PENDING_REVIEW`, `APPROVED` and `DECLINED` are reported by Meta. `PIN_MISMATCH` and `REGISTRATION_FAILED` mean Meta approved the name but it could not be applied; `EXPIRED` means Meta's 14-day window to apply an approved name elapsed. In all three cases, re-submit the same `profile.name` to retry. `COMPLETED` means the name was approved and applied — `name` now returns it. A `COMPLETED` change is reported for 14 days after it completes and is absent afterwards, so treat its presence as \"recently completed\" rather than a permanent flag; use `pending_display_name_status_date` to tell how recent. Absent when no display name change has been made.  |
**PendingDisplayNameStatusDate** | Pointer to [**time.Time**](time.Time.md) | WhatsApp only. The date and time in UTC when `pending_display_name_status` last changed, specified in ISO 8601 format. Absent whenever `pending_display_name_status` is absent, so the three `pending_display_name*` fields are always present or absent together.  |
**Properties** | Pointer to [**MessagingV2Properties**](MessagingV2Properties.md) |  |
**OfflineReasons** | Pointer to [**[]MessagingV2Items**](MessagingV2Items.md) | The reasons why the sender is offline. |
**Compliance** | Pointer to [**MessagingV2RcsComplianceResponse**](MessagingV2RcsComplianceResponse.md) |  |
**Url** | Pointer to **string** | The URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


