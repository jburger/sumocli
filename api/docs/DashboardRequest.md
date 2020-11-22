# DashboardRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the dashboard. | [default to null]
**Description** | **string** | Description of the dashboard. | [optional] [default to null]
**FolderId** | **string** | The identifier of the folder to save the dashboard in. By default it is saved in your personal folder.  | [optional] [default to null]
**TopologyLabelMap** | [***TopologyLabelMap**](TopologyLabelMap.md) |  | [optional] [default to null]
**RefreshInterval** | **int32** | Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. This functionality is currently not supported.  | [optional] [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [default to null]
**Panels** | [**[]Panel**](Panel.md) | Panels in the dashboard. | [optional] [default to null]
**Layout** | [***Layout**](Layout.md) |  | [optional] [default to null]
**Variables** | [**[]Variable**](Variable.md) | Variables to apply to the panels. | [optional] [default to null]
**Theme** | **string** | Theme for the dashboard. Either &#x60;Light&#x60; or &#x60;Dark&#x60;. | [optional] [default to Light]
**ColoringRules** | [**[]ColoringRule**](ColoringRule.md) | Rules to set the color of data. This is an internal field and is not current supported by UI. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

