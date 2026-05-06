# IdentifierConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdType** | **string** | Name of the identifier type. Usual values are email, phone, external_id etc. |
**MatchingAlgo** | **string** | The algorithm to use for matching identifier values.    - `exact`, exact string match.   - `fuzzy`, low precision match allowing for some variations. |[optional] [default to "exact"]
**MatchingThreshold** | **int** | The `fuzzy` matching threshold percentage. |[optional] [default to 75]
**Limit** | **int** | Maximum number of historical values to retain. |[optional] [default to 100]
**LimitPolicy** | **string** | Removal policy to apply when the number of values exceeds the limit and is based on the timestamp of the request when the identifier was added. - `fifo`: First In First Out, removes the oldest values first. - `lifo`: Last In First Out, removes the most recent values first. |[optional] [default to "fifo"]
**EnforceUnique** | **bool** | When enabled, more than one profile may not share the same identifier value. Adding a shared identifier to a second profile may trigger a merge. Disabling creates a compound identifier where merges are only triggered if two or more identifiers satisfy a matching rule. |[optional] [default to true]
**Normalization** | **string** | Normalization to apply to the identifier value before storing and matching. - `phone`: Normalize phone numbers to E.164 format. - `email`: Normalize email addresses by coverting to lowercase and removing spaces. - `trim`: Removes spaces from both ends of the string. - `none`: No normalization.  **Important Note for Phone Number Normalization:** When using the `phone` normalization option, please adhere to the following guidelines to ensure proper formatting: - All US numbers must be at least valid 10 digit phone numbers with area code. They may also include or omit the \"+1\" prefix. - Non-US numbers must include a 1-3 digit country code (a \"+\" prefix is optional). Non-US national formats are not supported and will be interpreted as US numbers. |[optional] [default to "trim"]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


