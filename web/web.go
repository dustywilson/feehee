package main

import (
	"feehee"
	"fmt"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
	"github.com/nytimes/gziphandler"
)

var db *bolt.DB

func main() {
	db, err = bolt.Open("feehee.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	http.Handle("/", gziphandler.GzipHandler(http.FileServer(http.Dir("www"))))
	http.ListenAndServe(":8088", nil)
}
