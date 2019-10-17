/*
Copyright 2019 The CRDS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package static

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// DataFunc is a function that returns data used to populate a template.
type DataFunc func() interface{}

// Server will serve a static html page.
type Server struct {
	startupMsg string
	logMsg     string
	template   *template.Template
	dataFunc   DataFunc
}

// New creates a new static Server.
func New(startupMsg string, logMsg string, template *template.Template, dataFunc DataFunc) *Server {
	return &Server{
		startupMsg: startupMsg,
		logMsg:     logMsg,
		template:   template,
		dataFunc:   dataFunc,
	}
}

// Serve starts a static site http server.
func (s *Server) Serve(port string) error {
	log.Print("Started serving static site.")

	http.HandleFunc("/", handler(s.logMsg, s.template, s.dataFunc))

	if port == "" {
		port = "8080"
	}
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handler(logMsg string, template *template.Template, datafunc DataFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := datafunc()
		log.Print(logMsg)
		if err := template.Execute(w, data); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}
