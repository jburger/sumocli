# UpdateScheduledViewDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataForwardingId** | **string** | An optional ID of a data forwarding configuration to be used by the scheduled view. | [optional] [default to null]
**RetentionPeriod** | **int32** | The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled. | [optional] [default to -1]
**ReduceRetentionPeriodImmediately** | **bool** | This is required if the newly specified &#x60;retentionPeriod&#x60; is less than the existing retention period.  In such a situation, a value of &#x60;true&#x60; says that data between the existing retention period and the new retention period should be deleted immediately; if &#x60;false&#x60;, such data will be deleted after seven days. This property is optional and ignored if the specified &#x60;retentionPeriod&#x60; is greater than or equal to the current retention period. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

