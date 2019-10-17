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

package cmd

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/hasheddan/crds/pkg/acronym"
	"github.com/hasheddan/crds/pkg/static"
	"github.com/spf13/cobra"
)

var (
	f, _      = filepath.Abs("./public/templates/index.html")
	indexTmpl = template.Must(
		template.ParseFiles(f),
	)
)

type indexData struct {
	Acronym string
}

// helloCmd represents the hello command
var staticCmd = &cobra.Command{
	Use:   "static",
	Short: "Serves the static crds website.",
	Long:  `Serves the static crds website using golang templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		dataFunc := func() interface{} {
			return indexData{
				Acronym: acronym.Random(),
			}
		}

		log.Fatal(static.New("Started serving static site.", "Request recieved.", indexTmpl, dataFunc).Serve(port))
	},
}

func init() {
	rootCmd.AddCommand(staticCmd)
}
