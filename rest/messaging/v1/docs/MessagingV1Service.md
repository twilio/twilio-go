# MessagingV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Service resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**InboundRequestUrl** | Pointer to **string** | The URL we call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. All messages sent to the Twilio phone number or short code will not be logged and received on the Account. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `inbound_request_url` defined for the Messaging Service. |
**InboundMethod** | Pointer to **string** | The HTTP method we use to call `inbound_request_url`. Can be `GET` or `POST`. |
**FallbackUrl** | Pointer to **string** | The URL that we call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. If the `use_inbound_webhook_on_number` field is enabled then the webhook url defined on the phone number will override the `fallback_url` defined for the Messaging Service. |
**FallbackMethod** | Pointer to **string** | The HTTP method we use to call `fallback_url`. Can be: `GET` or `POST`. |
**StatusCallback** | Pointer to **string** | The URL we call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery. |
**StickySender** | Pointer to **bool** | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance. |
**MmsConverter** | Pointer to **bool** | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance. |
**SmartEncoding** | Pointer to **bool** | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance. |
**ScanMessageContent** | Pointer to [**string**](ServiceEnumScanMessageContent.md) |  |
**FallbackToLongCode** | Pointer to **bool** | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance. |
**AreaCodeGeomatch** | Pointer to **bool** | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance. |
**SynchronousValidation** | Pointer to **bool** | Reserved. |
**ValidityPeriod** | Pointer to **int** | How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`. |
**Url** | Pointer to **string** | The absolute URL of the Service resource. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources. |
**Usecase** | Pointer to **string** | A string that describes the scenario in which the Messaging Service will be used. Examples: [notification, marketing, verification, poll ..] |
**UsAppToPersonRegistered** | Pointer to **bool** | Whether US A2P campaign is registered for this Service. |
**UseInboundWebhookOnNumber** | Pointer to **bool** | A boolean value that indicates either the webhook url configured on the phone number will be used or `inbound_request_url`/`fallback_url` url will be called when a message is received from the phone number. If this field is enabled then the webhook url defined on the phone number will override the `inbound_request_url`/`fallback_url` defined for the Messaging Service. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


