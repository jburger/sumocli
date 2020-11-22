# CreateRoleDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the role. | [default to null]
**Description** | **string** | Description of the role. | [optional] [default to null]
**FilterPredicate** | **string** | A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using &#x27;!_sourceCategory&#x3D;billing&#x27; as a filter predicate will prevent users assigned to the role from viewing logs from the source category named &#x27;billing&#x27;. | [optional] [default to null]
**Users** | **[]string** | List of user identifiers to assign the role to. | [optional] [default to null]
**Capabilities** | **[]string** | List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - dataVolumeIndex   - viewConnections   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - manageDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - manageMonitors   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipWhitelisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex | [optional] [default to null]
**AutofillDependencies** | **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

