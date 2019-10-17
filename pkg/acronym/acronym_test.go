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

package acronym

import (
	"regexp"
	"testing"
)

func wordCount(value string) int {
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(value, -1)
	return len(results)
}

func TestRandom(t *testing.T) {
	cases := []struct {
		name      string
		wordCount int
	}{
		{
			name:      "Random Acronym",
			wordCount: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Random()
			if wordCount(got) != tc.wordCount {
				t.Errorf("Random(...): length got %v, want %v", got, tc.wordCount)
			}
		})
	}
}
