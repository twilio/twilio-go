# TaskrouterV1WorkersCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**StartTime** | [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated |[optional] 
**EndTime** | [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated |[optional] 
**ActivityDurations** | **[]interface{}** | The minimum, average, maximum, and total time that Workers spent in each Activity |[optional] 
**ReservationsCreated** | **int** | The total number of Reservations that were created |[optional] 
**ReservationsAccepted** | **int** | The total number of Reservations that were accepted |[optional] 
**ReservationsRejected** | **int** | The total number of Reservations that were rejected |[optional] 
**ReservationsTimedOut** | **int** | The total number of Reservations that were timed out |[optional] 
**ReservationsCanceled** | **int** | The total number of Reservations that were canceled |[optional] 
**ReservationsRescinded** | **int** | The total number of Reservations that were rescinded |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the Workers |[optional] 
**Url** | **string** | The absolute URL of the Workers statistics resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


