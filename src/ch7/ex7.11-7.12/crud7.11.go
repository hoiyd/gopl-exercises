package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mutex sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	mutex.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mutex.Lock()
	price, ok := db[item]
	mutex.Unlock()
	if !ok {
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	if price, err := strconv.Atoi(priceStr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong price: %s\n", priceStr)
		return
	} else {
		mutex.Lock()
		db[item] = dollars(price)
		mutex.Unlock()
		fmt.Fprintf(w, "created %s: $%d\n", item, price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	if price, err := strconv.Atoi(priceStr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong price: %s\n", priceStr)
		return
	} else {
		mutex.Lock()
		if _, ok := db[item]; !ok {
			mutex.Unlock()
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "item does not exits: %s\n", item)
		} else {
			db[item] = dollars(price)
			mutex.Unlock()
			fmt.Fprintf(w, "updated %s: $%d\n", item, price)
		}
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mutex.Lock()
	if _, ok := db[item]; !ok {
		mutex.Unlock()
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item does not exits: %s\n", item)
	} else {
		delete(db, item)
		mutex.Unlock()
		fmt.Fprintf(w, "item deleted: %s\n", item)
	}
}
