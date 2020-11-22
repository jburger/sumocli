# UserModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**FirstName** | **string** | First name of the user. | [default to null]
**LastName** | **string** | Last name of the user. | [default to null]
**Email** | **string** | Email address of the user. | [default to null]
**RoleIds** | **[]string** | List of roleIds associated with the user. | [default to null]
**Id** | **string** | Unique identifier for the user. | [default to null]
**IsActive** | **bool** | True if the user is active. | [optional] [default to null]
**IsLocked** | **bool** | This has the value &#x60;true&#x60; if the user&#x27;s account has been locked. If a user tries to log into their account several times and fails, his or her account will be locked for security reasons. | [optional] [default to null]
**IsMfaEnabled** | **bool** | True if multi factor authentication is enabled for the user. | [optional] [default to null]
**LastLoginTimestamp** | [**time.Time**](time.Time.md) | Timestamp of the last login for the user in UTC. Will be null if the user has never logged in. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

