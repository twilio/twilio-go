# MessagingV1UsAppToPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies a US A2P Compliance resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**BrandRegistrationSid** | Pointer to **string** | A2P Brand Registration SID |
**MessagingServiceSid** | Pointer to **string** | The SID of the Messaging Service the resource is associated with |
**Description** | Pointer to **string** | A short description of what this SMS campaign does |
**MessageSamples** | Pointer to **[]string** | Message samples |
**UsAppToPersonUsecase** | Pointer to **string** | A2P Campaign Use Case. |
**HasEmbeddedLinks** | Pointer to **bool** | Indicate that this SMS campaign will send messages that contain links |
**HasEmbeddedPhone** | Pointer to **bool** | Indicates that this SMS campaign will send messages that contain phone numbers |
**CampaignStatus** | Pointer to **string** | Campaign status |
**CampaignId** | Pointer to **string** | The Campaign Registry (TCR) Campaign ID. |
**IsExternallyRegistered** | Pointer to **bool** | Indicates whether the campaign was registered externally or not |
**RateLimits** | Pointer to **interface{}** | Rate limit and/or classification set by each carrier |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the US App to Person resource |
**Mock** | Pointer to **bool** | A boolean that specifies whether campaign is a mock or not. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


