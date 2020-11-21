# Variable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the variable. | [optional] [default to null]
**Name** | **string** | Name of the variable. The variable name is case-insensitive. Only alphanumeric, and underscores are allowed in the variable name.  | [default to null]
**DisplayName** | **string** | Display name of the variable shown in the UI. If this field is empty, the name field will be used. The display name is case-insensitive. Only numbers, and underscores are allowed in the variable name. This field is not yet supported by the UI.  | [optional] [default to null]
**DefaultValue** | **string** | Default value of the variable. | [optional] [default to null]
**SourceDefinition** | [***VariableSourceDefinition**](VariableSourceDefinition.md) |  | [default to null]
**AllowMultiSelect** | **bool** | Allow multiple selections in the values dropdown. | [optional] [default to false]
**IncludeAllOption** | **bool** | Include an \&quot;All\&quot; option at the top of the variable&#x27;s values dropdown. | [optional] [default to true]
**HideFromUI** | **bool** | Hide the variable in the dashboard UI. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

