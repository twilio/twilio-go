# MessagingV1UsAppToPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies a US A2P Compliance resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**BrandRegistrationSid** | **string** | A2P Brand Registration SID |[optional] 
**MessagingServiceSid** | **string** | The SID of the Messaging Service the resource is associated with |[optional] 
**Description** | **string** | A short description of what this SMS campaign does |[optional] 
**MessageSamples** | **[]string** | Message samples |[optional] 
**UsAppToPersonUsecase** | **string** | A2P Campaign Use Case. |[optional] 
**HasEmbeddedLinks** | **bool** | Indicate that this SMS campaign will send messages that contain links |[optional] 
**HasEmbeddedPhone** | **bool** | Indicates that this SMS campaign will send messages that contain phone numbers |[optional] 
**CampaignStatus** | **string** | Campaign status |[optional] 
**CampaignId** | **string** | The Campaign Registry (TCR) Campaign ID. |[optional] 
**IsExternallyRegistered** | **bool** | Indicates whether the campaign was registered externally or not |[optional] 
**RateLimits** | Pointer to **interface{}** | Rate limit and/or classification set by each carrier |
**MessageFlow** | **string** | Consumer opt-in flow |[optional] 
**OptInMessage** | **string** | Opt In Message |[optional] 
**OptOutMessage** | **string** | Opt Out Message |[optional] 
**HelpMessage** | **string** | Help Message |[optional] 
**OptInKeywords** | **[]string** | Opt In Keywords |[optional] 
**OptOutKeywords** | **[]string** | Opt Out Keywords |[optional] 
**HelpKeywords** | **[]string** | Help Keywords |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the US App to Person resource |[optional] 
**Mock** | **bool** | A boolean that specifies whether campaign is a mock or not. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


