# tplxjson

A super simple tool for rendering JSON into go-style template. 

## Demand

We have a JSON file in `/tmp/j.json`:

```json
{
  "Title" : "404 Not Found",
  "Html"  : "<h1>404 Not Found</h1><p>File does not exist.</p>"
}
```

And we have a go-style template in `/tmp/t.tpl`:

```html
<html><head>

  <meta charset='utf8'>
  <title>{{.Title}}</title>

</head><body>

  {{.Content}}

</body></html>
```

Now we want to use command to combine this two files together.

## Solution

```sh
go get github.com/fr440305/tplxjson
go install github.com/fr440305/tplxjson

tplxjson -tpl=/tmp/t.tpl -json=/tmp/j.json
```

And we should be able to see the console output.

```html
<html><head>

  <meta charset='utf8'>
  <title>404 Not Found</title>

</head><body>

  <h1>404 Not Found</h1><p>File does not exist.</p>

</body></html>
```

