package nargs

type Args struct {
	HostUUID                       *string
	GlobalUUID                     *string
	UUID                           *string
	NetworkUUID                    *string
	DeviceUUID                     *string
	PointUUID                      *string
	PointSourceUUID                *string
	SourceUUID                     *string
	DeviceId                       *string
	AddressUUID                    *string
	AddressID                      *string
	ObjectType                     *string
	IoNumber                       *string
	MemberUUID                     *string
	Name                           *string
	SearchKeyword                  *string
	Tag                            *string
	MetaTag                        *string
	MetaTags                       *string
	Statuses                       *[]string
	Target                         *string
	ByPluginName                   bool
	Notified                       *bool
	ShowCloneNetworks              bool
	WidgetOrder                    *string
	Order                          *string
	Offset                         *int
	Limit                          *int
	IdGt                           *string
	TimestampGt                    *string
	TimestampLt                    *string
	StartDatetime                  *string
	EndDatetime                    *string
	WithGroups                     bool
	WithHosts                      bool
	WithNetworks                   bool
	WithDevices                    bool
	WithPoints                     bool
	WithPriority                   bool
	WithAlerts                     bool
	WithTransactions               bool
	WithMembers                    bool
	WithMemberDevices              bool
	WithTeams                      bool
	WithViews                      bool
	WithComments                   bool
	WithWidgets                    bool
	WithViewTemplateWidgets        bool
	WithViewTemplateWidgetPointers bool
	WithTickets                    bool
	WithTags                       bool
	WithMetaTags                   bool
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
	Target                         = "target"
	ByPluginName                   = "by_plugin_name"
	Notified                       = "notified"
	ShowCloneNetworks              = "show_clone_networks"
	WidgetOrder                    = "widget_order"
	Order                          = "order"
	Offset                         = "offset"
	Limit                          = "limit"
	IdGt                           = "id_gt"
	TimestampGt                    = "timestamp_gt"
	TimestampLt                    = "timestamp_lt"
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
)
