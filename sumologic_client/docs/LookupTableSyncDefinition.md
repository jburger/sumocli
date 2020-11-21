# LookupTableSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the lookup table. | [default to null]
**Fields** | [**[]LookupTableField**](LookupTableField.md) | The list of fields in the lookup table. | [default to null]
**PrimaryKeys** | **[]string** | The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain. | [default to null]
**Ttl** | **int32** | A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically. | [optional] [default to 0]
**SizeLimitAction** | **string** | The action that needs to be taken when the size limit is reached for the table. The possible values can be &#x60;StopIncomingMessages&#x60; or &#x60;DeleteOldData&#x60;. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached. | [optional] [default to StopIncomingMessages]
**Type_** | **string** | The item type. Dashboard links are not supported. | [default to null]
**Name** | **string** | The name of the item. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

