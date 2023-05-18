# NumbersV1BulkEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | The SID of the bulk eligibility check that you want to know about. |
**Url** | Pointer to **string** | This is the url of the request that you're trying to reach out to locate the resource. |
**Results** | Pointer to **[]interface{}** | The result set that contains the eligibility check response for each requested number, each result has at least the following attributes:  phone_number: The requested phone number ,hosting_account_sid: The account sid where the phone number will be hosted, country: Phone number’s country, eligibility_status: Indicates the eligibility status of the PN (Eligible/Ineligible), eligibility_sub_status: Indicates the sub status of the eligibility , ineligibility_reason: Reason for number's ineligibility (if applicable), next_step: Suggested next step in the hosting process based on the eligibility status. |
**FriendlyName** | Pointer to **string** | This is the string that you assigned as a friendly name for describing the eligibility check request. |
**Status** | Pointer to **string** | This is the status of the bulk eligibility check request. (Example: COMPLETE, IN_PROGRESS) |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) |  |
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


