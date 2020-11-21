# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllHealthEvents**](HealthEventsApi.md#ListAllHealthEvents) | **Get** /v1/healthEvents | Get a list of health events.
[**ListAllHealthEventsForResources**](HealthEventsApi.md#ListAllHealthEventsForResources) | **Post** /v1/healthEvents/resources | Health events for specific resources.

# **ListAllHealthEvents**
> ListHealthEventResponse ListAllHealthEvents(ctx, optional)
Get a list of health events.

Get a list of all the unresolved health events in your account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HealthEventsApiListAllHealthEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HealthEventsApiListAllHealthEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit the number of health events returned in the response. The number of health events returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.String**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListHealthEventResponse**](ListHealthEventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllHealthEventsForResources**
> ListHealthEventResponse ListAllHealthEventsForResources(ctx, body, optional)
Health events for specific resources.

Get a list of all the unresolved events in your account that belong to the supplied resource identifiers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceIdentities**](ResourceIdentities.md)| Resource identifiers to request health events from. | 
 **optional** | ***HealthEventsApiListAllHealthEventsForResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HealthEventsApiListAllHealthEventsForResourcesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.**| Limit the number of health events returned in the response. The number of health events returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **optional.**| Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListHealthEventResponse**](ListHealthEventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

