# IngestBudgetDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the ingest budget. | [default to null]
**FieldValue** | **string** | Custom field value that is used to assign Collectors to the ingest budget. | [default to null]
**CapacityBytes** | **int64** | Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit. | [default to null]
**Timezone** | **string** | Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | [default to null]
**ResetTime** | **string** | Reset time of the ingest budget in HH:MM format. | [default to null]
**Description** | **string** | Description of the ingest budget. | [optional] [default to null]
**Action** | **string** | Action to take when ingest budget&#x27;s capacity is reached. All actions are audited. Supported values are:   * &#x60;stopCollecting&#x60;   * &#x60;keepCollecting&#x60; | [default to null]
**AuditThreshold** | **int32** | The threshold as a percentage of when an ingest budget&#x27;s capacity usage is logged in the Audit Index. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

