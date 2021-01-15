# CreateQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Constraints the query to a given Field with an task. Useful when you know the Field you are expecting. It accepts one field in the format *task-unique-name-1*:*field-unique-name* | [optional] 
**Language** | **string** | An ISO language-country string of the sample. | 
**ModelBuild** | **string** | The Model Build Sid or unique name of the Model Build to be queried. | [optional] 
**Query** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. It can be up to 2048 characters long. | 
**Tasks** | **string** | Constraints the query to a set of tasks. Useful when you need to constrain the paths the user can take. Tasks should be comma separated *task-unique-name-1*, *task-unique-name-2* | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


