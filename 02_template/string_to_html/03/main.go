package main

import (
	"fmt"
	"os"
)

func main() {
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + os.Args[1] + `</h1>
	</body>
	</html>
	`

	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(tpl)
}
