# InsightsV1StatsBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Window** | [**InsightsV1Window**](InsightsV1Window.md) |  |[optional] 
**Calls** | **map[string]int** | Call counts keyed by call direction/kind (dynamic keys) |[optional] 
**Pdd** | [**InsightsV1Metrics**](InsightsV1Metrics.md) |  |[optional] 
**Duration** | [**InsightsV1Metrics**](InsightsV1Metrics.md) |  |[optional] 
**CallStates** | **map[string]int** | Count of calls per call state |[optional] 
**Direction** | **map[string]int** | Count of calls per direction |[optional] 
**Disconnected** | **map[string]int** | Count of calls by who disconnected (caller / callee / unknown) |[optional] 
**LastSipResponse** | **map[string]int** | Count of calls by last SIP response code |[optional] 
**CalleeCountry** | **map[string]int** | Count of calls by callee country code |[optional] 
**CallerCountry** | **map[string]int** | Count of calls by caller country code |[optional] 
**InsightTags** | **map[string]int** | Count of calls per insight tag |[optional] 
**SdkTags** | **map[string]int** | Count of calls per SDK tag |[optional] 
**SliTags** | **map[string]int** | Count of calls per SLI error tag |[optional] 
**InConference** | **map[string]int** | Count of calls split by whether they were in a conference |[optional] 
**CallDirection** | **map[string]int** | Count of calls per call direction |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


