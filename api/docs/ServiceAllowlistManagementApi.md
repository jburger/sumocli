# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAllowlistedCidrs**](ServiceAllowlistManagementApi.md#AddAllowlistedCidrs) | **Post** /v1/serviceAllowlist/addresses/add | Allowlist CIDRs/IP addresses.
[**DeleteAllowlistedCidrs**](ServiceAllowlistManagementApi.md#DeleteAllowlistedCidrs) | **Post** /v1/serviceAllowlist/addresses/remove | Remove allowlisted CIDRs/IP addresses.
[**DisableAllowlisting**](ServiceAllowlistManagementApi.md#DisableAllowlisting) | **Post** /v1/serviceAllowlist/disable | Disable service allowlisting.
[**EnableAllowlisting**](ServiceAllowlistManagementApi.md#EnableAllowlisting) | **Post** /v1/serviceAllowlist/enable | Enable service allowlisting.
[**GetAllowlistingStatus**](ServiceAllowlistManagementApi.md#GetAllowlistingStatus) | **Get** /v1/serviceAllowlist/status | Get the allowlisting status.
[**ListAllowlistedCidrs**](ServiceAllowlistManagementApi.md#ListAllowlistedCidrs) | **Get** /v1/serviceAllowlist/addresses | List all allowlisted CIDRs/IP addresses.

# **AddAllowlistedCidrs**
> CidrList AddAllowlistedCidrs(ctx, body)
Allowlist CIDRs/IP addresses.

Add CIDR notations and/or IP addresses to the allowlist of the organization if not already there. When service allowlisting functionality is enabled, CIDRs/IP addresses that are allowlisted will have access to Sumo Logic and/or content sharing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CidrList**](CidrList.md)| List of all CIDR notations and/or IP addresses to be added to the allowlist of the organization. | 

### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllowlistedCidrs**
> CidrList DeleteAllowlistedCidrs(ctx, body)
Remove allowlisted CIDRs/IP addresses.

Remove allowlisted CIDR notations and/or IP addresses from the organization. Removed CIDRs/IPs will immediately lose access to Sumo Logic and content sharing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CidrList**](CidrList.md)| List of all CIDR notations and/or IP addresses to be removed from the allowlist of the organization. | 

### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableAllowlisting**
> DisableAllowlisting(ctx, allowlistType)
Disable service allowlisting.

Disable service allowlisting functionality for login/API authentication or content sharing for the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allowlistType** | **string**| The type of allowlisting to be disabled. It can be one of: &#x60;Login&#x60;, &#x60;Content&#x60;, or &#x60;Both&#x60;. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAllowlisting**
> EnableAllowlisting(ctx, allowlistType)
Enable service allowlisting.

Enable service allowlisting functionality for the organization. The service allowlisting can be for 1. Login: If enabled, access to Sumo Logic is granted only to CIDRs/IP addresses that are allowlisted. 2. Content: If enabled, dashboards can be shared with users connecting from CIDRs/IP addresses that are allowlisted without logging in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allowlistType** | **string**| The type of allowlisting to be enabled. It can be one of: &#x60;Login&#x60;, &#x60;Content&#x60;, or &#x60;Both&#x60;. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllowlistingStatus**
> AllowlistingStatus GetAllowlistingStatus(ctx, )
Get the allowlisting status.

Get the status of the service allowlisting functionality for login/API authentication or content sharing for the organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AllowlistingStatus**](AllowlistingStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllowlistedCidrs**
> CidrList ListAllowlistedCidrs(ctx, )
List all allowlisted CIDRs/IP addresses.

Get a list of all allowlisted CIDR notations and/or IP addresses for the organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

