# MonitorsLibraryMonitorExport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the monitor or folder. | [default to null]
**Description** | **string** | Description of the monitor or folder. | [optional] [default to null]
**Type_** | **string** | Type of the object model. | [default to null]
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor. | [default to null]
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | [default to null]
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | [default to null]
**Notifications** | [**[]MonitorNotification**](MonitorNotification.md) | The notifications the monitor will send when the respective trigger condition is met. | [optional] [default to []]
**IsDisabled** | **bool** | Whether or not the monitor is disabled. Disabled monitors will not run, and will not generate or send notifications. | [optional] [default to false]
**GroupNotifications** | **bool** | Whether or not to group notifications for individual items that meet the trigger condition. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

