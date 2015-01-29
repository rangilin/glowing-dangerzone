package main

// BaseTemplateContent is HTML for base.tmpl
const BaseTemplateContent = `{{ define "base" }}
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title></title>
  </head>
  <body>
    <div>{{ template "content" }}</div>
  </body>
</html>
{{ end }}
`

// IndexTemplateContent is HTML for index page
const IndexTemplateContent = `{{ define "content" }}
<div>Index</div>
{{ end }}
`

// PostTemplateContent is HTML for post page
const PostTemplateContent = `{{ define "content"}}
<div>{{.Content}}</div>
{{ end }}
`
