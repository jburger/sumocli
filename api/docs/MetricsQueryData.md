# MetricsQueryData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The metric of the query. | [default to null]
**AggregationType** | **string** | The type of aggregation. Can be &#x60;Count&#x60;, &#x60;Minimum&#x60;, &#x60;Maximum&#x60;, &#x60;Sum&#x60;, &#x60;Average&#x60; or &#x60;None&#x60;. | [optional] [default to null]
**GroupBy** | **string** | The field to group the results by. | [optional] [default to null]
**Filters** | [**[]MetricsFilter**](MetricsFilter.md) | A list of filters for the metrics query. | [default to null]
**Operators** | [**[]OperatorData**](OperatorData.md) | A list of operator data for the metrics query. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

