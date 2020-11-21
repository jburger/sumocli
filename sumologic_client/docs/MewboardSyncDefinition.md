# MewboardSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The item type. Dashboard links are not supported. | [default to null]
**Name** | **string** | The name of the item. | [default to null]
**Description** | **string** | A description of the dashboard. | [optional] [default to null]
**Title** | **string** | The title of the dashboard. | [default to null]
**RootPanel** | [***ContainerPanel**](ContainerPanel.md) |  | [optional] [default to null]
**Theme** | **string** | Theme for the dashboard. Must be &#x60;light&#x60; or &#x60;dark&#x60;. | [optional] [default to light]
**TopologyLabelMap** | [***TopologyLabelMap**](TopologyLabelMap.md) |  | [optional] [default to null]
**RefreshInterval** | **int32** | Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. | [optional] [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] [default to null]
**Layout** | [***Layout**](Layout.md) |  | [optional] [default to null]
**Panels** | [**[]Panel**](Panel.md) | Children panels that the container panel contains. | [optional] [default to null]
**Variables** | [**[]Variable**](Variable.md) | Variables that could be applied to the panel&#x27;s children. | [optional] [default to null]
**ColoringRules** | [**[]ColoringRule**](ColoringRule.md) | Coloring rules to color the panel/data with. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

