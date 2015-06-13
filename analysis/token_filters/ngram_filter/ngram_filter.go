//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ngram_filter

import (
	"bytes"
	"fmt"
	"unicode/utf8"

	"github.com/tukdesk/bleve/analysis"
	"github.com/tukdesk/bleve/registry"
)

const Name = "ngram"

type NgramFilter struct {
	minLength int
	maxLength int
}

func NewNgramFilter(minLength, maxLength int) *NgramFilter {
	return &NgramFilter{
		minLength: minLength,
		maxLength: maxLength,
	}
}

func (s *NgramFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	rv := make(analysis.TokenStream, 0, len(input))

	for _, token := range input {
		runeCount := utf8.RuneCount(token.Term)
		runes := bytes.Runes(token.Term)
		for i := 0; i < runeCount; i++ {
			// index of the starting rune for this token
			for ngramSize := s.minLength; ngramSize <= s.maxLength; ngramSize++ {
				// build an ngram of this size starting at i
				if i+ngramSize <= runeCount {
					ngramTerm := buildTermFromRunes(runes[i : i+ngramSize])
					token := analysis.Token{
						Position: token.Position,
						Start:    token.Start,
						End:      token.End,
						Type:     token.Type,
						Term:     ngramTerm,
					}
					rv = append(rv, &token)
				}
			}
		}
	}

	return rv
}

func buildTermFromRunes(runes []rune) []byte {
	rv := make([]byte, 0, len(runes)*4)
	for _, r := range runes {
		runeBytes := make([]byte, utf8.RuneLen(r))
		utf8.EncodeRune(runeBytes, r)
		rv = append(rv, runeBytes...)
	}
	return rv
}

func NgramFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	minVal, ok := config["min"].(float64)
	if !ok {
		return nil, fmt.Errorf("must specify min")
	}
	min := int(minVal)
	maxVal, ok := config["max"].(float64)
	if !ok {
		return nil, fmt.Errorf("must specify max")
	}
	max := int(maxVal)

	return NewNgramFilter(min, max), nil
}

func init() {
	registry.RegisterTokenFilter(Name, NgramFilterConstructor)
}
