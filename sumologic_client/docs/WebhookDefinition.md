# WebhookDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Type of connection. Valid values are &#x60;WebhookDefinition&#x60;, &#x60;ServiceNowDefinition&#x60;. | [default to null]
**Name** | **string** | Name of connection. Name should be a valid alphanumeric value. | [default to null]
**Description** | **string** | Description of the connection. | [optional] 
**Url** | **string** | URL for the webhook connection. | [default to null]
**Headers** | [**[]Header**](Header.md) | List of access authorization headers. | [optional] [default to []]
**CustomHeaders** | [**[]Header**](Header.md) | List of custom webhook headers. | [optional] [default to []]
**DefaultPayload** | **string** | Default payload of the webhook. | [default to null]
**WebhookType** | **string** | Type of webhook. Valid values are &#x60;AWSLambda&#x60;, &#x60;Azure&#x60;, &#x60;Datadog&#x60;, &#x60;HipChat&#x60;, &#x60;NewRelic&#x60;, &#x60;Opsgenie&#x60;, &#x60;PagerDuty&#x60;, &#x60;Slack&#x60; and &#x60;Webhook&#x60;. &#x60;Jira&#x60; and &#x60;Opsgenie&#x60; webhooks are in beta and not available until given access. Please reach out to your Sumo Logic representative to add new webhook types. | [optional] [default to Webhook]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

