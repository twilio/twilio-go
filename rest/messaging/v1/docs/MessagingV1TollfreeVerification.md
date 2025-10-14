# MessagingV1TollfreeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string to identify Tollfree Verification. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Tollfree Verification resource. |
**CustomerProfileSid** | Pointer to **string** | Customer's Profile Bundle BundleSid. |
**TrustProductSid** | Pointer to **string** | Tollfree TrustProduct Bundle BundleSid. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**RegulatedItemSid** | Pointer to **string** | The SID of the Regulated Item. |
**BusinessName** | Pointer to **string** | The name of the business or organization using the Tollfree number. |
**BusinessStreetAddress** | Pointer to **string** | The address of the business or organization using the Tollfree number. |
**BusinessStreetAddress2** | Pointer to **string** | The address of the business or organization using the Tollfree number. |
**BusinessCity** | Pointer to **string** | The city of the business or organization using the Tollfree number. |
**BusinessStateProvinceRegion** | Pointer to **string** | The state/province/region of the business or organization using the Tollfree number. |
**BusinessPostalCode** | Pointer to **string** | The postal code of the business or organization using the Tollfree number. |
**BusinessCountry** | Pointer to **string** | The country of the business or organization using the Tollfree number. |
**BusinessWebsite** | Pointer to **string** | The website of the business or organization using the Tollfree number. |
**BusinessContactFirstName** | Pointer to **string** | The first name of the contact for the business or organization using the Tollfree number. |
**BusinessContactLastName** | Pointer to **string** | The last name of the contact for the business or organization using the Tollfree number. |
**BusinessContactEmail** | Pointer to **string** | The email address of the contact for the business or organization using the Tollfree number. |
**BusinessContactPhone** | Pointer to **string** | The E.164 formatted phone number of the contact for the business or organization using the Tollfree number. |
**NotificationEmail** | Pointer to **string** | The email address to receive the notification about the verification result. . |
**UseCaseCategories** | Pointer to **[]string** | The category of the use case for the Tollfree Number. List as many are applicable.. |
**UseCaseSummary** | Pointer to **string** | Use this to further explain how messaging is used by the business or organization. |
**ProductionMessageSample** | Pointer to **string** | An example of message content, i.e. a sample message. |
**OptInImageUrls** | Pointer to **[]string** | Link to an image that shows the opt-in workflow. Multiple images allowed and must be a publicly hosted URL. |
**OptInType** | Pointer to [**string**](TollfreeVerificationEnumOptInType.md) |  |
**MessageVolume** | Pointer to **string** | Estimate monthly volume of messages from the Tollfree Number. |
**AdditionalInformation** | Pointer to **string** | Additional information to be provided for verification. |
**TollfreePhoneNumberSid** | Pointer to **string** | The SID of the Phone Number associated with the Tollfree Verification. |
**TollfreePhoneNumber** | Pointer to **string** | The E.164 formatted toll-free phone number associated with the verification. |
**Status** | Pointer to [**string**](TollfreeVerificationEnumStatus.md) |  |
**Url** | Pointer to **string** | The absolute URL of the Tollfree Verification resource. |
**RejectionReason** | Pointer to **string** | The rejection reason given when a Tollfree Verification has been rejected. |
**ErrorCode** | Pointer to **int** | The error code given when a Tollfree Verification has been rejected. |
**EditExpiration** | Pointer to [**time.Time**](time.Time.md) | The date and time when the ability to edit a rejected verification expires. |
**EditAllowed** | Pointer to **bool** | If a rejected verification is allowed to be edited/resubmitted. Some rejection reasons allow editing and some do not. |
**BusinessRegistrationNumber** | Pointer to **string** | A legally recognized business registration number |
**BusinessRegistrationAuthority** | Pointer to **string** | The organizational authority for business registrations |
**BusinessRegistrationCountry** | Pointer to **string** | Country business is registered in |
**BusinessType** | Pointer to **string** | The type of business, valid values are PRIVATE_PROFIT, PUBLIC_PROFIT, NON_PROFIT, SOLE_PROPRIETOR, GOVERNMENT |
**BusinessRegistrationPhoneNumber** | Pointer to **string** | The E.164 formatted number associated with the business. |
**DoingBusinessAs** | Pointer to **string** | Trade name, sub entity, or downstream business name of business being submitted for verification |
**OptInConfirmationMessage** | Pointer to **string** | The confirmation message sent to users when they opt in to receive messages. |
**HelpMessageSample** | Pointer to **string** | A sample help message provided to users. |
**PrivacyPolicyUrl** | Pointer to **string** | The URL to the privacy policy for the business or organization. |
**TermsAndConditionsUrl** | Pointer to **string** | The URL of the terms and conditions for the business or organization. |
**AgeGatedContent** | Pointer to **bool** | Indicates if the content is age gated. |
**OptInKeywords** | Pointer to **[]string** | List of keywords that users can send to opt in or out of messages. |
**RejectionReasons** | Pointer to **[]interface{}** | A list of rejection reasons and codes describing why a Tollfree Verification has been rejected. |
**ResourceLinks** | Pointer to **interface{}** | The URLs of the documents associated with the Tollfree Verification resource. |
**ExternalReferenceId** | Pointer to **string** | An optional external reference ID supplied by customer and echoed back on status retrieval. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


