# ExtractionRuleDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field extraction rule. Use a name that makes it easy to identify the rule. | [default to null]
**Scope** | **string** | Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You&#x27;ll use the Scope to run a search against the rule. | [default to null]
**ParseExpression** | **string** | Describes the fields to be parsed. | [default to null]
**Enabled** | **bool** | Is the field extraction rule enabled. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

