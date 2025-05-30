# FlexV1Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Configuration resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Configuration resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Configuration resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Attributes** | Pointer to **map[string]interface{}** | An object that contains application-specific data. |
**Status** | Pointer to [**string**](ConfigurationEnumStatus.md) |  |
**TaskrouterWorkspaceSid** | Pointer to **string** | The SID of the TaskRouter Workspace. |
**TaskrouterTargetWorkflowSid** | Pointer to **string** | The SID of the TaskRouter target Workflow. |
**TaskrouterTargetTaskqueueSid** | Pointer to **string** | The SID of the TaskRouter Target TaskQueue. |
**TaskrouterTaskqueues** | Pointer to **[]map[string]interface{}** | The list of TaskRouter TaskQueues. |
**TaskrouterSkills** | Pointer to **[]map[string]interface{}** | The Skill description for TaskRouter workers. |
**TaskrouterWorkerChannels** | Pointer to **map[string]interface{}** | The TaskRouter default channel capacities and availability for workers. |
**TaskrouterWorkerAttributes** | Pointer to **map[string]interface{}** | The TaskRouter Worker attributes. |
**TaskrouterOfflineActivitySid** | Pointer to **string** | The TaskRouter SID of the offline activity. |
**RuntimeDomain** | Pointer to **string** | The URL where the Flex instance is hosted. |
**MessagingServiceInstanceSid** | Pointer to **string** | The SID of the Messaging service instance. |
**ChatServiceInstanceSid** | Pointer to **string** | The SID of the chat service this user belongs to. |
**FlexServiceInstanceSid** | Pointer to **string** | The SID of the Flex service instance. |
**FlexInstanceSid** | Pointer to **string** | The SID of the Flex instance. |
**UiLanguage** | Pointer to **string** | The primary language of the Flex UI. |
**UiAttributes** | Pointer to **map[string]interface{}** | The object that describes Flex UI characteristics and settings. |
**UiDependencies** | Pointer to **map[string]interface{}** | The object that defines the NPM packages and versions to be used in Hosted Flex. |
**UiVersion** | Pointer to **string** | The Pinned UI version. |
**ServiceVersion** | Pointer to **string** | The Flex Service version. |
**CallRecordingEnabled** | Pointer to **bool** | Whether call recording is enabled. |
**CallRecordingWebhookUrl** | Pointer to **string** | The call recording webhook URL. |
**CrmEnabled** | Pointer to **bool** | Whether CRM is present for Flex. |
**CrmType** | Pointer to **string** | The CRM type. |
**CrmCallbackUrl** | Pointer to **string** | The CRM Callback URL. |
**CrmFallbackUrl** | Pointer to **string** | The CRM Fallback URL. |
**CrmAttributes** | Pointer to **map[string]interface{}** | An object that contains the CRM attributes. |
**PublicAttributes** | Pointer to **map[string]interface{}** | The list of public attributes, which are visible to unauthenticated clients. |
**PluginServiceEnabled** | Pointer to **bool** | Whether the plugin service enabled. |
**PluginServiceAttributes** | Pointer to **map[string]interface{}** | The plugin service attributes. |
**Integrations** | Pointer to **[]map[string]interface{}** | A list of objects that contain the configurations for the Integrations supported in this configuration. |
**OutboundCallFlows** | Pointer to **map[string]interface{}** | The list of outbound call flows. |
**ServerlessServiceSids** | Pointer to **[]string** | The list of serverless service SIDs. |
**QueueStatsConfiguration** | Pointer to **map[string]interface{}** | Configurable parameters for Queues Statistics. |
**Notifications** | Pointer to **map[string]interface{}** | Configurable parameters for Notifications. |
**Markdown** | Pointer to **map[string]interface{}** | Configurable parameters for Markdown. |
**Url** | Pointer to **string** | The absolute URL of the Configuration resource. |
**FlexInsightsHr** | Pointer to **map[string]interface{}** | Object with enabled/disabled flag with list of workspaces. |
**FlexInsightsDrilldown** | Pointer to **bool** | Setting this to true will redirect Flex UI to the URL set in flex_url |
**FlexUrl** | Pointer to **string** | URL to redirect to in case drilldown is enabled. |
**ChannelConfigs** | Pointer to **[]map[string]interface{}** | Settings for different limits for Flex Conversations channels attachments. |
**DebuggerIntegration** | Pointer to **map[string]interface{}** | Configurable parameters for Debugger Integration. |
**FlexUiStatusReport** | Pointer to **map[string]interface{}** | Configurable parameters for Flex UI Status report. |
**AgentConvEndMethods** | Pointer to **map[string]interface{}** | Agent conversation end methods. |
**CitrixVoiceVdi** | Pointer to **map[string]interface{}** | Citrix voice vdi configuration and settings. |
**OfflineConfig** | Pointer to **map[string]interface{}** | Presence and presence ttl configuration |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


