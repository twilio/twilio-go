# CreateCampaignsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandRegistrationSid** | **string** | A2P BrandRegistration Sid | 
**Description** | **string** | A short description of what this SMS campaign does. | 
**HasEmbeddedLinks** | **bool** | Indicate that this SMS campaign will send messages that contain links. | 
**HasEmbeddedPhone** | **bool** | Indicates that this SMS campaign will send messages that contain phone numbers. | 
**MessageSamples** | **[]string** | Message samples, up to 5 sample messages, &lt;&#x3D;255 chars each. Example: [ \&quot;EXPRESS: Denim Days Event is ON\&quot;, \&quot;LAST CHANCE: Book your next flight for just 1 (ONE) EUR\&quot; ] | 
**MessagingServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from. | 
**UseCase** | **string** | A2P Campaign UseCase. One of [ 2FA, EMERGENCY, MARKETING ] | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


