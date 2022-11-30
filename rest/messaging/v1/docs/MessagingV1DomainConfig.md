# MessagingV1DomainConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSid** | Pointer to **string** | The unique string that we created to identify the Domain resource. |
**ConfigSid** | Pointer to **string** | The unique string that we created to identify the Domain config (prefix ZK). |
**MessagingServiceSids** | Pointer to **[]string** | A list of messagingServiceSids (with prefix MG). |
**FallbackUrl** | Pointer to **string** | We will redirect requests to urls we are unable to identify to this url. |
**CallbackUrl** | Pointer to **string** | URL to receive click events to your webhook whenever the recipients click on the shortened links. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | Date this Domain Config was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | Date that this Domain Config was last updated. |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


