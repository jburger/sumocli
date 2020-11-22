# Content

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [default to null]
**CreatedBy** | **string** | Identifier of the user who created the resource. | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Last modification timestamp in UTC. | [default to null]
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | [default to null]
**Id** | **string** | Identifier of the content item. | [default to null]
**Name** | **string** | The name of the content item. | [default to null]
**ItemType** | **string** | Type of the content item. Supported values are:   1. Folder   2. Search   3. Dashboard | [default to null]
**ParentId** | **string** | Identifier of the parent content item. | [default to null]
**Permissions** | **[]string** | List of permissions the user has on the content item. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

