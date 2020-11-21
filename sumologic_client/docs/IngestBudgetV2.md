# IngestBudgetV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the ingest budget. | [default to null]
**Scope** | **string** | A scope is a constraint that will be used to identify the messages on which budget needs to be applied. A scope is consists of key and value separated by &#x3D;. The field must be enabled in the fields table. Value supports wildcard. e.g. _sourceCategory&#x3D;*prod*payment*, cluster&#x3D;kafka. If the scope is defined _sourceCategory&#x3D;*nginx* in this budget will be applied on messages having fields _sourceCategory&#x3D;prod/nginx, _sourceCategory&#x3D;dev/nginx, or _sourceCategory&#x3D;dev/nginx/error | [default to null]
**CapacityBytes** | **int64** | Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit. | [default to null]
**Timezone** | **string** | Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | [default to null]
**ResetTime** | **string** | Reset time of the ingest budget in HH:MM format. | [default to null]
**Description** | **string** | Description of the ingest budget. | [optional] [default to null]
**Action** | **string** | Action to take when ingest budget&#x27;s capacity is reached. All actions are audited. Supported values are:   * &#x60;stopCollecting&#x60;   * &#x60;keepCollecting&#x60; | [default to null]
**AuditThreshold** | **int32** | The threshold as a percentage of when an ingest budget&#x27;s capacity usage is logged in the Audit Index. | [optional] [default to null]
**Id** | **string** | Unique identifier for the ingest budget. | [default to null]
**UsageBytes** | **int64** | Current usage since the last reset, in bytes. | [optional] [default to null]
**UsageStatus** | **string** | Status of the current usage. Can be &#x60;Normal&#x60;, &#x60;Approaching&#x60;, &#x60;Exceeded&#x60;, or &#x60;Unknown&#x60; (unable to retrieve usage). | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The creation timestamp in UTC of the Ingest Budget. | [default to null]
**CreatedBy** | **string** | The identifier of the user who created the Ingest Budget. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | The modified timestamp in UTC of the Ingest Budget. | [default to null]
**ModifiedBy** | **string** | The identifier of the user who modified the Ingest Budget. | [default to null]
**BudgetVersion** | **int32** | The version of the Ingest Budget | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

