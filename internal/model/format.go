package model

// OutputFormat represents the output mode.
type OutputFormat int

// Output formats.
const (
	TextFormat OutputFormat = iota
	JSONFormat
)

// GetOutputFormat returns an OutputMode instance.
func GetOutputFormat(format string) OutputFormat {
	switch format {
	case "text":
		return TextFormat
	case "json":
		return JSONFormat
	default:
		return TextFormat
	}
}

// String returns the string representation of the output mode.
func (o OutputFormat) String() string {
	switch o {
	case TextFormat:
		return "PlainText"
	case JSONFormat:
		return "JSON"
	default:
		return "Unknown"
	}
}
