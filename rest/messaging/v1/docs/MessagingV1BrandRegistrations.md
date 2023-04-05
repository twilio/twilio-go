# MessagingV1BrandRegistrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string to identify Brand Registration. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Brand Registration resource. |
**CustomerProfileBundleSid** | Pointer to **string** | A2P Messaging Profile Bundle BundleSid. |
**A2pProfileBundleSid** | Pointer to **string** | A2P Messaging Profile Bundle BundleSid. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**BrandType** | Pointer to **string** | Type of brand. One of: \"STANDARD\", \"SOLE_PROPRIETOR\". SOLE_PROPRIETOR is for the low volume, SOLE_PROPRIETOR campaign use case. There can only be one SOLE_PROPRIETOR campaign created per SOLE_PROPRIETOR brand. STANDARD is for all other campaign use cases. Multiple campaign use cases can be created per STANDARD brand. |
**Status** | Pointer to [**string**](BrandRegistrationsEnumStatus.md) |  |
**TcrId** | Pointer to **string** | Campaign Registry (TCR) Brand ID. Assigned only after successful brand registration. |
**FailureReason** | Pointer to **string** | A reason why brand registration has failed. Only applicable when status is FAILED. |
**Url** | Pointer to **string** | The absolute URL of the Brand Registration resource. |
**BrandScore** | Pointer to **int** | The secondary vetting score if it was done. Otherwise, it will be the brand score if it's returned from TCR. It may be null if no score is available. |
**BrandFeedback** | Pointer to [**[]string**](BrandRegistrationsEnumBrandFeedback.md) | Feedback on how to improve brand score |
**IdentityStatus** | Pointer to [**string**](BrandRegistrationsEnumIdentityStatus.md) |  |
**Russell3000** | Pointer to **bool** | Publicly traded company identified in the Russell 3000 Index |
**GovernmentEntity** | Pointer to **bool** | Identified as a government entity |
**TaxExemptStatus** | Pointer to **string** | Nonprofit organization tax-exempt status per section 501 of the U.S. tax code. |
**SkipAutomaticSecVet** | Pointer to **bool** | A flag to disable automatic secondary vetting for brands which it would otherwise be done. |
**Mock** | Pointer to **bool** | A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided. |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


