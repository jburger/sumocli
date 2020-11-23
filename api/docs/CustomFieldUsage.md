# CustomFieldUsage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Field name. | [default to null]
**FieldId** | **string** | Identifier of the field. | [default to null]
**DataType** | **string** | Field type. Possible values are &#x60;String&#x60;, &#x60;Long&#x60;, &#x60;Int&#x60;, &#x60;Double&#x60;, &#x60;Boolean&#x60;. | [default to null]
**State** | **string** | Indicates whether the field is enabled and its values are being accepted. Possible values are &#x60;Enabled&#x60; and &#x60;Disabled&#x60;. | [default to null]
**FieldExtractionRules** | **[]string** | An array of hexadecimal identifiers of field extraction rules which use this field. | [optional] [default to null]
**Roles** | **[]string** | An array of hexadecimal identifiers of roles which use this field in the search filter. | [optional] [default to null]
**Partitions** | **[]string** | An array of hexadecimal identifiers of partitions which use this field in the routing expression. | [optional] [default to null]
**CollectorsCount** | **int32** | Total number of collectors using this field. | [optional] [default to null]
**SourcesCount** | **int32** | Total number of sources using this field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
