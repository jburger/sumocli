# SumoSearchPanel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the panel. | [optional] [default to null]
**Key** | **string** | Key for the panel. Used to create searches for the queries in the panel and configure the layout of the panel in the dashboard.  | [default to null]
**Title** | **string** | Title of the panel. | [optional] [default to null]
**VisualSettings** | **string** | Visual settings of the panel. | [optional] [default to null]
**KeepVisualSettingsConsistentWithParent** | **bool** | Keeps the visual settings, like series colors, consistent with the settings of the parent panel. | [optional] [default to true]
**PanelType** | **string** | Type of panel. | [default to null]
**Queries** | [**[]Query**](Query.md) | Metrics and log queries of the panel. | [default to null]
**Description** | **string** | Description of the panel. | [optional] [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] [default to null]
**ColoringRules** | [**[]ColoringRule**](ColoringRule.md) | Rules to set the color of data. | [optional] [default to null]
**LinkedDashboards** | [**[]LinkedDashboard**](LinkedDashboard.md) | List of linked dashboards. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
