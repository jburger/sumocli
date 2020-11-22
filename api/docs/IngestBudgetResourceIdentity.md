# IngestBudgetResourceIdentity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the resource. | [default to null]
**Name** | **string** | The name of the resource. | [optional] [default to Unknown]
**Type_** | **string** | -&gt; Resource type. Supported types are - &#x60;Collector&#x60;, &#x60;Source&#x60;, &#x60;IngestBudget&#x60; and &#x60;Organisation&#x60;. | [default to null]
**IngestBudgetFieldValue** | **string** | The unique field value of the ingest budget v1. This will be empty for v2 budgets. | [optional] [default to Unknown]
**Scope** | **string** | The scope of the ingest budget v2. This will be empty for v1 budgets. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

