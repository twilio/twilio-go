# MessagingV1TollfreeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | Tollfree Verification Sid |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**CustomerProfileSid** | **string** | Customer's Profile Bundle BundleSid |[optional] 
**TrustProductSid** | **string** | Tollfree TrustProduct Bundle BundleSid |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**RegulatedItemSid** | **string** | The SID of the Regulated Item |[optional] 
**BusinessName** | **string** | The name of the business or organization using the Tollfree number |[optional] 
**BusinessStreetAddress** | **string** | The address of the business or organization using the Tollfree number |[optional] 
**BusinessStreetAddress2** | **string** | The address of the business or organization using the Tollfree number |[optional] 
**BusinessCity** | **string** | The city of the business or organization using the Tollfree number |[optional] 
**BusinessStateProvinceRegion** | **string** | The state/province/region of the business or organization using the Tollfree number |[optional] 
**BusinessPostalCode** | **string** | The postal code of the business or organization using the Tollfree number |[optional] 
**BusinessCountry** | **string** | The country of the business or organization using the Tollfree number |[optional] 
**BusinessWebsite** | **string** | The website of the business or organization using the Tollfree number |[optional] 
**BusinessContactFirstName** | **string** | The first name of the contact for the business or organization using the Tollfree number |[optional] 
**BusinessContactLastName** | **string** | The last name of the contact for the business or organization using the Tollfree number |[optional] 
**BusinessContactEmail** | **string** | The email address of the contact for the business or organization using the Tollfree number |[optional] 
**BusinessContactPhone** | **string** | The phone number of the contact for the business or organization using the Tollfree number |[optional] 
**NotificationEmail** | **string** | The email address to receive the notification about the verification result.  |[optional] 
**UseCaseCategories** | **[]string** | The category of the use case for the Tollfree Number. List as many are applicable. |[optional] 
**UseCaseSummary** | **string** | Further explaination on how messaging is used by the business or organization |[optional] 
**ProductionMessageSample** | **string** | An example of message content, i.e. a sample message |[optional] 
**OptInImageUrls** | **[]string** | Link to an image that shows the opt-in workflow. Multiple images allowed and must be a publicly hosted URL |[optional] 
**OptInType** | Pointer to [**string**](TollfreeVerificationEnumOptInType.md) |  |
**MessageVolume** | **string** | Estimate monthly volume of messages from the Tollfree Number |[optional] 
**AdditionalInformation** | **string** | Additional information to be provided for verification |[optional] 
**TollfreePhoneNumberSid** | **string** | The SID of the Phone Number associated with the Tollfree Verification |[optional] 
**Status** | Pointer to [**string**](TollfreeVerificationEnumStatus.md) |  |
**Url** | **string** | The absolute URL of the Tollfree Verification |[optional] 
**ResourceLinks** | Pointer to **interface{}** | The URLs of the documents associated with the Tollfree Verification resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


