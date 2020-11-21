# LookupUpdateDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | **int32** | A time to live for each entry in the lookup table (in minutes). 0 is a special value. A TTL of 0 implies entry will never be deleted from the table. | [default to 0]
**Description** | **string** | The description of the lookup table. The description cannot be blank. | [default to null]
**SizeLimitAction** | **string** | The action that needs to be taken when the size limit is reached for the table. The possible values can be &#x60;StopIncomingMessages&#x60; or &#x60;DeleteOldData&#x60;. DeleteOldData will starting deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached. | [optional] [default to StopIncomingMessages]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

