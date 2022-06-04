package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	item := req.PostForm.Get("item")
	priceStr := req.PostForm.Get("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
		return
	}

	if _, ok := db[item]; !ok {
		db[item] = dollars(price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 400
		fmt.Fprintf(w, "already exists item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	item := req.PostForm.Get("item")
	priceStr := req.PostForm.Get("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
		return
	}

	if _, ok := db[item]; ok {
		db[item] = dollars(price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 400
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; ok {
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 400
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
