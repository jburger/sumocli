# {{classname}}

All URIs are relative to *https://api.au.sumologic.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWhitelistedCidrs**](ServiceWhitelistManagementApi.md#AddWhitelistedCidrs) | **Post** /v1/serviceWhitelist/addresses/add | Whitelist CIDRs/IP addresses.
[**DeleteWhitelistedCidrs**](ServiceWhitelistManagementApi.md#DeleteWhitelistedCidrs) | **Post** /v1/serviceWhitelist/addresses/remove | Remove whitelisted CIDRs/IP addresses.
[**DisableWhitelisting**](ServiceWhitelistManagementApi.md#DisableWhitelisting) | **Post** /v1/serviceWhitelist/disable | Disable service whitelisting.
[**EnableWhitelisting**](ServiceWhitelistManagementApi.md#EnableWhitelisting) | **Post** /v1/serviceWhitelist/enable | Enable service whitelisting.
[**GetWhitelistingStatus**](ServiceWhitelistManagementApi.md#GetWhitelistingStatus) | **Get** /v1/serviceWhitelist/status | Get the whitelisting status.
[**ListWhitelistedCidrs**](ServiceWhitelistManagementApi.md#ListWhitelistedCidrs) | **Get** /v1/serviceWhitelist/addresses | List all whitelisted CIDRs/IP addresses.

# **AddWhitelistedCidrs**
> CidrList AddWhitelistedCidrs(ctx, body)
Whitelist CIDRs/IP addresses.

Add CIDR notations and/or IP addresses to the whitelist of the organization if not already there. When service whitelisting functionality is enabled, CIDRs/IP addresses that are whitelisted will have access to Sumo Logic and/or content sharing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CidrList**](CidrList.md)| List of all CIDR notations and/or IP addresses to be added to the whitelist of the organization. | 

### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWhitelistedCidrs**
> CidrList DeleteWhitelistedCidrs(ctx, body)
Remove whitelisted CIDRs/IP addresses.

Remove whitelisted CIDR notations and/or IP addresses from the organization. Removed CIDRs/IPs will immediately lose access to Sumo Logic and content sharing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CidrList**](CidrList.md)| List of all CIDR notations and/or IP addresses to be removed from the whitelist of the organization. | 

### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableWhitelisting**
> DisableWhitelisting(ctx, whitelistType)
Disable service whitelisting.

Disable service whitelisting functionality for login/API authentication or content sharing for the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whitelistType** | **string**| The type of whitelisting to be disabled. It can be one of: &#x60;Login&#x60;, &#x60;Content&#x60;, or &#x60;Both&#x60;. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableWhitelisting**
> EnableWhitelisting(ctx, whitelistType)
Enable service whitelisting.

Enable service whitelisting functionality for the organization. The service whitelisting can be for 1. Login: If enabled, access to Sumo Logic is granted only to CIDRs/IP addresses that are whitelisted. 2. Content: If enabled, dashboards can be shared with users connecting from CIDRs/IP addresses that are whitelisted without logging in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whitelistType** | **string**| The type of whitelisting to be enabled. It can be one of: &#x60;Login&#x60;, &#x60;Content&#x60;, or &#x60;Both&#x60;. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWhitelistingStatus**
> AllowlistingStatus GetWhitelistingStatus(ctx, )
Get the whitelisting status.

Get the status of the service whitelisting functionality for login/API authentication or content sharing for the organization.

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

# **ListWhitelistedCidrs**
> CidrList ListWhitelistedCidrs(ctx, )
List all whitelisted CIDRs/IP addresses.

Get a list of all whitelisted CIDR notations and/or IP addresses for the organization.

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

