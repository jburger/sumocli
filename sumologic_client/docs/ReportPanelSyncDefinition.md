# ReportPanelSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The title of the panel. | [default to null]
**ViewerType** | **string** | Type of [area chart](https://help.sumologic.com/Dashboards-and-Alerts/Dashboards/Chart-Panel-Types). Supported values are:   1. &#x60;table&#x60; for Table   2. &#x60;bar&#x60; for Bar Chart   3. &#x60;column&#x60; for Column Chart   4. &#x60;line&#x60; for Line Chart   5. &#x60;area&#x60; for Area Chart   6. &#x60;pie&#x60; for Pie Chart   7. &#x60;svv&#x60; for Single Value Viewer   8. &#x60;title&#x60; for Title Panel   9. &#x60;text&#x60; for Text Panel  Values 1-7 are used for Data Panels. | [default to null]
**DetailLevel** | **int32** | Supported values are:   - &#x60;1&#x60; for small   - &#x60;2&#x60; for medium   - &#x60;3&#x60; for large | [default to null]
**QueryString** | **string** | The query to run, for panels associated to log searches. | [default to null]
**MetricsQueries** | [**[]MetricsQuerySyncDefinition**](MetricsQuerySyncDefinition.md) | The query or queries to run, for panels associated to metrics searches. | [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [default to null]
**X** | **int32** | The horizontal position of the panel. A sumo screen is divided into 24 columns. The value for x can be any integer from 0 to 24. | [default to null]
**Y** | **int32** | The vertical position of the panel. A sumo screen is divided into 24 rows. The value for y can be any integer from 0 to 24. | [default to null]
**Width** | **int32** | The width of the panel. | [default to null]
**Height** | **int32** | The height of the panel. | [default to null]
**Properties** | **string** | Visual settings for the panel. | [default to null]
**Id** | **string** | A string identifier that you can use to refer to the panel in filters.panelIds. | [default to null]
**DesiredQuantizationInSecs** | **int32** | The quantization interval aligns your time series data to common intervals on the time axis (for example every one minute) to optimize the visualization and performance. | [optional] [default to null]
**QueryParameters** | [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | The parameters for parameterized searches. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

