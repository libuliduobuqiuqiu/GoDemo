package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, prices := range db {
		fmt.Fprintf(w, "%s: %s \n", item, prices)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 20}
	log.Fatal(http.ListenAndServe(":8999", db))
}
