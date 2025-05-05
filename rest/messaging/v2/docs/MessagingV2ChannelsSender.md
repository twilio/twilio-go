# MessagingV2ChannelsSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Sender. |
**Status** | Pointer to [**string**](ChannelsSenderEnumStatus.md) |  |
**SenderId** | Pointer to **string** | The ID of this Sender prefixed with the channel, e.g., `whatsapp:E.164` |
**Configuration** | Pointer to [**MessagingV2Configuration**](MessagingV2Configuration.md) |  |
**Webhook** | Pointer to [**MessagingV2Webhook**](MessagingV2Webhook.md) |  |
**Profile** | Pointer to [**MessagingV2Profile**](MessagingV2Profile.md) |  |
**Properties** | Pointer to [**MessagingV2Properties**](MessagingV2Properties.md) |  |
**OfflineReasons** | Pointer to [**[]MessagingV2Items**](MessagingV2Items.md) | Reasons why the sender is offline., e.g., [{\"code\": \"21211400\", \"message\": \"Whatsapp business account is banned by provider {provider_name} | Credit line is assigned to another BSP\", \"more_info\": \"https://www.twilio.com/docs/errors/21211400\"}] |
**Url** | Pointer to **string** | The URL of this resource, relative to `https://messaging.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


