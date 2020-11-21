# MissingDataRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedTimeSeries** | **string** | Defines when an alert should be raised: either when all or any time series are missing data. Accepted values for this field are: &#x60;all&#x60; and &#x60;any&#x60;.  | [default to null]
**TimeWindow** | **string** | Monitored time window. Currently, the only accepted values are &#x60;5m&#x60;, &#x60;10m&#x60;, &#x60;15m&#x60;, &#x60;30m&#x60;, &#x60;60m&#x60;, &#x60;2h&#x60;, &#x60;3h&#x60;, &#x60;6h&#x60;, &#x60;12h&#x60;, &#x60;24h&#x60;. The minimum value is equal to the query quantization. | [default to null]
**Notifications** | [***Notifications**](Notifications.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

