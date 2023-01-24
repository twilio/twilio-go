# VoiceV1DialingPermissionsCountryBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateCount** | Pointer to **int** | The number of countries updated |
**UpdateRequest** | Pointer to **string** | A bulk update request to change voice dialing country permissions stored as a URL-encoded, JSON array of update objects. For example : `[ { \"iso_code\": \"GB\", \"low_risk_numbers_enabled\": \"true\", \"high_risk_special_numbers_enabled\":\"true\", \"high_risk_tollfraud_numbers_enabled\": \"false\" } ]` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


