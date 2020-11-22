# WebhookSearchNotificationSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskType** | **string** | Delivery channel for notifications. | [default to null]
**WebhookId** | **string** | Identifier of the webhook connection. | [default to null]
**Payload** | **string** | A JSON object in the format required by the target WebHook URL. For details on variables that can be used as parameters within your JSON object, please refer to Sumo Logic Doc Hub. | [optional] [default to null]
**ItemizeAlerts** | **bool** | If this field is set to true, one webhook per result will be sent when the trigger conditions are met | [optional] [default to false]
**MaxItemizedAlerts** | **int32** | The maximum number of results for which we send separate alerts. This value should be between 1 and 100. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

