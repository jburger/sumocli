# SaveMetricsSearchRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Item title in the content library. | [default to null]
**Description** | **string** | Item description in the content library. | [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [default to null]
**LogQuery** | **string** | Log query used to add an overlay to the chart. | [optional] [default to null]
**MetricsQueries** | [**[]MetricsSearchQuery**](MetricsSearchQuery.md) | Metrics queries, up to the maximum of six. | [default to null]
**DesiredQuantizationInSecs** | **int32** | Desired quantization in seconds. | [optional] [default to 0]
**Properties** | **string** | Chart properties, like line width, color palette, and the fill missing data method. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.  | [optional] [default to null]
**ParentId** | **string** | Identifier of a folder to which the metrics search should be added. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

