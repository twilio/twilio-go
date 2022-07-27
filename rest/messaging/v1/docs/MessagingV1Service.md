# MessagingV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**InboundRequestUrl** | Pointer to **string** | The URL we call using inbound_method when a message is received by any phone number or short code in the Service. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled. |
**InboundMethod** | Pointer to **string** | The HTTP method we use to call inbound_request_url |
**FallbackUrl** | Pointer to **string** | The URL that we call using fallback_method if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled. |
**FallbackMethod** | Pointer to **string** | The HTTP method we use to call fallback_url |
**StatusCallback** | Pointer to **string** | The URL we call to pass status updates about message delivery |
**StickySender** | Pointer to **bool** | Whether to enable Sticky Sender on the Service instance |
**MmsConverter** | Pointer to **bool** | Whether to enable the MMS Converter for messages sent through the Service instance |
**SmartEncoding** | Pointer to **bool** | Whether to enable Encoding for messages sent through the Service instance |
**ScanMessageContent** | Pointer to [**string**](ServiceEnumScanMessageContent.md) |  |
**FallbackToLongCode** | Pointer to **bool** | Whether to enable Fallback to Long Code for messages sent through the Service instance |
**AreaCodeGeomatch** | Pointer to **bool** | Whether to enable Area Code Geomatch on the Service Instance |
**SynchronousValidation** | Pointer to **bool** | Reserved |
**ValidityPeriod** | Pointer to **int** | How long, in seconds, messages sent from the Service are valid |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources |
**Usecase** | Pointer to **string** | A string describing the scenario in which the Messaging Service will be used |
**UsAppToPersonRegistered** | Pointer to **bool** | Whether US A2P campaign is registered for this Service. |
**UseInboundWebhookOnNumber** | Pointer to **bool** | If enabled, the webhook url configured on the phone number will be used and will override the `inbound_request_url`/`fallback_url` url called when an inbound message is received. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


