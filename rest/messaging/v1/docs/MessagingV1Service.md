# MessagingV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**InboundRequestUrl** | **string** | The URL we call using inbound_method when a message is received by any phone number or short code in the Service. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled. |[optional] 
**InboundMethod** | **string** | The HTTP method we use to call inbound_request_url |[optional] 
**FallbackUrl** | **string** | The URL that we call using fallback_method if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. This field will be overridden if the `use_inbound_webhook_on_number` field is enabled. |[optional] 
**FallbackMethod** | **string** | The HTTP method we use to call fallback_url |[optional] 
**StatusCallback** | **string** | The URL we call to pass status updates about message delivery |[optional] 
**StickySender** | **bool** | Whether to enable Sticky Sender on the Service instance |[optional] 
**MmsConverter** | **bool** | Whether to enable the MMS Converter for messages sent through the Service instance |[optional] 
**SmartEncoding** | **bool** | Whether to enable Encoding for messages sent through the Service instance |[optional] 
**ScanMessageContent** | Pointer to [**string**](ServiceEnumScanMessageContent.md) |  |
**FallbackToLongCode** | **bool** | Whether to enable Fallback to Long Code for messages sent through the Service instance |[optional] 
**AreaCodeGeomatch** | **bool** | Whether to enable Area Code Geomatch on the Service Instance |[optional] 
**SynchronousValidation** | **bool** | Reserved |[optional] 
**ValidityPeriod** | **int** | How long, in seconds, messages sent from the Service are valid |[optional] 
**Url** | **string** | The absolute URL of the Service resource |[optional] 
**Links** | **map[string]interface{}** | The absolute URLs of related resources |[optional] 
**Usecase** | **string** | A string describing the scenario in which the Messaging Service will be used |[optional] 
**UsAppToPersonRegistered** | **bool** | Whether US A2P campaign is registered for this Service. |[optional] 
**UseInboundWebhookOnNumber** | **bool** | If enabled, the webhook url configured on the phone number will be used and will override the `inbound_request_url`/`fallback_url` url called when an inbound message is received. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


