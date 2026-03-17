package dto

type PollQueueStatistics struct {
	Enable bool `json:"enable"`

	// References
	FFNetworkUUID string `json:"ff_network_uuid"`
	NetworkName   string `json:"network_name"`
	FFPluginUUID  string `json:"ff_plugin_uuid"`
	PluginName    string `json:"plugin_name"`

	// Queue state
	TotalStandbyPointsLength int64 `json:"total_standby_points_length"`

	// Poll rate counts (attempts)
	FastPollCount    int64 `json:"fast_poll_count"`
	NormalPollCount  int64 `json:"normal_poll_count"`
	SlowPollCount    int64 `json:"slow_poll_count"`
	TotalPollCount   int64 `json:"total_poll_count"`
	SuccessPollCount int64 `json:"success_poll_count"`
	ErrorPollCount   int64 `json:"error_poll_count"`

	// Execution and scheduler percentiles
	PollDurationP01 string `json:"poll_duration_p01"`
	PollDurationP50 string `json:"poll_duration_p50"`
	PollDurationP99 string `json:"poll_duration_p99"`
	QueueWaitP50    string `json:"queue_wait_p50"`
	QueueWaitP99    string `json:"queue_wait_p99"`
	ScheduleLagP50  string `json:"schedule_lag_p50"`
	ScheduleLagP99  string `json:"schedule_lag_p99"`

	EnabledTime string `json:"enabled_time"`
}
