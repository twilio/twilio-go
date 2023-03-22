# MessagingV1DomainConfigMessagingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSid** | Pointer to **string** | The unique string that we created to identify the Domain resource. |
**ConfigSid** | Pointer to **string** | The unique string that we created to identify the Domain config (prefix ZK). |
**MessagingServiceSid** | Pointer to **string** | The unique string that identifies the messaging service |
**FallbackUrl** | Pointer to **string** | Any requests we receive to this domain that do not match an existing shortened message will be redirected to the fallback url. These will likely be either expired messages, random misdirected traffic, or intentional scraping. |
**CallbackUrl** | Pointer to **string** | URL to receive click events to your webhook whenever the recipients click on the shortened links. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | Date this Domain Config was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | Date that this Domain Config was last updated. |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


