# CountryRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCountry** | **string** | Iso country code as per ISO 3166-1 alpha-2 standard |
**RegistrationRequired** | **bool** | Whether Sender ID needs to be pre-registered for the country |
**SlaInDays** | **int** | Twilio SLA for Sender Id Registration process in business days. For countries requiring dynamic registration, it will be set to 0. |
**PromotionalSupported** | **bool** | Whether promotional usage for Sender ID is supported |[optional] 
**PromotionalSenderIdPrefix** | Pointer to **string** | Mandatory prefix string for Sender ID when used for promotional purpose in the country |
**PromotionalSenderIdSuffix** | Pointer to **string** | Mandatory suffix string for Sender ID when used for promotional purpose in the country |
**PricingScheme** | Pointer to **string** | Represents pricing requirements for country with free-flowing string format |
**DocumentationUrl** | Pointer to **string** | Represents public Twilio support URL which has information regarding the instructions and documents required for registration |
**DocumentationTemplateUrl** | Pointer to **string** | Represents the Twilio public URL for documentation template required to be filled for the Sender ID registration |
**DocumentTypeMachineNames** | **[]string** | List of document type machine names |[optional] 
**DomesticDocumentTypeMachineNames** | **[]string** | List of document type machine names for Domestic traffic reach |[optional] 
**InternationalDocumentTypeMachineNames** | **[]string** | List of document type machine names for International traffic reach |[optional] 
**SenderIdRegistrationRules** | Pointer to **string** | Sender ID string rules for the country |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


