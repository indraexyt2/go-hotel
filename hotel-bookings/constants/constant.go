package constants

var MidtransStatusMapping = map[string][]string{
	"pending":        {"settlement", "cancel", "expire", "failure", "success"},
	"settlement":     {"completed"},
	"capture":        {"settlement", "refund"},
	"deny":           {},
	"cancel":         {},
	"expire":         {},
	"failure":        {},
	"refund":         {"partial_refund"},
	"partial_refund": {"refund"},
	"chargeback":     {},
}
