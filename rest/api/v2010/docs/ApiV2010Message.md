# ApiV2010Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The message text. Can be up to 1,600 characters long. |
**NumSegments** | Pointer to **string** | The number of segments that make up the complete message. A message body that is too large to be sent in a single SMS message is segmented and charged as multiple messages. Inbound messages over 160 characters are reassembled when the message is received. Note: When using a Messaging Service to send messages, `num_segments` will always be 0 in Twilio's response to your API request. |
**Direction** | Pointer to [**string**](MessageEnumDirection.md) |  |
**From** | Pointer to **string** | The phone number (in [E.164](https://en.wikipedia.org/wiki/E.164) format), [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or [Wireless SIM](https://www.twilio.com/docs/wireless/tutorials/communications-guides/how-to-send-and-receive-text-messages) that initiated the message. For incoming messages, this will be the number of the sending phone. For outgoing messages, this value will be one of your Twilio phone numbers or the alphanumeric sender ID used. |
**To** | Pointer to **string** | The phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format that received the message. For incoming messages, this will be one of your Twilio phone numbers. For outgoing messages, this will be the sending phone. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Price** | Pointer to **string** | The amount billed for the message, in the currency specified by `price_unit`.  Note that your account is charged for each segment we send to the handset. Populated after the message has been sent. May not be immediately available. |
**ErrorMessage** | Pointer to **string** | The description of the `error_code` if your message `status` is `failed` or `undelivered`. If the message was successful, this value is null. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that sent the message that created the resource. |
**NumMedia** | Pointer to **string** | The number of media files associated with the message. A message can send up to 10 media files. |
**Status** | Pointer to [**string**](MessageEnumStatus.md) |  |
**MessagingServiceSid** | Pointer to **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) used with the message. The value is null if a Messaging Service was not used. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the Message resource. |
**DateSent** | Pointer to **string** | The date and time in GMT that the resource was sent specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. For outgoing messages, this is when we sent the message. For incoming messages, this is when we made the HTTP request to your application.  |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**ErrorCode** | Pointer to **int** | The error code returned if your message `status` is `failed` or `undelivered`. The error_code provides more information about the failure. If the message was successful, this value is null. |
**PriceUnit** | Pointer to **string** | The currency in which `price` is measured, in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`). |
**ApiVersion** | Pointer to **string** | The API version used to process the message. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their URIs relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


