# LogSearchQueryTimeRangeBase

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | Query to perform. | [default to null]
**TimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [default to null]
**RunByReceiptTime** | **bool** | This has the value &#x60;true&#x60; if the search is to be run by receipt time and &#x60;false&#x60; if it is to be run by message time. | [optional] [default to false]
**ParsingMode** | **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to Manual]
**QueryParameters** | [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | Definition of the query parameters. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

