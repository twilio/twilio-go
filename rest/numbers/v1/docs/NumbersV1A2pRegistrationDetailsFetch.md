# NumbersV1A2pRegistrationDetailsFetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | Account Sid that the phone number belongs to in Twilio. This is only returned for phone numbers that already exist in Twilio’s inventory and belong to your account or sub account. |
**PhoneNumberSid** | **string** | Phone Number SID for the requested phone number resource |
**PhoneNumber** | **string** |  |
**ExternalPhoneNumberStatus** | **string** |  |
**CampaignSid** | Pointer to **string** | Campaign Sid associated with the phone number |
**MessagingServiceSid** | Pointer to **string** | Messaging Service Sid that the number is associated with |
**ExternalCampaignId** | Pointer to **string** | The identifier for a campaign in the registrar. Typically, this is the TCR Campaign Id. |
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time when the A2P registration details were last updated |
**CampaignComplianceRegistrationSid** | Pointer to **string** | Sid associated with campaign compliance registration |
**BrandComplianceRegistrationSid** | Pointer to **string** | Sid associated with brand compliance registration |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


