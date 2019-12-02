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
	"io/ioutil"
	"log"

	"github.com/hasheddan/crds/pkg/crder"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a CRD instance.",
	Long:  `Validates a CRD instance given a custom resource definition.`,
	Run: func(cmd *cobra.Command, args []string) {
		dat, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		crd, err := crder.NewCRDer(dat, true)
		if err != nil {
			panic(err)
		}
		inst, err := ioutil.ReadFile(args[1])
		if err != nil {
			panic(err)
		}
		if err := crd.Validate(inst); err != nil {
			panic(err)
		}

		log.Println("CRD instance is valid.")
	},
}
