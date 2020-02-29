package anime

import "strings"

type Type int

// anime type constants for Type
const (
	TV    = 0
	MOVIE = 1
	ONA   = 2
	OVA   = 3
)

func (Type) Equal(a Type, b Type) bool {
	return a == b
}

func (Type) ToString(t Type) string {
	switch t {
	case TV:
		return "TV Series"

	case MOVIE:
		return "Movie"

	case ONA:
		return "ONA"

	case OVA:
		return "OVA"

	default:
		return "Unknown"
	}
}

func AsType(val string) Type {
	val = strings.ToLower(val)
	switch val {
	case "tv":
	case "tv series":
		return TV

	case "movie":
		return MOVIE

	case "ONA":
		return ONA

	case "OVA":
		return OVA

	default:
		return UNKNOWN
	}

	return UNKNOWN
}
