# MessagingV1BrandRegistrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A2pProfileBundleSid** | Pointer to **string** | A2P Messaging Profile Bundle BundleSid |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**BrandFeedback** | Pointer to **[]string** | Brand feedback |
**BrandScore** | Pointer to **int** | Brand score |
**BrandType** | Pointer to **string** | Type of brand. One of: \"STANDARD\", \"STARTER\". |
**CustomerProfileBundleSid** | Pointer to **string** | A2P Messaging Profile Bundle BundleSid |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FailureReason** | Pointer to **string** | A reason why brand registration has failed |
**IdentityStatus** | Pointer to **string** | Identity Status |
**Links** | Pointer to **map[string]interface{}** |  |
**Mock** | Pointer to **bool** | A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided. |
**Russell3000** | Pointer to **bool** | Russell 3000 |
**Sid** | Pointer to **string** | A2P BrandRegistration Sid |
**SkipAutomaticSecVet** | Pointer to **bool** | Skip Automatic Secondary Vetting |
**Status** | Pointer to **string** | Brand Registration status. |
**TaxExemptStatus** | Pointer to **string** | Tax Exempt Status |
**TcrId** | Pointer to **string** | Campaign Registry (TCR) Brand ID |
**Url** | Pointer to **string** | The absolute URL of the Brand Registration |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


