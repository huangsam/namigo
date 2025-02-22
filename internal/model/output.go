package model

// OutputMode represents the output mode.
type OutputMode int

// Output modes.
const (
	TextMode OutputMode = iota
	JSONMode
)

// GetOutputMode returns an OutputMode instance.
func GetOutputMode(mode string) OutputMode {
	switch mode {
	case "text":
		return TextMode
	case "json":
		return JSONMode
	default:
		return TextMode
	}
}

// String returns the string representation of the output mode.
func (o OutputMode) String() string {
	switch o {
	case TextMode:
		return "PlainText"
	case JSONMode:
		return "JSON"
	default:
		return "Unknown"
	}
}
