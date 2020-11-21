# EmailSearchNotificationSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskType** | **string** | Delivery channel for notifications. | [default to null]
**ToList** | **[]string** | A list of email recipients. | [default to null]
**SubjectTemplate** | **string** | If the notification is scheduled with a threshold, the default subject template will be \&quot;Search Alert: {{AlertCondition}} results found for {{SearchName}}\&quot;. For email notifications without a threshold, the default subject template is \&quot;Search Results: {{SearchName}}\&quot;. | [optional] [default to null]
**IncludeQuery** | **bool** | A boolean value to indicate if the search query should be included in the notification email. | [optional] [default to true]
**IncludeResultSet** | **bool** | A boolean value to indicate if the search result set should be included in the notification email. | [optional] [default to true]
**IncludeHistogram** | **bool** | A boolean value to indicate if the search result histogram should be included in the notification email. | [optional] [default to true]
**IncludeCsvAttachment** | **bool** | A boolean value to indicate if the search results should be included in the notification email as a CSV attachment. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

