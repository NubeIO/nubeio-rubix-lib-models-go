package nargs

import "encoding/json"

type Args struct {
	HostUUID                       *string   `json:"host_uuid,omitempty"`
	GlobalUUID                     *string   `json:"global_uuid,omitempty"`
	UUID                           *string   `json:"uuid,omitempty"`
	NetworkUUID                    *string   `json:"network_uuid,omitempty"`
	DeviceUUID                     *string   `json:"device_uuid,omitempty"`
	PointUUID                      *string   `json:"point_uuid,omitempty"`
	PointSourceUUID                *string   `json:"point_source_uuid,omitempty"`
	SourceUUID                     *string   `json:"source_uuid,omitempty"`
	DeviceId                       *string   `json:"device_id,omitempty"`
	AddressUUID                    *string   `json:"address_uuid,omitempty"`
	AddressID                      *string   `json:"address_id,omitempty"`
	ObjectType                     *string   `json:"object_type,omitempty"`
	IoNumber                       *string   `json:"io_number,omitempty"`
	MemberUUID                     *string   `json:"member_uuid,omitempty"`
	Name                           *string   `json:"name,omitempty"`
	SearchKeyword                  *string   `json:"search_keyword,omitempty"`
	Tag                            *string   `json:"tag,omitempty"`
	MetaTag                        *string   `json:"meta_tag,omitempty"`
	MetaTags                       *string   `json:"meta_tags,omitempty"`
	Statuses                       *[]string `json:"statuses,omitempty"`
	Target                         *string   `json:"target,omitempty"`
	WriteValue                     *bool     `json:"write_value,omitempty"`
	HistoryEnabled                 *bool     `json:"history_enabled,omitempty"`
	ByPluginName                   bool      `json:"by_plugin_name,omitempty"`
	Notified                       *bool     `json:"notified,omitempty"`
	ShowCloneNetworks              bool      `json:"show_clone_networks,omitempty"`
	WidgetOrder                    *string   `json:"widget_order,omitempty"`
	Order                          *string   `json:"order,omitempty"`
	Offset                         *int      `json:"offset,omitempty"`
	Limit                          *int      `json:"limit,omitempty"`
	Count                          *int      `json:"count,omitempty"`
	Sort                           *string   `json:"sort,omitempty"`
	Id                             *string   `json:"id,omitempty"`
	IdGt                           *string   `json:"id_gt,omitempty"`
	Timestamp                      *string   `json:"timestamp,omitempty"`
	TimestampGt                    *string   `json:"timestamp_gt,omitempty"`
	TimestampLt                    *string   `json:"timestamp_lt,omitempty"`
	TimestampGte                   *string   `json:"timestamp_gte,omitempty"`
	TimestampLte                   *string   `json:"timestamp_lte,omitempty"`
	StartDatetime                  *string   `json:"start_datetime,omitempty"`
	EndDatetime                    *string   `json:"end_datetime,omitempty"`
	WithGroups                     bool      `json:"with_groups,omitempty"`
	WithHosts                      bool      `json:"with_hosts,omitempty"`
	WithNetworks                   bool      `json:"with_networks,omitempty"`
	WithDevices                    bool      `json:"with_devices,omitempty"`
	WithPoints                     bool      `json:"with_points,omitempty"`
	WithPriority                   bool      `json:"with_priority,omitempty"`
	WithAlerts                     bool      `json:"with_alerts,omitempty"`
	WithTransactions               bool      `json:"with_transactions,omitempty"`
	WithMembers                    bool      `json:"with_members,omitempty"`
	WithMemberDevices              bool      `json:"with_member_devices,omitempty"`
	WithTeams                      bool      `json:"with_teams,omitempty"`
	WithViews                      bool      `json:"with_views,omitempty"`
	WithComments                   bool      `json:"with_comments,omitempty"`
	WithWidgets                    bool      `json:"with_widgets,omitempty"`
	WithViewTemplateWidgets        bool      `json:"with_view_template_widgets,omitempty"`
	WithViewTemplateWidgetPointers bool      `json:"with_view_template_widget_pointers,omitempty"`
	WithTickets                    bool      `json:"with_tickets,omitempty"`
	WithTags                       bool      `json:"with_tags,omitempty"`
	WithMetaTags                   bool      `json:"with_meta_tags,omitempty"`
	Aggregate                      *string   `json:"aggregate,omitempty"`
	GroupLimit                     *int      `json:"group_limit,omitempty"`
	PointHistoryId                 *string   `json:"point_history_id,omitempty"`
	PointHistoryTimestamp          *string   `json:"point_history_timestamp,omitempty"`
	MetricLogId                    *string   `json:"metric_log_id,omitempty"`
	MetricLogTimestamp             *string   `json:"metric_log_timestamp,omitempty"`
	WithAlertConditions            bool      `json:"with_alert_conditions,omitempty"`
	EntityUUID                     *string   `json:"entity_uuid,omitempty"`
	CreatedAt                      *string   `json:"created_at"`
	CreatedAtGt                    *string   `json:"created_at_gt"`
	CreatedAtGte                   *string   `json:"created_at_gte"`
	CreatedAtLt                    *string   `json:"created_at_lt"`
	CreatedAtLte                   *string   `json:"created_at_lte"`
	LastUpdated                    *string   `json:"last_updated"`
	LastUpdatedGt                  *string   `json:"last_updated_gt"`
	LastUpdatedGte                 *string   `json:"last_updated_gte"`
	LastUpdatedLt                  *string   `json:"last_updated_lt"`
	LastUpdatedLte                 *string   `json:"last_updated_lte"`
	Title                          *string   `json:"title"`
	ShowClones                     bool      `json:"show_clones,omitempty"`
	Source                         *string   `json:"source,omitempty"`
	PointState                     *string   `json:"point_state,omitempty"`
	HostUUIDs                      *[]string `json:"host_uuids"`
	Sources                        *[]string `json:"sources"`
	EntityTypes                    *[]string `json:"entity_types,omitempty"`
	Types                          *[]string `json:"types,omitempty"`
	Severities                     *[]string `json:"severities,omitempty"`
}

const (
	HostUUID                       = "host_uuid"
	GlobalUUID                     = "global_uuid"
	UUID                           = "uuid"
	NetworkUUID                    = "network_uuid"
	DeviceUUID                     = "device_uuid"
	PointUUID                      = "point_uuid"
	PointSourceUUID                = "point_source_uuid"
	SourceUUID                     = "source_uuid"
	DeviceId                       = "device_id"
	AddressUUID                    = "address_uuid"
	AddressID                      = "address_id"
	ObjectType                     = "object_type"
	IoNumber                       = "io_number"
	Name                           = "name"
	NetworkName                    = "network_name"
	DeviceName                     = "device_name"
	PointName                      = "point_name"
	SearchKeyword                  = "search_keyword"
	Tag                            = "tag"
	MetaTag                        = "meta_tag"
	MetaTags                       = "meta_tags"
	Status                         = "status"
	HistoryEnabled                 = "history_enabled"
	Target                         = "target"
	ByPluginName                   = "by_plugin_name"
	Notified                       = "notified"
	ShowCloneNetworks              = "show_clone_networks"
	WidgetOrder                    = "widget_order"
	Order                          = "order"
	Offset                         = "offset"
	Limit                          = "limit"
	Count                          = "count"
	Sort                           = "sort"
	Id                             = "id"
	IdGt                           = "id_gt"
	Timestamp                      = "timestamp"
	TimestampGt                    = "timestamp_gt"
	TimestampLt                    = "timestamp_lt"
	TimestampGte                   = "timestamp_gte"
	TimestampLte                   = "timestamp_lte"
	StartDatetime                  = "start_datetime"
	EndDatetime                    = "end_datetime"
	WithGroups                     = "with_groups"
	WithHosts                      = "with_hosts"
	WithNetworks                   = "with_networks"
	WithDevices                    = "with_devices"
	WithPoints                     = "with_points"
	WithPriority                   = "with_priority"
	WithAlerts                     = "with_alerts"
	WithTransactions               = "with_transactions"
	WithMembers                    = "with_members"
	WithMemberDevices              = "with_member_devices"
	WithTeams                      = "with_teams"
	WithViews                      = "with_views"
	WithComments                   = "with_comments"
	WithWidgets                    = "with_widgets"
	WithViewTemplateWidgets        = "with_view_template_widgets"
	WithViewTemplateWidgetPointers = "with_view_template_widget_pointers"
	WithTickets                    = "with_tickets"
	WithTags                       = "with_tags"
	WithMetaTags                   = "with_meta_tags"
	GroupLimit                     = "group_limit"
	PointHistoryId                 = "point_history_id"
	PointHistoryTimestamp          = "point_history_timestamp"
	MetricLogId                    = "metric_log_id"
	MetricLogTimestamp             = "metric_log_timestamp"
	WithAlertConditions            = "with_alert_conditions"
	EntityUUID                     = "entity_uuid"
	CreatedAt                      = "created_at"
	CreatedAtGt                    = "created_at_gt"
	CreatedAtGte                   = "created_at_gte"
	CreatedAtLt                    = "created_at_lt"
	CreatedAtLte                   = "created_at_lte"
	LastUpdated                    = "last_updated"
	LastUpdatedGt                  = "last_updated_gt"
	LastUpdatedGte                 = "last_updated_gte"
	LastUpdatedLt                  = "last_updated_lt"
	LastUpdatedLte                 = "last_updated_lte"
	Title                          = "title"
	ShowClones                     = "show_clones"
	Source                         = "source"
	PointState                     = "point_state"
	EntityType                     = "entity_type"
	Type                           = "type"
	Severity                       = "severity"
)

const (
	DefaultByPluginName                   = "false"
	DefaultShowCloneNetworks              = "false"
	DefaultWithHosts                      = "false"
	DefaultWithGroups                     = "false"
	DefaultWithNetworks                   = "false"
	DefaultWithDevices                    = "false"
	DefaultWithPoints                     = "false"
	DefaultWithPriority                   = "false"
	DefaultWithAlerts                     = "false"
	DefaultWithTransactions               = "false"
	DefaultWithTags                       = "false"
	DefaultWithMetaTags                   = "false"
	DefaultWithMembers                    = "false"
	DefaultWithMemberDevices              = "false"
	DefaultWithTeams                      = "false"
	DefaultWithViews                      = "false"
	DefaultWithComments                   = "false"
	DefaultWithWidgets                    = "false"
	DefaultWithViewTemplateWidgets        = "false"
	DefaultWithViewTemplateWidgetPointers = "false"
	DefaultWithTickets                    = "false"
	DefaultWithAlertConditions            = "false"
	DefaultShowClones                     = "false"
)

func SerializeArgs(args Args) (*string, error) {
	argsData, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	argsString := string(argsData)
	return &argsString, nil
}

func DeserializeArgs(args string) (*Args, error) {
	deserializedArgs := Args{}
	if len(args) == 0 {
		return &deserializedArgs, nil
	}
	err := json.Unmarshal([]byte(args), &deserializedArgs)
	if err != nil {
		return nil, err
	}
	return &deserializedArgs, nil
}
