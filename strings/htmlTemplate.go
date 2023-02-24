package main

import (
	"html/template"
	"log"
	"os"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head><meta name="viewport" content="width=device-width, initial-scale=1">
<meta charset="UTF-8">
<title>{{.Title}}</title>
</head>
<body>
{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
</body>
</html>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css">
`

func Example() {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(htmlTemplate)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(os.Stdout, noItems)
	check(err)

	// Output:
	// <!DOCTYPE html>
	// <html>
	// 	<head>
	// 		<meta charset="UTF-8">
	// 		<title>My page</title>
	// 	</head>
	// 	<body>
	// 		<div>My photos</div><div>My blog</div>
	// 	</body>
	// </html>
	// <!DOCTYPE html>
	// <html>
	// 	<head>
	// 		<meta charset="UTF-8">
	// 		<title>My another page</title>
	// 	</head>
	// 	<body>
	// 		<div><strong>no rows</strong></div>
	// 	</body>
	// </html>

}

func main() {
	Example()
}
