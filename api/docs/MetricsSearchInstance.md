# MetricsSearchInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**Title** | **string** | Item title in the content library. | [default to null]
**Description** | **string** | Item description in the content library. | [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [default to null]
**LogQuery** | **string** | Log query used to add an overlay to the chart. | [optional] [default to null]
**MetricsQueries** | [**[]MetricsSearchQuery**](MetricsSearchQuery.md) | Metrics queries, up to the maximum of six. | [default to null]
**DesiredQuantizationInSecs** | **int32** | Desired quantization in seconds. | [optional] [default to 0]
**Properties** | **string** | Chart properties, like line width, color palette, and the fill missing data method. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.  | [optional] [default to null]
**Id** | **string** | Identifier of the metrics search. | [default to null]
**ParentId** | **string** | Identifier of the parent element in the content library, such as folder. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

