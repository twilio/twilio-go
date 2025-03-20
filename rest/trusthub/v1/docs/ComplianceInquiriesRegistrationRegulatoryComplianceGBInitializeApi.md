# ComplianceInquiriesRegistrationRegulatoryComplianceGBInitializeApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComplianceRegistration**](ComplianceInquiriesRegistrationRegulatoryComplianceGBInitializeApi.md#CreateComplianceRegistration) | **Post** /v1/ComplianceInquiries/Registration/RegulatoryCompliance/GB/Initialize | Create a new Compliance Registration Inquiry for the authenticated account. This is necessary to start a new embedded session.
[**UpdateComplianceRegistration**](ComplianceInquiriesRegistrationRegulatoryComplianceGBInitializeApi.md#UpdateComplianceRegistration) | **Post** /v1/ComplianceInquiries/Registration/{RegistrationId}/RegulatoryCompliance/GB/Initialize | Resume a specific Regulatory Compliance Inquiry that has expired, or re-open a rejected Compliance Inquiry for editing.



## CreateComplianceRegistration

> TrusthubV1ComplianceRegistration CreateComplianceRegistration(ctx, optional)

Create a new Compliance Registration Inquiry for the authenticated account. This is necessary to start a new embedded session.

Create a new Compliance Registration Inquiry for the authenticated account. This is necessary to start a new embedded session.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateComplianceRegistrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndUserType** | [**string**](string.md) | 
**PhoneNumberType** | [**string**](string.md) | 
**BusinessIdentityType** | [**string**](string.md) | 
**BusinessRegistrationAuthority** | [**string**](string.md) | 
**BusinessLegalName** | **string** | he name of the business or organization using the Tollfree number.
**NotificationEmail** | **string** | he email address to receive the notification about the verification result.
**AcceptedNotificationReceipt** | **bool** | The email address to receive the notification about the verification result.
**BusinessRegistrationNumber** | **string** | Business registration number of the business
**BusinessWebsiteUrl** | **string** | The URL of the business website
**FriendlyName** | **string** | Friendly name for your business information
**AuthorizedRepresentative1FirstName** | **string** | First name of the authorized representative
**AuthorizedRepresentative1LastName** | **string** | Last name of the authorized representative
**AuthorizedRepresentative1Phone** | **string** | Phone number of the authorized representative
**AuthorizedRepresentative1Email** | **string** | Email address of the authorized representative
**AuthorizedRepresentative1DateOfBirth** | **string** | Birthdate of the authorized representative
**AddressStreet** | **string** | Street address of the business
**AddressStreetSecondary** | **string** | Street address of the business
**AddressCity** | **string** | City of the business
**AddressSubdivision** | **string** | State or province of the business
**AddressPostalCode** | **string** | Postal code of the business
**AddressCountryCode** | **string** | Country code of the business
**EmergencyAddressStreet** | **string** | Street address of the business
**EmergencyAddressStreetSecondary** | **string** | Street address of the business
**EmergencyAddressCity** | **string** | City of the business
**EmergencyAddressSubdivision** | **string** | State or province of the business
**EmergencyAddressPostalCode** | **string** | Postal code of the business
**EmergencyAddressCountryCode** | **string** | Country code of the business
**UseAddressAsEmergencyAddress** | **bool** | Use the business address as the emergency address
**FileName** | **string** | The name of the verification document to upload
**File** | **string** | The verification document to upload
**FirstName** | **string** | The first name of the Individual User.
**LastName** | **string** | The last name of the Individual User.
**DateOfBirth** | **string** | The date of birth of the Individual User.
**IndividualEmail** | **string** | The email address of the Individual User.
**IndividualPhone** | **string** | The phone number of the Individual User.
**IsIsvEmbed** | **bool** | Indicates if the inquiry is being started from an ISV embedded component.
**IsvRegisteringForSelfOrTenant** | **string** | Indicates if the isv registering for self or tenant.
**StatusCallbackUrl** | **string** | The url we call to inform you of bundle changes.
**ThemeSetId** | **string** | Theme id for styling the inquiry form.

### Return type

[**TrusthubV1ComplianceRegistration**](TrusthubV1ComplianceRegistration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComplianceRegistration

> TrusthubV1ComplianceRegistration UpdateComplianceRegistration(ctx, RegistrationIdoptional)

Resume a specific Regulatory Compliance Inquiry that has expired, or re-open a rejected Compliance Inquiry for editing.

Resume a specific Regulatory Compliance Inquiry that has expired, or re-open a rejected Compliance Inquiry for editing.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RegistrationId** | **string** | The unique RegistrationId matching the Regulatory Compliance Inquiry that should be resumed or resubmitted. This value will have been returned by the initial Regulatory Compliance Inquiry creation call.

### Other Parameters

Other parameters are passed through a pointer to a UpdateComplianceRegistrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**IsIsvEmbed** | **bool** | Indicates if the inquiry is being started from an ISV embedded component.
**ThemeSetId** | **string** | Theme id for styling the inquiry form.

### Return type

[**TrusthubV1ComplianceRegistration**](TrusthubV1ComplianceRegistration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

