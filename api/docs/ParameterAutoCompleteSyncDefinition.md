# ParameterAutoCompleteSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCompleteType** | **string** | The autocomplete parameter type. Supported values are:   1. &#x60;SKIP_AUTOCOMPLETE&#x60;   2. &#x60;CSV_AUTOCOMPLETE&#x60;   3. &#x60;AUTOCOMPLETE_KEY&#x60;   4. &#x60;VALUE_ONLY_AUTOCOMPLETE&#x60;   5. &#x60;VALUE_ONLY_LOOKUP_AUTOCOMPLETE&#x60;   6. &#x60;LABEL_VALUE_LOOKUP_AUTOCOMPLETE&#x60; | [default to null]
**AutoCompleteKey** | **string** | The autocomplete key to be used to fetch autocomplete values. | [optional] [default to null]
**AutoCompleteValues** | [**[]AutoCompleteValueSyncDefinition**](AutoCompleteValueSyncDefinition.md) | The array of values of the corresponding autocomplete parameter. | [optional] [default to null]
**LookupFileName** | **string** | The lookup file to use as a source for autocomplete values. | [optional] [default to null]
**LookupLabelColumn** | **string** | The column from the lookup file to use for autocomplete labels. | [optional] [default to null]
**LookupValueColumn** | **string** | The column from the lookup file to fill the actual value when a particular label is selected. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

