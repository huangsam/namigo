package generate

// promptName is a name for the template below.
const promptName = "name_brainstorm_with_json"

// promptData is the input for the prompt template.
type promptData struct {
	Purpose      string
	Demographics string
	Interests    string
	Theme        string
	MaxCount     int
	MaxLength    int
}

// promptTemplate is a Go template for prompting.
const promptTemplate = `Given the following parameters:

- Project Purpose: {{.Purpose}}
- Project Theme: {{.Theme}}
- Target Demographics: {{.Demographics}}
- Target Interests: {{.Interests}}

Generate up to {{.MaxCount}} names for a side business / project. All names
should have at most {{.MaxLength}} characters.

For each of the names generated, please provide the following:

- An explanation of the name's fit for project purpose, audience and theme.
- First impressions from the audience when they hear the name.
- Pros and cons of the name.
- Any other thoughts you have about the name.

Format the output as a JSON array of objects, stack ranked based on your
assessment of their suitability. Provide a suitability score in each JSON
array element, anywhere between 1.0 and 10.0.

The JSON output should adhere to the following structure:

[
  {
    "name": "Generated Name",
    "explanation": "Explanation of fit...",
    "firstImpressions": "First impressions...",
    "pros": ["Pro 1", "Pro 2"],
    "cons": ["Con 1", "Con 2"],
    "additionalThoughts": "Any other thoughts...",
    "suitabilityScore": 1.0
  }
]
`
