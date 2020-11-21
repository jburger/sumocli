# MetricsSavedSearchSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The item type. Dashboard links are not supported. | [default to null]
**Name** | **string** | The name of the item. | [default to null]
**Description** | **string** | Item description in the content library. | [optional] [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [default to null]
**LogQuery** | **string** | Query used to add an overlay to the chart. | [optional] [default to null]
**MetricsQueries** | [**[]MetricsSavedSearchQuerySyncDefinition**](MetricsSavedSearchQuerySyncDefinition.md) | Metrics queries. | [default to null]
**DesiredQuantizationInSecs** | **int32** | Desired quantization in seconds. | [default to null]
**Properties** | **string** | Chart properties. This field is optional. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

