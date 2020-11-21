# WebhookConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Type of connection. Valid values are &#x60;WebhookConnection&#x60;, &#x60;ServiceNowConnection&#x60;. | [default to null]
**Id** | **string** | Unique identifier for the connection. | [default to null]
**Name** | **string** | Name of the connection. | [default to null]
**Description** | **string** | Description of the connection. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**Url** | **string** | URL for the webhook connection. | [default to null]
**Headers** | [**[]Header**](Header.md) | List of access authorization headers. | [default to null]
**CustomHeaders** | [**[]Header**](Header.md) | List of custom webhook headers. | [default to null]
**DefaultPayload** | **string** | Default payload of the webhook. | [default to null]
**WebhookType** | **string** | Type of webhook. Valid values are &#x60;AWSLambda&#x60;, &#x60;Azure&#x60;, &#x60;Datadog&#x60;, &#x60;HipChat&#x60;, &#x60;NewRelic&#x60;, &#x60;Opsgenie&#x60;, &#x60;PagerDuty&#x60;, &#x60;Slack&#x60; and &#x60;Webhook&#x60;. &#x60;Jira&#x60; and &#x60;Opsgenie&#x60; webhooks are in beta and not available until given access. Please reach out to your Sumo Logic representative to add new webhook types. | [default to null]
**Warnings** | **[]string** | Webhook endpoint warning for incorrect variable names and syntax. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

