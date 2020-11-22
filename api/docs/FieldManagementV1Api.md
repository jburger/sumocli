# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateField**](FieldManagementV1Api.md#CreateField) | **Post** /v1/fields | Create a new field.
[**DeleteField**](FieldManagementV1Api.md#DeleteField) | **Delete** /v1/fields/{id} | Delete a custom field.
[**DisableField**](FieldManagementV1Api.md#DisableField) | **Delete** /v1/fields/{id}/disable | Disable a custom field.
[**EnableField**](FieldManagementV1Api.md#EnableField) | **Put** /v1/fields/{id}/enable | Enable custom field with a specified identifier.
[**GetBuiltInField**](FieldManagementV1Api.md#GetBuiltInField) | **Get** /v1/fields/builtin/{id} | Get a built-in field.
[**GetCustomField**](FieldManagementV1Api.md#GetCustomField) | **Get** /v1/fields/{id} | Get a custom field.
[**GetFieldQuota**](FieldManagementV1Api.md#GetFieldQuota) | **Get** /v1/fields/quota | Get capacity information.
[**ListBuiltInFields**](FieldManagementV1Api.md#ListBuiltInFields) | **Get** /v1/fields/builtin | Get a list of built-in fields.
[**ListCustomFields**](FieldManagementV1Api.md#ListCustomFields) | **Get** /v1/fields | Get a list of all custom fields.
[**ListDroppedFields**](FieldManagementV1Api.md#ListDroppedFields) | **Get** /v1/fields/dropped | Get a list of dropped fields.

# **CreateField**
> CustomField CreateField(ctx, body)
Create a new field.

Adding a field will define it in the Fields schema allowing it to be assigned as metadata to your logs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FieldName**](FieldName.md)| Name of a field to add. The name is used as the key in the key-value pair. | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteField**
> DeleteField(ctx, id)
Delete a custom field.

Deleting a field does not delete historical data assigned with that field. If you  delete a field by mistake and one or more of those dependencies break, you can  re-add the field to get things working properly again. You should always disable  a field and ensure things are behaving as expected before deleting a field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of a field to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableField**
> DisableField(ctx, id)
Disable a custom field.

After disabling a field Sumo Logic will start dropping its incoming values at ingest. As a result, they won't be searchable or usable. Historical values are not removed and remain searchable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of a field to disable. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableField**
> EnableField(ctx, id)
Enable custom field with a specified identifier.

Fields have to be enabled to be assigned to your data. This operation ensures that a specified field is enabled and Sumo Logic will treat it as safe to process. All manually created custom fields are  enabled by default.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of a field to enable. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuiltInField**
> BuiltinField GetBuiltInField(ctx, id)
Get a built-in field.

Get the details of a built-in field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of a built-in field. | 

### Return type

[**BuiltinField**](BuiltinField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomField**
> CustomField GetCustomField(ctx, id)
Get a custom field.

Get the details of a custom field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of a field. | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFieldQuota**
> FieldQuotaUsage GetFieldQuota(ctx, )
Get capacity information.

Every account has a limited number of fields available. This endpoint returns your account limitations and remaining quota.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FieldQuotaUsage**](FieldQuotaUsage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBuiltInFields**
> ListBuiltinFieldsResponse ListBuiltInFields(ctx, )
Get a list of built-in fields.

Built-in fields are created automatically by Sumo Logic for standard configuration purposes. They include `_sourceHost` and `_sourceCategory`. Built-in fields can't be deleted or disabled.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListBuiltinFieldsResponse**](ListBuiltinFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomFields**
> ListCustomFieldsResponse ListCustomFields(ctx, )
Get a list of all custom fields.

Request a list of all the custom fields configured in your account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListCustomFieldsResponse**](ListCustomFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDroppedFields**
> ListDroppedFieldsResponse ListDroppedFields(ctx, )
Get a list of dropped fields.

Dropped fields are fields sent to Sumo Logic, but are ignored since they are not defined in your Fields schema. In order to save these values a field must both exist and be enabled.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListDroppedFieldsResponse**](ListDroppedFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

