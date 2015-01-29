// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine,secure

package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/Go-zh/tools/playground" // this registers the /compile and /share handlers
	"github.com/Go-zh/tools/present"
)

var (
	basePath = flag.String("base", "./present/", "base path of tools/cmd/present")
	httpAddr = flag.String("http", "127.0.0.1:8082", "HTTP service address (e.g., '127.0.0.1:8082')")
)

func main() {
	flag.Parse()

	initTemplates(*basePath)
	playScript(*basePath, "HTTPTransport")
	present.PlayEnabled = true
	http.Handle("/static/", http.FileServer(http.Dir(*basePath)))

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

func playable(c present.Code) bool {
	return present.PlayEnabled && c.Play && c.Ext == ".go"
}
