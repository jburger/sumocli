# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](TransformationRuleManagementApi.md#CreateRule) | **Post** /v1/transformationRules | Create a new transformation rule.
[**DeleteRule**](TransformationRuleManagementApi.md#DeleteRule) | **Delete** /v1/transformationRules/{id} | Delete a transformation rule.
[**GetTransformationRule**](TransformationRuleManagementApi.md#GetTransformationRule) | **Get** /v1/transformationRules/{id} | Get a transformation rule.
[**GetTransformationRules**](TransformationRuleManagementApi.md#GetTransformationRules) | **Get** /v1/transformationRules | Get a list of transformation rules.
[**UpdateTransformationRule**](TransformationRuleManagementApi.md#UpdateTransformationRule) | **Put** /v1/transformationRules/{id} | Update a transformation rule.

# **CreateRule**
> TransformationRuleResponse CreateRule(ctx, body)
Create a new transformation rule.

Create a new transformation rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransformationRuleRequest**](TransformationRuleRequest.md)| The configuration of the transformation rule to create. | 

### Return type

[**TransformationRuleResponse**](TransformationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRule**
> DeleteRule(ctx, id)
Delete a transformation rule.

Delete a transformation rule with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the transformation rule to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransformationRule**
> TransformationRuleResponse GetTransformationRule(ctx, id)
Get a transformation rule.

Get a transformation rule with the given identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of transformation rule to return. | 

### Return type

[**TransformationRuleResponse**](TransformationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransformationRules**
> TransformationRulesResponse GetTransformationRules(ctx, optional)
Get a list of transformation rules.

Get a list of transformation rules in the organization. The response is paginated with a default limit of 100 rules per page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransformationRuleManagementApiGetTransformationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransformationRuleManagementApiGetTransformationRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of transformation rules returned in the response. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**TransformationRulesResponse**](TransformationRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransformationRule**
> TransformationRuleResponse UpdateTransformationRule(ctx, body, id)
Update a transformation rule.

Update an existing transformation rule. All properties specified in the request are replaced. Missing properties will remain the same.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransformationRuleRequest**](TransformationRuleRequest.md)| Information to update about the transformation rule. | 
  **id** | **string**| Identifier of the transformation rule to update. | 

### Return type

[**TransformationRuleResponse**](TransformationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

