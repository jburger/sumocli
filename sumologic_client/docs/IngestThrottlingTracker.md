# IngestThrottlingTracker

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Event type. | [optional] [default to null]
**TrackerId** | **string** | Name that uniquely identifies the health event. It focuses on what happened rather than why. | [default to null]
**Error_** | **string** | Description of the underlying reason for the event change. | [default to null]
**Description** | **string** | A more elaborate description of why the event occurred. | [default to null]
**DataType** | **string** | The type of data for which the rate limit was enabled. The possible values are &#x60;LogIngest&#x60; and &#x60;MetricsIngest&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

