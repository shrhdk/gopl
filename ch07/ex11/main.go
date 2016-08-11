package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
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
	http.HandleFunc("/item/", db.item)

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
	<h2>Create</h2>
	<form action="/item/" method="post">
		<input type="hidden" name="_method" value="POST" />
		<input type="text" name="item" placeholder="item" />
		<input type="text" name="price" placeholder="price" />
		<input type="submit" />
	</form>
	<h2>Get</h2>
	<form action="/item/" method="GET" onsubmit="this.action+=this.item.value; delete this.item;">
		<input type="text" name="item" placeholder="item" />
		<input type="submit" />
	</form>
	<h2>Update</h2>
	<form action="/item/" method="post" onsubmit="this.action+=this.item.value;">
		<input type="hidden" name="_method" value="PUT" />
		<input type="text" name="item" placeholder="item" />
		<input type="text" name="price" placeholder="price" />
		<input type="submit" />
	</form>
	<h2>Delete</h2>
	<form action="/item/" method="post" onsubmit="this.action+=this.item.value;">
		<input type="hidden" name="_method" value="DELETE" />
		<input type="text" name="item" placeholder="item" />
		<input type="submit" />
	</form>
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

func (db database) item(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	if method == "POST" && req.FormValue("_method") != "" {
		method = req.FormValue("_method")
	}

	switch method {
	case "POST":
		db.post(w, req)
	case "GET":
		db.get(w, req)
	case "PUT":
		db.put(w, req)
	case "DELETE":
		db.del(w, req)
	default:
		http.Error(w, "Unsupported Method: "+req.Method, 405)
	}
}

func (db database) post(w http.ResponseWriter, req *http.Request) {
	if _, file := path.Split(req.URL.Path); file != "" {
		http.Error(w, "Unsupported Method: "+req.Method, 405)
		return
	}

	name := req.FormValue("item")
	price, err := strconv.ParseFloat(req.FormValue("price"), 32)

	if err != nil {
		http.Error(w, "Invalid price: "+req.FormValue("price"), 400)
		return
	}

	db[name] = dollars(price)

	w.WriteHeader(201)
	fmt.Fprintf(w, "Created %s %s\n", name, db[name])
}

func (db database) get(w http.ResponseWriter, req *http.Request) {
	_, name := path.Split(req.URL.Path)

	price, ok := db[name]

	if !ok {
		http.Error(w, "No such item: "+name, 404)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) put(w http.ResponseWriter, req *http.Request) {
	_, name := path.Split(req.URL.Path)
	price, err := strconv.ParseFloat(req.FormValue("price"), 32)

	if err != nil {
		http.Error(w, "Invalid price: "+req.FormValue("price"), 400)
		return
	}

	oldPrice, ok := db[name]

	if !ok {
		http.Error(w, "No such item: "+name, 404)
		return
	}

	db[name] = dollars(price)

	w.WriteHeader(200)
	fmt.Fprintf(w, "Changed %s price %s from %s\n", name, db[name], oldPrice)
}

func (db database) del(w http.ResponseWriter, req *http.Request) {
	_, name := path.Split(req.URL.Path)
	_, ok := db[name]

	if !ok {
		http.Error(w, "No such item: "+name, 404)
		return
	}

	delete(db, name)

	w.WriteHeader(200)
	fmt.Fprintf(w, "Deleted %s\n", name)
}
