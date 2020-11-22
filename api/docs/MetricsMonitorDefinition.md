# MetricsMonitorDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Monitor name. | [default to null]
**Description** | **string** | Monitor description. | [optional] [default to null]
**AlertQueries** | [**[]MetricsAlertQuery**](MetricsAlertQuery.md) | Monitor queries. | [default to null]
**Timezone** | **string** | Monitor time zone in [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) format. Date time ranges shown in emails and sent to webhooks are expressed in this time zone.  | [default to null]
**MonitorRules** | [***MonitorRules**](MonitorRules.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

