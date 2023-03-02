package main

const (
	configTemplate = `{
  "Contacts": ["mailto:security[at]EXAMPLE.com", "https://hackerone.com/EXAMPLECO"],
  "Acknowledgments": "https://hackerone.com/EXAMPLECO/thanks?type=team",
  "Languages": "en, XX",
  "Canonical": "https://www.example.com/.well-known/security.txt",
  "Policy": "https://hackerone.com/EXAMPLECO/policy"
}`

	securityTemplate = `
{{range .Contacts}}Contact: {{.}}
{{end}}Expires: {{.Expires}}
Acknowledgments: {{.Acknowledgments}}
Preferred-Languages: {{.Languages}}
Canonical: {{.Canonical}}
Policy: {{.Policy}}
`
)
