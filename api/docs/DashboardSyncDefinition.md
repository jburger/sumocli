# DashboardSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The item type. Dashboard links are not supported. | [default to null]
**Name** | **string** | The name of the item. | [default to null]
**Description** | **string** | A description of the dashboard. | [default to null]
**DetailLevel** | **int32** | Supported values are:   - &#x60;1&#x60; for small   - &#x60;2&#x60; for medium   - &#x60;3&#x60; for large | [default to null]
**Properties** | **string** | Visual settings for the panel. | [default to null]
**Panels** | [**[]ReportPanelSyncDefinition**](ReportPanelSyncDefinition.md) | The panels of the dashboard. _Dashboard links are not supported._ | [default to null]
**Filters** | [**[]ReportFilterSyncDefinition**](ReportFilterSyncDefinition.md) | The filters for the dashboard. Filters allow you to control the amount of information displayed in your dashboards. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

