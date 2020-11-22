# Folder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the content item. | [default to null]
**Name** | **string** | The name of the content item. | [default to null]
**ItemType** | **string** | Type of the content item. Supported values are:   1. Folder   2. Search   3. Dashboard | [default to null]
**ParentId** | **string** | Identifier of the parent content item. | [default to null]
**Permissions** | **[]string** | List of permissions the user has on the content item. | [default to null]
**Description** | **string** | The description of the folder. | [optional] [default to null]
**Children** | [**[]Content**](Content.md) | A list of the content items. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

