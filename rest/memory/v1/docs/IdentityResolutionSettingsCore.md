# IdentityResolutionSettingsCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierConfigs** | [**[]IdentifierConfig**](IdentifierConfig.md) | List of identifier types and their resolution settings. |[optional] 
**MatchingRules** | **[]string** | Priority list of identifiers to locate profiles to apply new data to, or for  determining if two existing profiles should merge. Rules are evaluated in order. - If no rule matches against existing profiles, a new profile will be created.  - If a rule matches to a single existing profile, the profile will be updated.  - If a rule matches to multiple existing profiles, those existing profiles will be merged. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


