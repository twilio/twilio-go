# TrusthubV1A2pBrandRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandType** | **string** | The type of brand being registered. |
**ThemeSetId** | **string** | Theme ID for styling the inquiry form. |[optional] 
**FriendlyName** | **string** | A human-readable name for the brand registration. |[optional] 
**NotificationEmail** | **string** | The email address to receive notifications about the registration. |
**BusinessName** | **string** | The legal name of the business. |[optional] 
**BusinessRegistrationAuthority** | **string** | The authority with which the business is registered. |[optional] 
**BusinessRegistrationNumber** | **string** | The business registration number (e.g., EIN, CBN, ACN). Maximum 21 characters. |[optional] 
**BusinessIndustry** | **string** | The industry of the business. |[optional] 
**BusinessWebsite** | **string** | The business website URL (must use https://). |[optional] 
**BusinessType** | **string** | The type of business entity. |[optional] 
**BusinessStockSymbol** | **string** | The stock symbol for publicly traded companies. |[optional] 
**BusinessStockExchange** | **string** | The stock exchange where the business is listed. |[optional] 
**BusinessTaxExemptStatus** | **string** | Nonprofit/political organization tax-exempt status per the U.S. tax code (e.g. 501c3, 501c1, 527). A value of 527 (political organization) requires a Campaign Verify token in brandExternalVettingToken. |[optional] 
**BrandExternalVettingToken** | **string** | Campaign Verify token for externally vetted brands. |[optional] 
**BusinessStreetAddress** | **string** | The street address of the business. |[optional] 
**BusinessStreetAddress2** | **string** | Additional street address information. |[optional] 
**BusinessCity** | **string** | The city of the business. |[optional] 
**BusinessStateProvinceRegion** | **string** | The state, province, or region of the business. |[optional] 
**BusinessPostalCode** | **string** | The postal code of the business. |[optional] 
**BusinessCountry** | **string** | The two-letter ISO country code of the business. |[optional] 
**BusinessContactFirstName** | **string** | The first name of the business contact. |[optional] 
**BusinessContactLastName** | **string** | The last name of the business contact. |[optional] 
**BusinessContactEmail** | **string** | The email address of the business contact. |[optional] 
**BusinessContactPhone** | **string** | The phone number of the business contact in E.164 format. |[optional] 
**AuthorizedContactVerificationEmail** | **string** | The email address for authorized contact verification. |[optional] 
**AuthorizedContactMobilePhoneNumberE164** | **string** | The mobile phone number for authorized contact in E.164 format. |[optional] 
**IsTest** | **bool** | Whether this is a test brand registration. |[optional] 
**SkipAutomaticSecVet** | **bool** | Whether to skip automatic secondary vetting. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


