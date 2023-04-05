# MessagingV1UsAppToPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies a US A2P Compliance resource `QE2c6890da8086d771620e9b13fadeba0b`. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the Campaign belongs to. |
**BrandRegistrationSid** | Pointer to **string** | The unique string to identify the A2P brand. |
**MessagingServiceSid** | Pointer to **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) that the resource is associated with. |
**Description** | Pointer to **string** | A short description of what this SMS campaign does. Min length: 40 characters. Max length: 4096 characters. |
**MessageSamples** | Pointer to **[]string** | Message samples, at least 1 and up to 5 sample messages (at least 2 for starter/sole proprietor), >=20 chars, <=1024 chars each. |
**UsAppToPersonUsecase** | Pointer to **string** | A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING, SOLE_PROPRIETOR...]. SOLE_PROPRIETOR campaign use cases can only be created by SOLE_PROPRIETOR Brands, and there can only be one SOLE_PROPRIETOR campaign created per SOLE_PROPRIETOR Brand. |
**HasEmbeddedLinks** | Pointer to **bool** | Indicate that this SMS campaign will send messages that contain links. |
**HasEmbeddedPhone** | Pointer to **bool** | Indicates that this SMS campaign will send messages that contain phone numbers. |
**CampaignStatus** | Pointer to **string** | Campaign status. Examples: IN_PROGRESS, VERIFIED, FAILED. |
**CampaignId** | Pointer to **string** | The Campaign Registry (TCR) Campaign ID. |
**IsExternallyRegistered** | Pointer to **bool** | Indicates whether the campaign was registered externally or not. |
**RateLimits** | Pointer to **interface{}** | Rate limit and/or classification set by each carrier, Ex. AT&T or T-Mobile. |
**MessageFlow** | Pointer to **string** | Details around how a consumer opts-in to their campaign, therefore giving consent to receive their messages. If multiple opt-in methods can be used for the same campaign, they must all be listed. 40 character minimum. 2048 character maximum. |
**OptInMessage** | Pointer to **string** | If end users can text in a keyword to start receiving messages from this campaign, the auto-reply messages sent to the end users must be provided. The opt-in response should include the Brand name, confirmation of opt-in enrollment to a recurring message campaign, how to get help, and clear description of how to opt-out. This field is required if end users can text in a keyword to start receiving messages from this campaign. 20 character minimum. 320 character maximum. |
**OptOutMessage** | Pointer to **string** | Upon receiving the opt-out keywords from the end users, Twilio customers are expected to send back an auto-generated response, which must provide acknowledgment of the opt-out request and confirmation that no further messages will be sent. It is also recommended that these opt-out messages include the brand name. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum. |
**HelpMessage** | Pointer to **string** | When customers receive the help keywords from their end users, Twilio customers are expected to send back an auto-generated response; this may include the brand name and additional support contact information. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum. |
**OptInKeywords** | Pointer to **[]string** | If end users can text in a keyword to start receiving messages from this campaign, those keywords must be provided. This field is required if end users can text in a keyword to start receiving messages from this campaign. Values must be alphanumeric. 255 character maximum. |
**OptOutKeywords** | Pointer to **[]string** | End users should be able to text in a keyword to stop receiving messages from this campaign. Those keywords must be provided. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum. |
**HelpKeywords** | Pointer to **[]string** | End users should be able to text in a keyword to receive help. Those keywords must be provided as part of the campaign registration request. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the US App to Person resource. |
**Mock** | Pointer to **bool** | A boolean that specifies whether campaign is a mock or not. Mock campaigns will be automatically created if using a mock brand. Mock campaigns should only be used for testing purposes. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


