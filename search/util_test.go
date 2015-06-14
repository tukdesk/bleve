//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package search

import (
	"reflect"
	"testing"
)

func TestMergeLocations(t *testing.T) {
	flm1 := FieldTermLocationMap{
		"marty": map[string]TermLocationMap{
			"_": TermLocationMap{
				"name": {
					&Location{
						Pos:   1,
						Start: 0,
						End:   5,
					},
				},
			},
		},
	}

	flm2 := FieldTermLocationMap{
		"marty": map[string]TermLocationMap{
			"_": TermLocationMap{
				"description": {
					&Location{
						Pos:   5,
						Start: 20,
						End:   25,
					},
				},
			},
		},
	}

	flm3 := FieldTermLocationMap{
		"marty": map[string]TermLocationMap{
			"_abc": TermLocationMap{
				"description": {
					&Location{
						Pos:   10,
						Start: 25,
						End:   40,
					},
				},
			},
		},
	}

	flm4 := FieldTermLocationMap{
		"josh": map[string]TermLocationMap{
			"_": TermLocationMap{
				"description": {
					&Location{
						Pos:   5,
						Start: 20,
						End:   25,
					},
				},
			},
		},
	}

	expectedMerge := FieldTermLocationMap{
		"marty": map[string]TermLocationMap{
			"_": TermLocationMap{
				"description": {
					&Location{
						Pos:   5,
						Start: 20,
						End:   25,
					},
				},
				"name": {
					&Location{
						Pos:   1,
						Start: 0,
						End:   5,
					},
				},
			},
			"_abc": TermLocationMap{
				"description": {
					&Location{
						Pos:   10,
						Start: 25,
						End:   40,
					},
				},
			},
		},
		"josh": map[string]TermLocationMap{
			"_": TermLocationMap{
				"description": {
					&Location{
						Pos:   5,
						Start: 20,
						End:   25,
					},
				},
			},
		},
	}

	mergedLocations := MergeLocations([]FieldTermLocationMap{flm1, flm2, flm3, flm4})
	if !reflect.DeepEqual(expectedMerge, mergedLocations) {
		t.Errorf("expected %v, got %v", expectedMerge, mergedLocations)
	}
}
