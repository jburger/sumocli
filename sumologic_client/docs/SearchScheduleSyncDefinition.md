# SearchScheduleSyncDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | **string** | Cron-like expression specifying the search&#x27;s schedule. Field scheduleType must be set to \&quot;Custom\&quot;, otherwise, scheduleType takes precedence over cronExpression. | [optional] [default to null]
**DisplayableTimeRange** | **string** | A human-friendly text describing the query time range. For e.g. \&quot;-2h\&quot;, \&quot;last three days\&quot;, \&quot;team default time\&quot; | [optional] [default to null]
**ParseableTimeRange** | [***ResolvableTimeRange**](ResolvableTimeRange.md) |  | [default to null]
**TimeZone** | **string** | Time zone identifier for time specification. Either an abbreviation such as \&quot;PST\&quot;, a full name such as \&quot;America/Los_Angeles\&quot;, or a custom ID such as \&quot;GMT-8:00\&quot;. Note that the support of abbreviations is for JDK 1.1.x compatibility only and full names should be used. | [default to null]
**Threshold** | [***NotificationThresholdSyncDefinition**](NotificationThresholdSyncDefinition.md) |  | [optional] [default to null]
**Notification** | [***ScheduleNotificationSyncDefinition**](ScheduleNotificationSyncDefinition.md) |  | [default to null]
**ScheduleType** | **string** | Run schedule of the scheduled search. Set to \&quot;Custom\&quot; to specify the schedule with a CRON expression. Possible schedule types are:   - &#x60;RealTime&#x60;   - &#x60;15Minutes&#x60;   - &#x60;1Hour&#x60;   - &#x60;2Hours&#x60;   - &#x60;4Hours&#x60;   - &#x60;6Hours&#x60;   - &#x60;8Hours&#x60;   - &#x60;12Hours&#x60;   - &#x60;1Day&#x60;   - &#x60;1Week&#x60;   - &#x60;Custom&#x60; | [default to null]
**MuteErrorEmails** | **bool** | If enabled, emails are not sent out in case of errors with the search. | [optional] [default to null]
**Parameters** | [**[]ScheduleSearchParameterSyncDefinition**](ScheduleSearchParameterSyncDefinition.md) | A list of scheduled search parameters. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

