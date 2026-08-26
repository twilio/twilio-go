# TrusthubV1A2pCampaignRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A2pBrandRegistrationSid** | **string** | The legacy Twilio A2P Brand SID (BN-prefixed) for the brand this campaign is registered under. This is a different identifier from the `id` returned by the Initialize Brand endpoint (which is formatted as `tri1.us1.account.AC.../registration.BU...`). Retrieve the BN brand SID via [GET /v1/a2p/BrandRegistrations/{Sid}](https://www.twilio.com/docs/messaging/api/brand-registration-resource) on `messaging.twilio.com`. |
**MessagingServiceSid** | **string** | The SID of the Messaging Service. |[optional] 
**ThemeSetId** | **string** | Theme ID for styling the inquiry form. |[optional] 
**UseCaseCategories** | **[]string** | The categories of the messaging use case. |[optional] 
**UseCaseDescription** | **string** | A description of the messaging use case. |[optional] 
**UseCaseSampleMessage1** | **string** | A sample message for the use case. |[optional] 
**UseCaseSampleMessage2** | **string** | A second sample message for the use case. |[optional] 
**UseCaseSampleMessage3** | **string** | A third sample message for the use case. |[optional] 
**UseCaseSampleMessage4** | **string** | A fourth sample message for the use case. |[optional] 
**UseCaseSampleMessage5** | **string** | A fifth sample message for the use case. |[optional] 
**UseCaseOptInTypes** | **[]string** | The opt-in methods for the use case. |[optional] 
**UseCaseOptInDescription** | **string** | A description of the opt-in process. |[optional] 
**HasEmbeddedLinks** | **bool** | Whether messages will contain embedded links. |[optional] 
**HasEmbeddedPhone** | **bool** | Whether messages will contain embedded phone numbers. |[optional] 
**EmbeddedUrlSample** | **string** | A sample URL that will be embedded in messages (must use https://). |[optional] 
**DirectLending** | **bool** | Whether the campaign involves direct lending. |[optional] 
**AgeGated** | **bool** | Whether the content is age-gated. |[optional] 
**PrivacyPolicyUrl** | **string** | The URL to the privacy policy. |[optional] 
**TermsAndConditionsUrl** | **string** | The URL to the terms and conditions. |[optional] 
**OptInKeywords** | **[]string** | Keywords users can text to opt in. |[optional] 
**OptInMessageSample** | **string** | A sample opt-in confirmation message. |[optional] 
**OptOutKeywords** | **[]string** | Keywords users can text to opt out. |[optional] 
**OptOutMessageSample** | **string** | A sample opt-out confirmation message. |[optional] 
**HelpKeywords** | **[]string** | Keywords users can text for help. |[optional] 
**HelpMessageSample** | **string** | A sample help message. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


