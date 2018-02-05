package main

import (
	"github.com/appscode/go/ioutil"
	"log"
	"github.com/tamalsaha/go-oneliners"
	"time"
)

func main() {
	w, err := ioutil.NewAtomicWriter("/tmp/tester")
	if err != nil {
		log.Fatal(err)
	}
	payload := map[string]ioutil.FileProjection{
		"b.txt": {[]byte(time.Now().String()), 0755},
	}
	oneliners.FILE(w.Write(payload))
}