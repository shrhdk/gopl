package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// GET    /list
// POST   /?item={name}&price={price} -> Create Item
// GET    /{name}                     -> Get Item Price
// PUT    /{name}?price={price}       -> Update Item Price
// DELETE /{name}                     -> Delete Item

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/", db.list)

	sockAddress := "localhost:8000"
	fmt.Println("Serving on " + sockAddress)
	log.Fatal(http.ListenAndServe(sockAddress, nil))
}

var dbList = template.Must(template.New("dblist").Parse(`<!DOCTYPE html>
<html>
<head>
	<title>DB</title>
</head>
<body>
	<h1>DB</h1>
	<h2>List</h2>
	<table>
		<tr>
		<th>Item</th>
		<th>Price</th>
		</tr>
		{{range $index, $element := .}}
		<tr>
			<td>{{$index}}</td>
			<td>{{$element}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>`))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := dbList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}
