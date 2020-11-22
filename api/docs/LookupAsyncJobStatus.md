# LookupAsyncJobStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | An identifier returned in response to an asynchronous request. | [default to null]
**Status** | **string** | Whether or not the request is pending (&#x60;Pending&#x60;), in progress (&#x60;InProgress&#x60;), has completed successfully (&#x60;Success&#x60;), has completed partially with warnings (&#x60;PartialSuccess&#x60;) or has completed with an error (&#x60;Failed&#x60;). | [default to null]
**StatusMessages** | **[]string** | Additional status messages generated if any if the status is &#x60;Success&#x60;. | [optional] [default to null]
**Errors** | [**[]ErrorDescription**](ErrorDescription.md) | More information about the failures, if the status is &#x60;Failed&#x60;. | [optional] [default to null]
**Warnings** | [**[]WarningDescription**](warningDescription.md) | More information about the warnings, if the status is &#x60;PartialSuccess&#x60;. | [optional] [default to null]
**LookupContentId** | **string** | Content id of lookup table on which this operation was performed. | [default to null]
**LookupName** | **string** | Name of lookup table on which this operation was performed. | [default to null]
**LookupContentPath** | **string** | Content path of lookup table on which this operation was performed. | [default to null]
**RequestType** | **string** | Type of asynchronous request made:   - &#x60;BulkMerge&#x60;   - &#x60;BulkReplace&#x60;   - &#x60;Truncate&#x60; | [optional] [default to null]
**UserId** | **string** | User id of user who initiated this operation. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time of this job in UTC. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Timestamp in UTC when status was last updated. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

