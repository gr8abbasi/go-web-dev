package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}

	name := "Kashif Abbasi"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	f.WriteString(tpl)
}
