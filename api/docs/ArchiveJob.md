# ArchiveJob

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the ingestion job. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The starting timestamp of the ingestion job. | [default to null]
**EndTime** | [**time.Time**](time.Time.md) | The ending timestamp of the ingestion job. | [default to null]
**Id** | **string** | The unique identifier of the ingestion job. | [default to null]
**TotalObjectsScanned** | **int64** | The total number of objects scanned by the ingestion job. | [default to null]
**TotalObjectsIngested** | **int64** | The total number of objects ingested by the ingestion job. | [default to null]
**TotalBytesIngested** | **int64** | The total bytes ingested by the ingestion job. | [default to null]
**Status** | **string** | The status of the ingestion job, either &#x60;Pending&#x60;,&#x60;Scanning&#x60;,&#x60;Ingesting&#x60;,&#x60;Failed&#x60;, or &#x60;Succeeded&#x60;. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The creation timestamp in UTC of the ingestion job. | [default to null]
**CreatedBy** | **string** | The identifier of the user who created the ingestion job. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

