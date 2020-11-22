# SavedSearchSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryText** | **string** | The text of a Sumo Logic query. | [default to null]
**DefaultTimeRange** | **string** | Default time range for the search. Possible types of time ranges are:   - relative time range: e.g. \&quot;-1d -12h\&quot; represents a time range from one day ago to 12 hours ago.   - absolute time range: e.g. \&quot;01-04-2017 20:32:00 to 01-04-2017 20:35:00\&quot; represents a time range     from April 1st, 2017 at 8:32 PM until April 1st, 2017 at 8:35 PM. | [default to null]
**ByReceiptTime** | **bool** | Set it to true to run the search using receipt time. By default, searches do not run by receipt time. | [default to false]
**ViewName** | **string** | The name of the Scheduled View that has indexed the data you want to search. | [optional] [default to null]
**ViewStartTime** | [**time.Time**](time.Time.md) | Start timestamp of the Scheduled View in UTC format. | [optional] [default to null]
**QueryParameters** | [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | An array of search query parameter objects. | [default to null]
**ParsingMode** | **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to Manual]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

