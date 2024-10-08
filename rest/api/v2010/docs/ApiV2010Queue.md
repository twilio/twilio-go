# ApiV2010Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateUpdated** | Pointer to **string** | The date and time in GMT that this resource was last updated, specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**CurrentSize** | **int** | The number of calls currently in the queue. |[optional] [default to 0]
**FriendlyName** | Pointer to **string** | A string that you assigned to describe this resource. |
**Uri** | Pointer to **string** | The URI of this resource, relative to `https://api.twilio.com`. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Queue resource. |
**AverageWaitTime** | **int** |  The average wait time in seconds of the members in this queue. This is calculated at the time of the request. |[optional] [default to 0]
**Sid** | Pointer to **string** | The unique string that that we created to identify this Queue resource. |
**DateCreated** | Pointer to **string** | The date and time in GMT that this resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**MaxSize** | **int** |  The maximum number of calls that can be in the queue. The default is 1000 and the maximum is 5000. |[optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


