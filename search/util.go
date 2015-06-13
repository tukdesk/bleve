//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package search

func MergeLocations(locations []FieldTermLocationMap) FieldTermLocationMap {
	rv := locations[0]

	for i := 1; i < len(locations); i++ {
		nextLocations := locations[i]
		for field, outterMap := range nextLocations {
			rvArrayPositionsTermLocationMap, rvHasField := rv[field]
			if rvHasField {
				for arrayPositionsStr, termLocationMap := range outterMap {
					rvTermLocationMap, rvHasArrayPositionsStr := rv[field][arrayPositionsStr]
					if rvHasArrayPositionsStr {
						rv[field][arrayPositionsStr] = MergeTermLocationMaps(rvTermLocationMap, termLocationMap)
					} else {
						rv[field][arrayPositionsStr] = termLocationMap
					}
				}
			} else {
				rv[field] = rvArrayPositionsTermLocationMap
			}
		}
	}

	return rv
}

func MergeTermLocationMaps(rv, other TermLocationMap) TermLocationMap {
	for term, locationMap := range other {
		// for a given term/document there cannot be different locations
		// if they came back from different clauses, overwrite is ok
		rv[term] = locationMap
	}
	return rv
}
