# MessagingV1BrandRegistrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A2P BrandRegistration Sid |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CustomerProfileBundleSid** | Pointer to **string** | A2P Messaging Profile Bundle BundleSid |
**A2pProfileBundleSid** | Pointer to **string** | A2P Messaging Profile Bundle BundleSid |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**BrandType** | Pointer to **string** | Type of brand. One of: \"STANDARD\", \"STARTER\". |
**Status** | Pointer to [**string**](BrandRegistrationsEnumStatus.md) |  |
**TcrId** | Pointer to **string** | Campaign Registry (TCR) Brand ID |
**FailureReason** | Pointer to **string** | A reason why brand registration has failed |
**Url** | Pointer to **string** | The absolute URL of the Brand Registration |
**BrandScore** | Pointer to **int** | Brand score |
**BrandFeedback** | Pointer to [**[]string**](BrandRegistrationsEnumBrandFeedback.md) | Brand feedback |
**IdentityStatus** | Pointer to [**string**](BrandRegistrationsEnumIdentityStatus.md) |  |
**Russell3000** | Pointer to **bool** | Russell 3000 |
**GovernmentEntity** | Pointer to **bool** | Government Entity |
**TaxExemptStatus** | Pointer to **string** | Tax Exempt Status |
**SkipAutomaticSecVet** | Pointer to **bool** | Skip Automatic Secondary Vetting |
**Mock** | Pointer to **bool** | A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided. |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


