# MessagingV1BrandRegistrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A2P BrandRegistration Sid |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**CustomerProfileBundleSid** | **string** | A2P Messaging Profile Bundle BundleSid |[optional] 
**A2pProfileBundleSid** | **string** | A2P Messaging Profile Bundle BundleSid |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**BrandType** | **string** | Type of brand. One of: \"STANDARD\", \"STARTER\". |[optional] 
**Status** | Pointer to [**string**](BrandRegistrationsEnumStatus.md) |  |
**TcrId** | **string** | Campaign Registry (TCR) Brand ID |[optional] 
**FailureReason** | **string** | A reason why brand registration has failed |[optional] 
**Url** | **string** | The absolute URL of the Brand Registration |[optional] 
**BrandScore** | Pointer to **int** | Brand score |
**BrandFeedback** | Pointer to [**[]string**](BrandRegistrationsEnumBrandFeedback.md) | Brand feedback |
**IdentityStatus** | Pointer to [**string**](BrandRegistrationsEnumIdentityStatus.md) |  |
**Russell3000** | **bool** | Russell 3000 |[optional] 
**GovernmentEntity** | **bool** | Government Entity |[optional] 
**TaxExemptStatus** | Pointer to **string** | Tax Exempt Status |
**SkipAutomaticSecVet** | **bool** | Skip Automatic Secondary Vetting |[optional] 
**Mock** | **bool** | A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided. |[optional] 
**Links** | **map[string]interface{}** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


