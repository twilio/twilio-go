# MessagingV1Service

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AreaCodeGeomatch** | Pointer to **bool** | Whether to enable Area Code Geomatch on the Service Instance |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FallbackMethod** | Pointer to **string** | The HTTP method we use to call fallback_url |
**FallbackToLongCode** | Pointer to **bool** | Whether to enable Fallback to Long Code for messages sent through the Service instance |
**FallbackUrl** | Pointer to **string** | The URL that we call using fallback_method if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**InboundMethod** | Pointer to **string** | The HTTP method we use to call inbound_request_url |
**InboundRequestUrl** | Pointer to **string** | The URL we call using inbound_method when a message is received by any phone number or short code in the Service. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources |
**MmsConverter** | Pointer to **bool** | Whether to enable the MMS Converter for messages sent through the Service instance |
**ScanMessageContent** | Pointer to **string** | Reserved |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SmartEncoding** | Pointer to **bool** | Whether to enable Encoding for messages sent through the Service instance |
**StatusCallback** | Pointer to **string** | The URL we call to pass status updates about message delivery |
**StickySender** | Pointer to **bool** | Whether to enable Sticky Sender on the Service instance |
**SynchronousValidation** | Pointer to **bool** | Reserved |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**UseInboundWebhookOnNumber** | Pointer to **bool** | If enabled, the webhook url configured on the phone number will be used and will override the `inbound_request_url`/`fallback_url` url called when an inbound message is received. |
**ValidityPeriod** | Pointer to **int32** | How long, in seconds, messages sent from the Service are valid |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


