package generate

// promptName is a name for the template below.
const promptName = "name_brainstorm_with_json"

// promptTemplate is a Go template for prompting.
const promptTemplate = `Given the following parameters:

Project Purpose: {{.Purpose}}
Project Theme: {{.Theme}}
Target Demographics: {{.Demographics}}
Target Interests: {{.Interests}}

Generate 3-5 possible names for a side business / project.

For each of the names generated, please provide the following:

- An explanation of the name's fit for project purpose, audience and theme.
- First impressions from the audience when they hear the name.
- Pros and cons of the name.
- Any other thoughts you have about the name.

Format the output as a JSON array of objects, stack ranked based on your
assessment of their suitability. Provide a suitability score in each JSON
array element, anywhere between 1.0 and 5.0.

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
