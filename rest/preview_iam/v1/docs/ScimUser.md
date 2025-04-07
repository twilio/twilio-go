# ScimUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Twilio user sid |[optional] 
**ExternalId** | **string** | External unique resource id defined by provisioning client |[optional] 
**UserName** | **string** | Unique username, MUST be same as primary email address |
**DisplayName** | **string** | User friendly display name |[optional] 
**Name** | [**ScimName**](ScimName.md) |  |[optional] 
**Emails** | [**[]ScimEmailAddress**](ScimEmailAddress.md) | Email address list of the user. Primary email must be defined if there are more than 1 email. Primary email must match the username. |[optional] 
**Active** | **bool** | Indicates whether the user is active |[optional] 
**Locale** | **string** | User's locale |[optional] 
**Timezone** | **string** | User's time zone |[optional] 
**Schemas** | **[]string** | An array of URIs that indicate the schemas supported for this user resource |[optional] 
**Meta** | [**ScimMeta**](ScimMeta.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


