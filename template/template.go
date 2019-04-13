package main

import (
	"html/template"
	"os"
)

const html = `
<script>var foo = {{.}};</script>
<a href="{{.URL}}">
	{{.Text}}
</a>
`

func main() {

	html := template.Must(template.New("example").Parse(html))


	data := struct {
		FOO  string
		URL  string
		Text string
	}{
		FOO:  `some "quoted" string`,
		URL:  `" onClick="alert('XSS!');`,
		Text: `The <- operator is for channel sends and receives`,
	}

	_ = html.Execute(os.Stdout, data)
}
