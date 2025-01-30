package constants

var UpdateStatusMapping = map[string][]string{
	"pending":     {"active", "canceled"},
	"active":      {"rescheduled", "refunded", "completed"},
	"canceled":    {},
	"completed":   {},
	"refunded":    {},
	"rescheduled": {"active", "refunded"},
}
