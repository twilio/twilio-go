# MessagingV1RequestManagedCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSid** | Pointer to **string** | The unique string that we created to identify the Domain resource. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | Date that this Domain was last updated. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | Date that this Domain was registered to the Twilio platform to create a new Domain object. |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | Date that the private certificate associated with this domain expires. This is the expiration date of your existing cert. |
**DomainName** | Pointer to **string** | Full url path for this domain. |
**CertificateSid** | Pointer to **string** | The unique string that we created to identify this Certificate resource. |
**Url** | Pointer to **string** |  |
**Managed** | Pointer to **bool** | A boolean flag indicating if the certificate is managed by Twilio. |
**Requesting** | Pointer to **bool** | A boolean flag indicating if a managed certificate needs to be fulfilled by Twilio. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


