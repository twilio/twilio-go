# LookupsV2PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingCountryCode** | **string** | International dialing prefix |[optional] 
**CountryCode** | **string** | Phone number's ISO country code |[optional] 
**PhoneNumber** | **string** | Phone number in E.164 format |[optional] 
**NationalFormat** | **string** | Phone number in national format |[optional] 
**Valid** | **bool** | Boolean which indicates if the phone number is valid |[optional] 
**ValidationErrors** | Pointer to [**[]string**](PhoneNumberEnumValidationError.md) | Contains reasons why a phone number is invalid |
**CallerName** | Pointer to **interface{}** | An object that contains caller name information |
**SimSwap** | Pointer to **interface{}** | An object that contains SIM swap information |
**CallForwarding** | Pointer to **interface{}** | An object that contains call forwarding status information |
**LiveActivity** | Pointer to **interface{}** | An object that contains live activity information |
**LineTypeIntelligence** | Pointer to **interface{}** | An object that contains line type information |
**IdentityMatch** | Pointer to **interface{}** | An object that contains identity match information |
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


