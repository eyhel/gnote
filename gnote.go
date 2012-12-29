// Copyright 2012 The NWS Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
)

/*
const (
	GNOTE_DIR = "./gnote"
)
*/

var GNOTE_DIR string
var config map[string] string

func init() {
	file, err := os.Open("config.json.default")
	if err != nil {
		println("[SYSTEM] load configure file failed")
		panic(err)
		os.Exit(1)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&config)

	if err != nil {
		println("[SYSTEM] read configure file failed")
		panic(err)
		os.Exit(1)
	}

	GNOTE_DIR = config["gnote_dir"]
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", GNOTE_DIR+"/assets", 1)
	mux.HandleFunc("/", safeHandler(gnoteHandler))
	mux.HandleFunc("/gnote/view", safeHandler(gnoteViewHandler))
	mux.HandleFunc("/gnote/add", safeHandler(gnoteAddHandler))
	mux.HandleFunc("/gnote/del", safeHandler(gnoteDeleteHandler))
	mux.HandleFunc("/gnote/change", safeHandler(gnoteChangeHandler))

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}