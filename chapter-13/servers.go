package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(res, `
<Doctype html>
<html>
<head>
	<title>David Saint</title>
</head>
<body>
	<p>This is my very first GO server</p>
</body>
</html>
		`)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
