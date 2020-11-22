# HealthEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The unique identifier of the event. | [default to null]
**EventName** | **string** | The name of the event. | [default to null]
**Details** | [***TrackerIdentity**](TrackerIdentity.md) |  | [default to null]
**ResourceIdentity** | [***ResourceIdentity**](ResourceIdentity.md) |  | [default to null]
**EventTime** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**Subsystem** | **string** | The product area of the event. | [default to null]
**SeverityLevel** | **string** | The criticality of the event. It is either &#x60;Error&#x60; or &#x60;Warning&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

