package datatype

type QOS byte

const (
	// AtMostOnce means the broker will deliver at most once
	AtMostOnce QOS = iota
	// AtLeastOnce means the broker will deliver c message at least once
	AtLeastOnce
	// ExactlyOnce means the broker will deliver c message exactly once
	ExactlyOnce
)
