package search

// FormatOption represents the output mode.
type FormatOption struct {
	Name  string
	Value string
}

var (
	// TextOption is the text output mode.
	TextOption = FormatOption{Name: "PlainText", Value: "text"}

	// JSONOption is the JSON output mode.
	JSONOption = FormatOption{Name: "JSON", Value: "json"}
)

var formatOptions = map[string]FormatOption{
	TextOption.Value: TextOption,
	JSONOption.Value: JSONOption,
}

// GetFormatOption returns a FormatOption instance.
func GetFormatOption(format string) FormatOption {
	option, ok := formatOptions[format]
	if ok {
		return option
	}
	return TextOption
}

// GetAllFormatOptionValues returns all available format option values.
func GetAllFormatOptionValues() []string {
	var values []string
	for _, option := range formatOptions {
		values = append(values, option.Value)
	}
	return values
}

// String returns the string representation of the output mode.
func (o FormatOption) String() string {
	return o.Name
}
