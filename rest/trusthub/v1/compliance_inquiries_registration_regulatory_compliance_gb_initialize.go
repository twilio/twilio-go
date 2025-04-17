/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateComplianceRegistration'
type CreateComplianceRegistrationParams struct {
	//
	EndUserType *string `json:"EndUserType,omitempty"`
	//
	PhoneNumberType *string `json:"PhoneNumberType,omitempty"`
	//
	BusinessIdentityType *string `json:"BusinessIdentityType,omitempty"`
	//
	BusinessRegistrationAuthority *string `json:"BusinessRegistrationAuthority,omitempty"`
	// he name of the business or organization using the Tollfree number.
	BusinessLegalName *string `json:"BusinessLegalName,omitempty"`
	// he email address to receive the notification about the verification result.
	NotificationEmail *string `json:"NotificationEmail,omitempty"`
	// The email address to receive the notification about the verification result.
	AcceptedNotificationReceipt *bool `json:"AcceptedNotificationReceipt,omitempty"`
	// Business registration number of the business
	BusinessRegistrationNumber *string `json:"BusinessRegistrationNumber,omitempty"`
	// The URL of the business website
	BusinessWebsiteUrl *string `json:"BusinessWebsiteUrl,omitempty"`
	// Friendly name for your business information
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// First name of the authorized representative
	AuthorizedRepresentative1FirstName *string `json:"AuthorizedRepresentative1FirstName,omitempty"`
	// Last name of the authorized representative
	AuthorizedRepresentative1LastName *string `json:"AuthorizedRepresentative1LastName,omitempty"`
	// Phone number of the authorized representative
	AuthorizedRepresentative1Phone *string `json:"AuthorizedRepresentative1Phone,omitempty"`
	// Email address of the authorized representative
	AuthorizedRepresentative1Email *string `json:"AuthorizedRepresentative1Email,omitempty"`
	// Birthdate of the authorized representative
	AuthorizedRepresentative1DateOfBirth *string `json:"AuthorizedRepresentative1DateOfBirth,omitempty"`
	// Street address of the business
	AddressStreet *string `json:"AddressStreet,omitempty"`
	// Street address of the business
	AddressStreetSecondary *string `json:"AddressStreetSecondary,omitempty"`
	// City of the business
	AddressCity *string `json:"AddressCity,omitempty"`
	// State or province of the business
	AddressSubdivision *string `json:"AddressSubdivision,omitempty"`
	// Postal code of the business
	AddressPostalCode *string `json:"AddressPostalCode,omitempty"`
	// Country code of the business
	AddressCountryCode *string `json:"AddressCountryCode,omitempty"`
	// Street address of the business
	EmergencyAddressStreet *string `json:"EmergencyAddressStreet,omitempty"`
	// Street address of the business
	EmergencyAddressStreetSecondary *string `json:"EmergencyAddressStreetSecondary,omitempty"`
	// City of the business
	EmergencyAddressCity *string `json:"EmergencyAddressCity,omitempty"`
	// State or province of the business
	EmergencyAddressSubdivision *string `json:"EmergencyAddressSubdivision,omitempty"`
	// Postal code of the business
	EmergencyAddressPostalCode *string `json:"EmergencyAddressPostalCode,omitempty"`
	// Country code of the business
	EmergencyAddressCountryCode *string `json:"EmergencyAddressCountryCode,omitempty"`
	// Use the business address as the emergency address
	UseAddressAsEmergencyAddress *bool `json:"UseAddressAsEmergencyAddress,omitempty"`
	// The name of the verification document to upload
	FileName *string `json:"FileName,omitempty"`
	// The verification document to upload
	File *string `json:"File,omitempty"`
	// The first name of the Individual User.
	FirstName *string `json:"FirstName,omitempty"`
	// The last name of the Individual User.
	LastName *string `json:"LastName,omitempty"`
	// The date of birth of the Individual User.
	DateOfBirth *string `json:"DateOfBirth,omitempty"`
	// The email address of the Individual User.
	IndividualEmail *string `json:"IndividualEmail,omitempty"`
	// The phone number of the Individual User.
	IndividualPhone *string `json:"IndividualPhone,omitempty"`
	// Indicates if the inquiry is being started from an ISV embedded component.
	IsIsvEmbed *bool `json:"IsIsvEmbed,omitempty"`
	// Indicates if the isv registering for self or tenant.
	IsvRegisteringForSelfOrTenant *string `json:"IsvRegisteringForSelfOrTenant,omitempty"`
	// The url we call to inform you of bundle changes.
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty"`
	// Theme id for styling the inquiry form.
	ThemeSetId *string `json:"ThemeSetId,omitempty"`
}

func (params *CreateComplianceRegistrationParams) SetEndUserType(EndUserType string) *CreateComplianceRegistrationParams {
	params.EndUserType = &EndUserType
	return params
}
func (params *CreateComplianceRegistrationParams) SetPhoneNumberType(PhoneNumberType string) *CreateComplianceRegistrationParams {
	params.PhoneNumberType = &PhoneNumberType
	return params
}
func (params *CreateComplianceRegistrationParams) SetBusinessIdentityType(BusinessIdentityType string) *CreateComplianceRegistrationParams {
	params.BusinessIdentityType = &BusinessIdentityType
	return params
}
func (params *CreateComplianceRegistrationParams) SetBusinessRegistrationAuthority(BusinessRegistrationAuthority string) *CreateComplianceRegistrationParams {
	params.BusinessRegistrationAuthority = &BusinessRegistrationAuthority
	return params
}
func (params *CreateComplianceRegistrationParams) SetBusinessLegalName(BusinessLegalName string) *CreateComplianceRegistrationParams {
	params.BusinessLegalName = &BusinessLegalName
	return params
}
func (params *CreateComplianceRegistrationParams) SetNotificationEmail(NotificationEmail string) *CreateComplianceRegistrationParams {
	params.NotificationEmail = &NotificationEmail
	return params
}
func (params *CreateComplianceRegistrationParams) SetAcceptedNotificationReceipt(AcceptedNotificationReceipt bool) *CreateComplianceRegistrationParams {
	params.AcceptedNotificationReceipt = &AcceptedNotificationReceipt
	return params
}
func (params *CreateComplianceRegistrationParams) SetBusinessRegistrationNumber(BusinessRegistrationNumber string) *CreateComplianceRegistrationParams {
	params.BusinessRegistrationNumber = &BusinessRegistrationNumber
	return params
}
func (params *CreateComplianceRegistrationParams) SetBusinessWebsiteUrl(BusinessWebsiteUrl string) *CreateComplianceRegistrationParams {
	params.BusinessWebsiteUrl = &BusinessWebsiteUrl
	return params
}
func (params *CreateComplianceRegistrationParams) SetFriendlyName(FriendlyName string) *CreateComplianceRegistrationParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateComplianceRegistrationParams) SetAuthorizedRepresentative1FirstName(AuthorizedRepresentative1FirstName string) *CreateComplianceRegistrationParams {
	params.AuthorizedRepresentative1FirstName = &AuthorizedRepresentative1FirstName
	return params
}
func (params *CreateComplianceRegistrationParams) SetAuthorizedRepresentative1LastName(AuthorizedRepresentative1LastName string) *CreateComplianceRegistrationParams {
	params.AuthorizedRepresentative1LastName = &AuthorizedRepresentative1LastName
	return params
}
func (params *CreateComplianceRegistrationParams) SetAuthorizedRepresentative1Phone(AuthorizedRepresentative1Phone string) *CreateComplianceRegistrationParams {
	params.AuthorizedRepresentative1Phone = &AuthorizedRepresentative1Phone
	return params
}
func (params *CreateComplianceRegistrationParams) SetAuthorizedRepresentative1Email(AuthorizedRepresentative1Email string) *CreateComplianceRegistrationParams {
	params.AuthorizedRepresentative1Email = &AuthorizedRepresentative1Email
	return params
}
func (params *CreateComplianceRegistrationParams) SetAuthorizedRepresentative1DateOfBirth(AuthorizedRepresentative1DateOfBirth string) *CreateComplianceRegistrationParams {
	params.AuthorizedRepresentative1DateOfBirth = &AuthorizedRepresentative1DateOfBirth
	return params
}
func (params *CreateComplianceRegistrationParams) SetAddressStreet(AddressStreet string) *CreateComplianceRegistrationParams {
	params.AddressStreet = &AddressStreet
	return params
}
func (params *CreateComplianceRegistrationParams) SetAddressStreetSecondary(AddressStreetSecondary string) *CreateComplianceRegistrationParams {
	params.AddressStreetSecondary = &AddressStreetSecondary
	return params
}
func (params *CreateComplianceRegistrationParams) SetAddressCity(AddressCity string) *CreateComplianceRegistrationParams {
	params.AddressCity = &AddressCity
	return params
}
func (params *CreateComplianceRegistrationParams) SetAddressSubdivision(AddressSubdivision string) *CreateComplianceRegistrationParams {
	params.AddressSubdivision = &AddressSubdivision
	return params
}
func (params *CreateComplianceRegistrationParams) SetAddressPostalCode(AddressPostalCode string) *CreateComplianceRegistrationParams {
	params.AddressPostalCode = &AddressPostalCode
	return params
}
func (params *CreateComplianceRegistrationParams) SetAddressCountryCode(AddressCountryCode string) *CreateComplianceRegistrationParams {
	params.AddressCountryCode = &AddressCountryCode
	return params
}
func (params *CreateComplianceRegistrationParams) SetEmergencyAddressStreet(EmergencyAddressStreet string) *CreateComplianceRegistrationParams {
	params.EmergencyAddressStreet = &EmergencyAddressStreet
	return params
}
func (params *CreateComplianceRegistrationParams) SetEmergencyAddressStreetSecondary(EmergencyAddressStreetSecondary string) *CreateComplianceRegistrationParams {
	params.EmergencyAddressStreetSecondary = &EmergencyAddressStreetSecondary
	return params
}
func (params *CreateComplianceRegistrationParams) SetEmergencyAddressCity(EmergencyAddressCity string) *CreateComplianceRegistrationParams {
	params.EmergencyAddressCity = &EmergencyAddressCity
	return params
}
func (params *CreateComplianceRegistrationParams) SetEmergencyAddressSubdivision(EmergencyAddressSubdivision string) *CreateComplianceRegistrationParams {
	params.EmergencyAddressSubdivision = &EmergencyAddressSubdivision
	return params
}
func (params *CreateComplianceRegistrationParams) SetEmergencyAddressPostalCode(EmergencyAddressPostalCode string) *CreateComplianceRegistrationParams {
	params.EmergencyAddressPostalCode = &EmergencyAddressPostalCode
	return params
}
func (params *CreateComplianceRegistrationParams) SetEmergencyAddressCountryCode(EmergencyAddressCountryCode string) *CreateComplianceRegistrationParams {
	params.EmergencyAddressCountryCode = &EmergencyAddressCountryCode
	return params
}
func (params *CreateComplianceRegistrationParams) SetUseAddressAsEmergencyAddress(UseAddressAsEmergencyAddress bool) *CreateComplianceRegistrationParams {
	params.UseAddressAsEmergencyAddress = &UseAddressAsEmergencyAddress
	return params
}
func (params *CreateComplianceRegistrationParams) SetFileName(FileName string) *CreateComplianceRegistrationParams {
	params.FileName = &FileName
	return params
}
func (params *CreateComplianceRegistrationParams) SetFile(File string) *CreateComplianceRegistrationParams {
	params.File = &File
	return params
}
func (params *CreateComplianceRegistrationParams) SetFirstName(FirstName string) *CreateComplianceRegistrationParams {
	params.FirstName = &FirstName
	return params
}
func (params *CreateComplianceRegistrationParams) SetLastName(LastName string) *CreateComplianceRegistrationParams {
	params.LastName = &LastName
	return params
}
func (params *CreateComplianceRegistrationParams) SetDateOfBirth(DateOfBirth string) *CreateComplianceRegistrationParams {
	params.DateOfBirth = &DateOfBirth
	return params
}
func (params *CreateComplianceRegistrationParams) SetIndividualEmail(IndividualEmail string) *CreateComplianceRegistrationParams {
	params.IndividualEmail = &IndividualEmail
	return params
}
func (params *CreateComplianceRegistrationParams) SetIndividualPhone(IndividualPhone string) *CreateComplianceRegistrationParams {
	params.IndividualPhone = &IndividualPhone
	return params
}
func (params *CreateComplianceRegistrationParams) SetIsIsvEmbed(IsIsvEmbed bool) *CreateComplianceRegistrationParams {
	params.IsIsvEmbed = &IsIsvEmbed
	return params
}
func (params *CreateComplianceRegistrationParams) SetIsvRegisteringForSelfOrTenant(IsvRegisteringForSelfOrTenant string) *CreateComplianceRegistrationParams {
	params.IsvRegisteringForSelfOrTenant = &IsvRegisteringForSelfOrTenant
	return params
}
func (params *CreateComplianceRegistrationParams) SetStatusCallbackUrl(StatusCallbackUrl string) *CreateComplianceRegistrationParams {
	params.StatusCallbackUrl = &StatusCallbackUrl
	return params
}
func (params *CreateComplianceRegistrationParams) SetThemeSetId(ThemeSetId string) *CreateComplianceRegistrationParams {
	params.ThemeSetId = &ThemeSetId
	return params
}

// Create a new Compliance Registration Inquiry for the authenticated account. This is necessary to start a new embedded session.
func (c *ApiService) CreateComplianceRegistration(params *CreateComplianceRegistrationParams) (*TrusthubV1ComplianceRegistration, error) {
	path := "/v1/ComplianceInquiries/Registration/RegulatoryCompliance/GB/Initialize"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.EndUserType != nil {
		data.Set("EndUserType", fmt.Sprint(*params.EndUserType))
	}
	if params != nil && params.PhoneNumberType != nil {
		data.Set("PhoneNumberType", fmt.Sprint(*params.PhoneNumberType))
	}
	if params != nil && params.BusinessIdentityType != nil {
		data.Set("BusinessIdentityType", fmt.Sprint(*params.BusinessIdentityType))
	}
	if params != nil && params.BusinessRegistrationAuthority != nil {
		data.Set("BusinessRegistrationAuthority", fmt.Sprint(*params.BusinessRegistrationAuthority))
	}
	if params != nil && params.BusinessLegalName != nil {
		data.Set("BusinessLegalName", *params.BusinessLegalName)
	}
	if params != nil && params.NotificationEmail != nil {
		data.Set("NotificationEmail", *params.NotificationEmail)
	}
	if params != nil && params.AcceptedNotificationReceipt != nil {
		data.Set("AcceptedNotificationReceipt", fmt.Sprint(*params.AcceptedNotificationReceipt))
	}
	if params != nil && params.BusinessRegistrationNumber != nil {
		data.Set("BusinessRegistrationNumber", *params.BusinessRegistrationNumber)
	}
	if params != nil && params.BusinessWebsiteUrl != nil {
		data.Set("BusinessWebsiteUrl", *params.BusinessWebsiteUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.AuthorizedRepresentative1FirstName != nil {
		data.Set("AuthorizedRepresentative1FirstName", *params.AuthorizedRepresentative1FirstName)
	}
	if params != nil && params.AuthorizedRepresentative1LastName != nil {
		data.Set("AuthorizedRepresentative1LastName", *params.AuthorizedRepresentative1LastName)
	}
	if params != nil && params.AuthorizedRepresentative1Phone != nil {
		data.Set("AuthorizedRepresentative1Phone", *params.AuthorizedRepresentative1Phone)
	}
	if params != nil && params.AuthorizedRepresentative1Email != nil {
		data.Set("AuthorizedRepresentative1Email", *params.AuthorizedRepresentative1Email)
	}
	if params != nil && params.AuthorizedRepresentative1DateOfBirth != nil {
		data.Set("AuthorizedRepresentative1DateOfBirth", *params.AuthorizedRepresentative1DateOfBirth)
	}
	if params != nil && params.AddressStreet != nil {
		data.Set("AddressStreet", *params.AddressStreet)
	}
	if params != nil && params.AddressStreetSecondary != nil {
		data.Set("AddressStreetSecondary", *params.AddressStreetSecondary)
	}
	if params != nil && params.AddressCity != nil {
		data.Set("AddressCity", *params.AddressCity)
	}
	if params != nil && params.AddressSubdivision != nil {
		data.Set("AddressSubdivision", *params.AddressSubdivision)
	}
	if params != nil && params.AddressPostalCode != nil {
		data.Set("AddressPostalCode", *params.AddressPostalCode)
	}
	if params != nil && params.AddressCountryCode != nil {
		data.Set("AddressCountryCode", *params.AddressCountryCode)
	}
	if params != nil && params.EmergencyAddressStreet != nil {
		data.Set("EmergencyAddressStreet", *params.EmergencyAddressStreet)
	}
	if params != nil && params.EmergencyAddressStreetSecondary != nil {
		data.Set("EmergencyAddressStreetSecondary", *params.EmergencyAddressStreetSecondary)
	}
	if params != nil && params.EmergencyAddressCity != nil {
		data.Set("EmergencyAddressCity", *params.EmergencyAddressCity)
	}
	if params != nil && params.EmergencyAddressSubdivision != nil {
		data.Set("EmergencyAddressSubdivision", *params.EmergencyAddressSubdivision)
	}
	if params != nil && params.EmergencyAddressPostalCode != nil {
		data.Set("EmergencyAddressPostalCode", *params.EmergencyAddressPostalCode)
	}
	if params != nil && params.EmergencyAddressCountryCode != nil {
		data.Set("EmergencyAddressCountryCode", *params.EmergencyAddressCountryCode)
	}
	if params != nil && params.UseAddressAsEmergencyAddress != nil {
		data.Set("UseAddressAsEmergencyAddress", fmt.Sprint(*params.UseAddressAsEmergencyAddress))
	}
	if params != nil && params.FileName != nil {
		data.Set("FileName", *params.FileName)
	}
	if params != nil && params.File != nil {
		data.Set("File", *params.File)
	}
	if params != nil && params.FirstName != nil {
		data.Set("FirstName", *params.FirstName)
	}
	if params != nil && params.LastName != nil {
		data.Set("LastName", *params.LastName)
	}
	if params != nil && params.DateOfBirth != nil {
		data.Set("DateOfBirth", *params.DateOfBirth)
	}
	if params != nil && params.IndividualEmail != nil {
		data.Set("IndividualEmail", *params.IndividualEmail)
	}
	if params != nil && params.IndividualPhone != nil {
		data.Set("IndividualPhone", *params.IndividualPhone)
	}
	if params != nil && params.IsIsvEmbed != nil {
		data.Set("IsIsvEmbed", fmt.Sprint(*params.IsIsvEmbed))
	}
	if params != nil && params.IsvRegisteringForSelfOrTenant != nil {
		data.Set("IsvRegisteringForSelfOrTenant", *params.IsvRegisteringForSelfOrTenant)
	}
	if params != nil && params.StatusCallbackUrl != nil {
		data.Set("StatusCallbackUrl", *params.StatusCallbackUrl)
	}
	if params != nil && params.ThemeSetId != nil {
		data.Set("ThemeSetId", *params.ThemeSetId)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1ComplianceRegistration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateComplianceRegistration'
type UpdateComplianceRegistrationParams struct {
	// Indicates if the inquiry is being started from an ISV embedded component.
	IsIsvEmbed *bool `json:"IsIsvEmbed,omitempty"`
	// Theme id for styling the inquiry form.
	ThemeSetId *string `json:"ThemeSetId,omitempty"`
}

func (params *UpdateComplianceRegistrationParams) SetIsIsvEmbed(IsIsvEmbed bool) *UpdateComplianceRegistrationParams {
	params.IsIsvEmbed = &IsIsvEmbed
	return params
}
func (params *UpdateComplianceRegistrationParams) SetThemeSetId(ThemeSetId string) *UpdateComplianceRegistrationParams {
	params.ThemeSetId = &ThemeSetId
	return params
}

// Resume a specific Regulatory Compliance Inquiry that has expired, or re-open a rejected Compliance Inquiry for editing.
func (c *ApiService) UpdateComplianceRegistration(RegistrationId string, params *UpdateComplianceRegistrationParams) (*TrusthubV1ComplianceRegistration, error) {
	path := "/v1/ComplianceInquiries/Registration/{RegistrationId}/RegulatoryCompliance/GB/Initialize"
	path = strings.Replace(path, "{"+"RegistrationId"+"}", RegistrationId, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.IsIsvEmbed != nil {
		data.Set("IsIsvEmbed", fmt.Sprint(*params.IsIsvEmbed))
	}
	if params != nil && params.ThemeSetId != nil {
		data.Set("ThemeSetId", *params.ThemeSetId)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1ComplianceRegistration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
