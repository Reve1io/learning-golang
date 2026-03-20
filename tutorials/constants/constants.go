package constants

const (
	OK = iota
	CANCELED
	UNKNOWN
)

func ErrorMessageToCode(msg string) int {

	switch msg {
	case "OK":
		return OK
	case "CANCELLED":
		return CANCELED
	case "UNKNOWN":
		return UNKNOWN
	default:
		return UNKNOWN
	}
}
