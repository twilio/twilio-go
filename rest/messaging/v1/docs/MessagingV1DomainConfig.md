# MessagingV1DomainConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSid** | **string** | The unique string that we created to identify the Domain resource. |[optional] 
**ConfigSid** | **string** | The unique string that we created to identify the Domain config (prefix ZK). |[optional] 
**MessagingServiceSids** | **[]string** | A list of messagingServiceSids (with prefix MG). |[optional] 
**FallbackUrl** | **string** | We will redirect requests to urls we are unable to identify to this url. |[optional] 
**CallbackUrl** | **string** | URL to receive click events to your webhook whenever the recipients click on the shortened links. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | Date this Domain Config was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | Date that this Domain Config was last updated. |[optional] 
**Url** | **string** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


