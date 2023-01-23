# MessagingV1DomainCertV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSid** | Pointer to **string** | The unique string that we created to identify the Domain resource. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | Date that this Domain was last updated. |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | Date that the private certificate associated with this domain expires. You will need to update the certificate before that date to ensure your shortened links will continue to work. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | Date that this Domain was registered to the Twilio platform to create a new Domain object. |
**DomainName** | Pointer to **string** | Full url path for this domain. |
**CertificateSid** | Pointer to **string** | The unique string that we created to identify this Certificate resource. |
**Url** | Pointer to **string** |  |
**Validated** | Pointer to **bool** | Boolean value indicating whether certificate has been validated |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


