package main

// BaseTemplateContent is HTML for base.tmpl
const BaseTemplateContent = `<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title></title>
  </head>
  <body>
    <div>{{.Content}}</div>
  </body>
</html>
`

// IndexTemplateContent is HTML for index page
const IndexTemplateContent = `
`

// PostTemplateContent is HTML for post page
const PostTemplateContent = `<div>
  {{.Content}}
</div>
`
