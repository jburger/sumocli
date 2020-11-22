# Email

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Connection type of the connection. Valid values:   1.  &#x60;Email&#x60;   2.  &#x60;AWSLambda&#x60;   3.  &#x60;AzureFunctions&#x60;   4.  &#x60;Datadog&#x60;   5.  &#x60;HipChat&#x60;   6.  &#x60;Jira&#x60;   7.  &#x60;NewRelic&#x60;   8. &#x60;Opsgenie&#x60;   8. &#x60;PagerDuty&#x60;   10. &#x60;Slack&#x60;   11. &#x60;Webhook&#x60; | [default to null]
**Recipients** | **[]string** | A list of email addresses to send to when the rule fires. | [default to null]
**Subject** | **string** | The subject line of the email. | [default to null]
**MessageBody** | **string** | The message body of the email to send. | [optional] [default to null]
**TimeZone** | **string** | Time zone for the email content. All dates/times will be displayed in this timeZone in the email payload. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

