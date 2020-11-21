# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtractionRule**](ExtractionRuleManagementApi.md#CreateExtractionRule) | **Post** /v1/extractionRules | Create a new field extraction rule.
[**DeleteExtractionRule**](ExtractionRuleManagementApi.md#DeleteExtractionRule) | **Delete** /v1/extractionRules/{id} | Delete a field extraction rule.
[**GetExtractionRule**](ExtractionRuleManagementApi.md#GetExtractionRule) | **Get** /v1/extractionRules/{id} | Get a field extraction rule.
[**ListExtractionRules**](ExtractionRuleManagementApi.md#ListExtractionRules) | **Get** /v1/extractionRules | Get a list of field extraction rules.
[**UpdateExtractionRule**](ExtractionRuleManagementApi.md#UpdateExtractionRule) | **Put** /v1/extractionRules/{id} | Update a field extraction rule.

# **CreateExtractionRule**
> ExtractionRule CreateExtractionRule(ctx, body)
Create a new field extraction rule.

Create a new field extraction rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExtractionRuleDefinition**](ExtractionRuleDefinition.md)| Information about the new field extraction rule. | 

### Return type

[**ExtractionRule**](ExtractionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExtractionRule**
> DeleteExtractionRule(ctx, id)
Delete a field extraction rule.

Delete a field extraction rule with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the field extraction rule to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtractionRule**
> ExtractionRule GetExtractionRule(ctx, id)
Get a field extraction rule.

Get a field extraction rule with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of field extraction rule to return. | 

### Return type

[**ExtractionRule**](ExtractionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExtractionRules**
> ListExtractionRulesResponse ListExtractionRules(ctx, optional)
Get a list of field extraction rules.

Get a list of all field extraction rules. The response is paginated with a default limit of 100 field extraction rules per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtractionRuleManagementApiListExtractionRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtractionRuleManagementApiListExtractionRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of field extraction rules returned in the response. The number of field extraction rules returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListExtractionRulesResponse**](ListExtractionRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtractionRule**
> ExtractionRule UpdateExtractionRule(ctx, body, id)
Update a field extraction rule.

Update an existing field extraction rule. All properties specified in the request are replaced. Missing properties are set to their default values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateExtractionRuleDefinition**](UpdateExtractionRuleDefinition.md)| Information to update about the field extraction rule. | 
  **id** | **string**| Identifier of the field extraction rule to update. | 

### Return type

[**ExtractionRule**](ExtractionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

