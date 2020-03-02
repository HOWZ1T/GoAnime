package types

import (
	"strings"
)

type Status int

// status constants for Status
const (
	UNKNOWN   = -1
	ON_GOING  = 0
	COMPLETED = 1
)

func (Status) Equal(a Status, b Status) bool {
	return a == b
}

func (s Status) ToString() string {
	switch s {
	case ON_GOING:
		return "Ongoing"

	case COMPLETED:
		return "Completed"

	default:
		return "Unknown"
	}
}

func AsStatus(val string) Status {
	val = strings.ToLower(val)
	switch val {
	case "completed":
		return COMPLETED

	case "on going":
	case "ongoing":
		return ON_GOING

	default:
		return UNKNOWN
	}

	return UNKNOWN
}
