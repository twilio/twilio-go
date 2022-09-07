# MessagingV1TollfreeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | Tollfree Verification Sid |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CustomerProfileSid** | Pointer to **string** | Customer's Profile Bundle BundleSid |
**TrustProductSid** | Pointer to **string** | Tollfree TrustProduct Bundle BundleSid |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**RegulatedItemSid** | Pointer to **string** | The SID of the Regulated Item |
**BusinessName** | Pointer to **string** | The name of the business or organization using the Tollfree number |
**BusinessStreetAddress** | Pointer to **string** | The address of the business or organization using the Tollfree number |
**BusinessStreetAddress2** | Pointer to **string** | The address of the business or organization using the Tollfree number |
**BusinessCity** | Pointer to **string** | The city of the business or organization using the Tollfree number |
**BusinessStateProvinceRegion** | Pointer to **string** | The state/province/region of the business or organization using the Tollfree number |
**BusinessPostalCode** | Pointer to **string** | The postal code of the business or organization using the Tollfree number |
**BusinessCountry** | Pointer to **string** | The country of the business or organization using the Tollfree number |
**BusinessWebsite** | Pointer to **string** | The website of the business or organization using the Tollfree number |
**BusinessContactFirstName** | Pointer to **string** | The first name of the contact for the business or organization using the Tollfree number |
**BusinessContactLastName** | Pointer to **string** | The last name of the contact for the business or organization using the Tollfree number |
**BusinessContactEmail** | Pointer to **string** | The email address of the contact for the business or organization using the Tollfree number |
**BusinessContactPhone** | Pointer to **string** | The phone number of the contact for the business or organization using the Tollfree number |
**NotificationEmail** | Pointer to **string** | The email address to receive the notification about the verification result.  |
**UseCaseCategories** | Pointer to **[]string** | The category of the use case for the Tollfree Number. List as many are applicable. |
**UseCaseSummary** | Pointer to **string** | Further explaination on how messaging is used by the business or organization |
**ProductionMessageSample** | Pointer to **string** | An example of message content, i.e. a sample message |
**OptInImageUrls** | Pointer to **[]string** | Link to an image that shows the opt-in workflow. Multiple images allowed and must be a publicly hosted URL |
**OptInType** | Pointer to [**string**](TollfreeVerificationEnumOptInType.md) |  |
**MessageVolume** | Pointer to **string** | Estimate monthly volume of messages from the Tollfree Number |
**AdditionalInformation** | Pointer to **string** | Additional information to be provided for verification |
**TollfreePhoneNumberSid** | Pointer to **string** | The SID of the Phone Number associated with the Tollfree Verification |
**Status** | Pointer to [**string**](TollfreeVerificationEnumStatus.md) |  |
**Url** | Pointer to **string** | The absolute URL of the Tollfree Verification |
**ResourceLinks** | Pointer to **interface{}** | The URLs of the documents associated with the Tollfree Verification resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


