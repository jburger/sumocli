# AtLeastOnceOverThreshold

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | **string** |  | [default to null]
**Threshold** | **float64** | Threshold for the monitor, data points above or below this threshold are treated as outliers. | [default to null]
**ThresholdType** | **string** | One of: &#x60;Above&#x60;, &#x60;Below&#x60;. | [default to null]
**TimeWindow** | **string** | Monitored time window. If there is a single threshold violation in this time frame then notification will be triggered. After all outliers go out of scope, the alert should be resolved. Currently, the only accepted values are &#x60;5m&#x60;, &#x60;10m&#x60;, &#x60;15m&#x60;, &#x60;30m&#x60;, &#x60;60m&#x60;, &#x60;2h&#x60;, &#x60;3h&#x60;, &#x60;6h&#x60;, &#x60;12h&#x60;, &#x60;24h&#x60;. | [default to null]
**Notifications** | [***Notifications**](Notifications.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

