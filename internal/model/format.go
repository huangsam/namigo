package model

// FormatOption represents the output mode.
type FormatOption int

const (
	TextOption FormatOption = iota
	JSONOption

	TextValue = "text"
	JSONValue = "json"
)

// GetOutputFormat returns an OutputMode instance.
func GetOutputFormat(format string) FormatOption {
	switch format {
	case TextValue:
		return TextOption
	case JSONValue:
		return JSONOption
	default:
		return TextOption
	}
}

// String returns the string representation of the output mode.
func (o FormatOption) String() string {
	switch o {
	case TextOption:
		return "PlainText"
	case JSONOption:
		return "JSON"
	default:
		return "Unknown"
	}
}
