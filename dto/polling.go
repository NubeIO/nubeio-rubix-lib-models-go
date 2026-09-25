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
	PollDurationFast    string `json:"poll_duration_fast"`    // p01 of poll duration
	PollDurationNormal  string `json:"poll_duration_normal"`  // p50 of poll duration
	PollDurationSlow    string `json:"poll_duration_slow"`    // p99 of poll duration
	QueueNormalWait     string `json:"queue_normal_wait"`     // p50 of time spent waiting in the queue
	QueueSlowWait       string `json:"queue_slow_wait"`       // p99 of time spent waiting in the queue
	ScheduleDelayNormal string `json:"schedule_delay_normal"` // p50 of delay past the scheduled poll time
	ScheduleDelayLong   string `json:"schedule_delay_long"`   // p99 of delay past the scheduled poll time

	EnabledTime string `json:"enabled_time"`
}
