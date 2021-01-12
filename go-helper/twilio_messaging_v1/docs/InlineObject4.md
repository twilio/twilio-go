# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaCodeGeomatch** | **bool** | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance. | [optional] 
**FallbackMethod** | **string** | The HTTP method we should use to call &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**FallbackToLongCode** | **bool** | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance. | [optional] 
**FallbackUrl** | **string** | The URL that we should call using &#x60;fallback_method&#x60; if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | [optional] 
**InboundMethod** | **string** | The HTTP method we should use to call &#x60;inbound_request_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | [optional] 
**InboundRequestUrl** | **string** | The URL we should call using &#x60;inbound_method&#x60; when a message is received by any phone number or short code in the Service. When this property is &#x60;null&#x60;, receiving inbound messages is disabled. | [optional] 
**MmsConverter** | **bool** | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance. | [optional] 
**ScanMessageContent** | **string** | Reserved. | [optional] 
**SmartEncoding** | **bool** | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance. | [optional] 
**StatusCallback** | **string** | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery. | [optional] 
**StickySender** | **bool** | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance. | [optional] 
**SynchronousValidation** | **bool** | Reserved. | [optional] 
**ValidityPeriod** | **int32** | How long, in seconds, messages sent from the Service are valid. Can be an integer from &#x60;1&#x60; to &#x60;14,400&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


