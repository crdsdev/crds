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
	"html/template"
	"testing"
)

func TestNew(t *testing.T) {
	blankTemp := &template.Template{}
	empty := func() interface{} { return nil }

	cases := []struct {
		name       string
		startupMsg string
		logMsg     string
		staticDir  string
		template   *template.Template
		dataFunc   DataFunc
		want       *Server
	}{
		{
			name:       "Successful",
			startupMsg: "Some message.",
			logMsg:     "Some other message.",
			staticDir:  "public/assets",
			template:   blankTemp,
			dataFunc:   empty,
			want: &Server{
				startupMsg: "Some message.",
				logMsg:     "Some other message.",
				staticDir:  "public/assets",
				template:   blankTemp,
				dataFunc:   empty,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := New(tc.startupMsg, tc.logMsg, tc.staticDir, tc.template, tc.dataFunc)
			if got.startupMsg != tc.want.startupMsg {
				t.Errorf("New(...): got %v, want %v", got.startupMsg, tc.want.startupMsg)
			}
			if got.logMsg != tc.want.logMsg {
				t.Errorf("New(...): got %v, want %v", got.logMsg, tc.want.logMsg)
			}
			if got.staticDir != tc.want.staticDir {
				t.Errorf("New(...): got %v, want %v", got.staticDir, tc.want.staticDir)
			}
			if got.template != tc.want.template {
				t.Errorf("New(...): got %v, want %v", got.template, tc.want.template)
			}
			if got.dataFunc() != tc.want.dataFunc() {
				t.Errorf("New(...): got %v, want %v", got.dataFunc(), tc.want.dataFunc())
			}
		})
	}
}
