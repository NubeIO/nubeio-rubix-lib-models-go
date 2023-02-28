package model

type PollQueueStatistics struct {
	Enable      bool   `json:"enable"`
	MaxPollRate string `json:"max_poll_rate"`

	// References
	FFNetworkUUID string `json:"ff_network_uuid"`
	NetworkName   string `json:"network_name"`
	FFPluginUUID  string `json:"ff_plugin_uuid"`
	PluginName    string `json:"plugin_name"`

	// Statistics
	MaxPollExecuteTime            string `json:"max_poll_execute_time"`             // time for polling to complete (poll response time, doesn't include the time in queue).
	AveragePollExecuteTime        string `json:"average_poll_execute_time"`         // time for polling to complete (poll response time, doesn't include the time in queue).
	MinPollExecuteTime            string `json:"min_poll_execute_time"`             // time for polling to complete (poll response time, doesn't include the time in queue).
	TotalPollQueueLength          int64  `json:"total_poll_queue_length"`           // number of polling points in the current queue.
	TotalStandbyPointsLength      int64  `json:"total_standby_points_length"`       // number of polling points in the standby list.
	TotalPointsOutForPolling      int64  `json:"total_points_out_for_polling"`      // number of points currently out for polling (currently being handled by the protocol plugin).
	ASAPPriorityPollQueueLength   int64  `json:"asap_priority_poll_queue_length"`   // number of ASAP priority polling points in the current queue.
	HighPriorityPollQueueLength   int64  `json:"high_priority_poll_queue_length"`   // number of High priority polling points in the current queue.
	NormalPriorityPollQueueLength int64  `json:"normal_priority_poll_queue_length"` // number of Normal priority polling points in the current queue.
	LowPriorityPollQueueLength    int64  `json:"low_priority_poll_queue_length"`    // number of Low priority polling points in the current queue.
	ASAPPriorityAveragePollTime   string `json:"asap_priority_average_poll_time"`   // average time between ASAP priority polling point added to current queue, and polling complete.
	HighPriorityAveragePollTime   string `json:"high_priority_average_poll_time"`   // average time between High priority polling point added to current queue, and polling complete.
	NormalPriorityAveragePollTime string `json:"normal_priority_average_poll_time"` // average time between Normal priority polling point added to current queue, and polling complete.
	LowPriorityAveragePollTime    string `json:"low_priority_average_poll_time"`    // average time between Low priority polling point added to current queue, and polling complete.
	TotalPollCount                int64  `json:"total_poll_count"`                  // total number of polls completed.
	ASAPPriorityPollCount         int64  `json:"asap_priority_poll_count"`          // total number of ASAP priority polls completed.
	HighPriorityPollCount         int64  `json:"high_priority_poll_count"`          // total number of High priority polls completed.
	NormalPriorityPollCount       int64  `json:"normal_priority_poll_count"`        // total number of Normal priority polls completed.
	LowPriorityPollCount          int64  `json:"low_priority_poll_count"`           // total number of Low priority polls completed.
	ASAPPriorityMaxCycleTime      string `json:"asap_priority_max_cycle_time"`      // threshold setting for triggering a lockup alert for ASAP priority.
	HighPriorityMaxCycleTime      string `json:"high_priority_max_cycle_time"`      // threshold setting for triggering a lockup alert for High priority.
	NormalPriorityMaxCycleTime    string `json:"normal_priority_max_cycle_time"`    // threshold setting for triggering a lockup alert for Normal priority.
	LowPriorityMaxCycleTime       string `json:"low_priority_max_cycle_time"`       // threshold setting for triggering a lockup alert for Low priority.
	ASAPPriorityLockupAlert       bool   `json:"asap_priority_lockup_alert"`        // alert if poll time has exceeded the ASAPPriorityMaxCycleTime
	HighPriorityLockupAlert       bool   `json:"high_priority_lockup_alert"`        // alert if poll time has exceeded the HighPriorityMaxCycleTime
	NormalPriorityLockupAlert     bool   `json:"normal_priority_lockup_alert"`      // alert if poll time has exceeded the NormalPriorityMaxCycleTime
	LowPriorityLockupAlert        bool   `json:"low_priority_lockup_alert"`         // alert if poll time has exceeded the LowPriorityMaxCycleTime
	BusyTime                      string `json:"busy_time"`                         // percent of the time that the plugin is actively polling.
	EnabledTime                   string `json:"enabled_time"`                      // time that the statistics have been running for.
	PortUnavailableTime           string `json:"port_unavailable_time"`             // time that the serial port has been unavailable.
}
