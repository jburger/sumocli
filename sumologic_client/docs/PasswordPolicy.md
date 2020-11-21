# PasswordPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLength** | **int32** | The minimum length of the password. | [default to 8]
**MaxLength** | **int32** | The maximum length of the password. | [default to 128]
**MustContainLowercase** | **bool** | If the password must contain lower case characters. | [default to true]
**MustContainUppercase** | **bool** | If the password must contain upper case characters. | [default to true]
**MustContainDigits** | **bool** | If the password must contain digits. | [default to true]
**MustContainSpecialChars** | **bool** | If the password must contain special characters. | [default to true]
**MaxPasswordAgeInDays** | **int32** | Maximum number of days that a password can be used before user is required to change it. Put -1 if the user should not have to change their password. | [default to 365]
**MinUniquePasswords** | **int32** | The minimum number of unique new passwords that a user must use before an old password can be reused. | [default to 10]
**AccountLockoutThreshold** | **int32** | Number of failed login attempts allowed before account is locked-out. | [default to 6]
**FailedLoginResetDurationInMins** | **int32** | The duration of time in minutes that must elapse from the first failed login attempt after which failed login count is reset to 0. | [default to 10]
**AccountLockoutDurationInMins** | **int32** | The duration of time in minutes that a locked-out account remained locked before getting unlocked automatically. | [default to 30]
**RequireMfa** | **bool** | If MFA should be required to log in. By default, this field is set to &#x60;false&#x60;. | [default to false]
**RememberMfa** | **bool** | If MFA should be remembered on the browser. | [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

