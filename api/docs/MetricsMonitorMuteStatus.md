# MetricsMonitorMuteStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMuted** | **bool** | True if the monitor is muted. | [default to null]
**MutedAt** | [**time.Time**](time.Time.md) | Date and time the monitor was last muted in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] [default to null]
**MutedUntil** | [**time.Time**](time.Time.md) | Date and time when the monitor will be unmuted in [RFC3339](https://tools.ietf.org/html/rfc3339) format. &#x60;null&#x60; when the monitor is muted indefinitely.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

