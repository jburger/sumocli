# MonitorsLibraryMonitorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the monitor or folder. | [default to null]
**Name** | **string** | Identifier of the monitor or folder. | [default to null]
**Description** | **string** | Description of the monitor or folder. | [default to null]
**Version** | **int64** | Version of the monitor or folder. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**ParentId** | **string** | Identifier of the parent folder. | [default to null]
**ContentType** | **string** | Type of the content. Valid values:   1) Monitor   2) Folder | [default to null]
**Type_** | **string** | Type of the object model. | [default to null]
**IsSystem** | **bool** | System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can&#x27;t be updated. | [default to null]
**IsMutable** | **bool** | Immutable objects are \&quot;READ-ONLY\&quot;. | [default to null]
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor. | [default to null]
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | [default to null]
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | [default to null]
**Notifications** | [**[]MonitorNotification**](MonitorNotification.md) | The notifications the monitor will send when the respective trigger condition is met. | [optional] [default to []]
**IsDisabled** | **bool** | Whether or not the monitor is disabled. Disabled monitors will not run, and will not generate or send notifications. | [optional] [default to false]
**Status** | **[]string** | The current status of the monitor. Each monitor can have one or more status values. Valid values:   1. &#x60;Normal&#x60;: The monitor is running normally and does not have any currently triggered conditions.   2. &#x60;Critical&#x60;: The Critical trigger condition has been met.   3. &#x60;Warning&#x60;: The Warning trigger condition has been met.   4. &#x60;MissingData&#x60;: The MissingData trigger condition has been met.   5. &#x60;Disabled&#x60;: The monitor has been disabled and is not currently running. | [optional] [default to null]
**GroupNotifications** | **bool** | Whether or not to group notifications for individual items that meet the trigger condition. | [optional] [default to true]
**Warnings** | **map[string]string** | Monitor manager warnings | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

